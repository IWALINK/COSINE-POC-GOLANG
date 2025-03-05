package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sync"
	"time"
    "fmt"
	"github.com/IWALINK/cosine/internal/utils"
)

// ScalingFactor represents a dynamic scaling factor with Kalman filter state
type ScalingFactor struct {
	Value float64   `json:"value"`    // Current scaling factor value
	P     float64   `json:"p_t"`      // Kalman filter error variance
	LastUpdated time.Time `json:"last_updated"`
}

// ValidatorInfo stores information about a validator
type ValidatorInfo struct {
	Stake       uint64    `json:"stake"`        // Staked tokens (minimum 100,000)
	Performance float64   `json:"performance"`  // Performance score in [0, 1]
	LastPing    time.Time `json:"last_ping"`    // Last ping timestamp
}

// GlobalStateObject represents the shared state for the entire protocol
type GlobalStateObject struct {
	mu sync.RWMutex
	logger *utils.Logger

	// Scaling factors for various credit score shifts
	ScalingFactors map[string]ScalingFactor `json:"scaling_factors"`

	// Network reward pool (COSINE tokens)
	NetworkRewardPool uint64 `json:"network_reward_pool"`

	// Total token supply on L2
	TotalTokenSupplyL2 uint64 `json:"total_token_supply_l2"`

	// Validator information keyed by validator address
	Validators map[string]ValidatorInfo `json:"validators"`

	// Association mapping from (chain_id, L1_address) to L2_wallet_address
	AssociationMapping map[string]string `json:"association_mapping"`

	// Bridged token states by chain ID
	BridgedTokens map[uint64]uint64 `json:"bridged_tokens"`

	// Last updated timestamp
	LastUpdated time.Time `json:"last_updated"`
}

// NewGlobalStateObject creates a new global state object with default values
func NewGlobalStateObject(logger *utils.Logger) *GlobalStateObject {
	now := time.Now()
	
	return &GlobalStateObject{
		logger: logger.WithComponent("GlobalStateObject"),
		
		// Initialize scaling factors with defaults
		ScalingFactors: map[string]ScalingFactor{
			"K_neg":   {Value: 1.0, P: 1.0, LastUpdated: now},
			"K_pos":   {Value: 1.0, P: 1.0, LastUpdated: now},
			"K_assoc": {Value: 1.0, P: 1.0, LastUpdated: now},
			"K_rehab": {Value: 1.0, P: 1.0, LastUpdated: now},
		},
		
		// Initialize with 40% of total supply (1 billion tokens)
		NetworkRewardPool: 400000000,
		
		// Initialize with remaining 60% of total supply
		TotalTokenSupplyL2: 600000000,
		
		// Empty maps initially
		Validators:         make(map[string]ValidatorInfo),
		AssociationMapping: make(map[string]string),
		BridgedTokens:      make(map[uint64]uint64),
		
		LastUpdated: now,
	}
}

// UpdateScalingFactor updates a scaling factor using the Kalman filter approach
func (gso *GlobalStateObject) UpdateScalingFactor(name string, observedImpact, expectedImpact float64, q, r float64) float64 {
	gso.mu.Lock()
	defer gso.mu.Unlock()
	
	sf, exists := gso.ScalingFactors[name]
	if !exists {
		sf = ScalingFactor{Value: 1.0, P: 1.0, LastUpdated: time.Now()}
	}
	
	// Create a Kalman filter with current values
	kf := utils.NewDynamicScalingKalmanFilter(sf.Value, sf.P, q, r)
	
	// Update scaling factor
	newValue := kf.UpdateWithImpact(observedImpact, expectedImpact)
	
	// Store updated values
	sf.Value = newValue
	sf.P = kf.P
	sf.LastUpdated = time.Now()
	gso.ScalingFactors[name] = sf
	gso.LastUpdated = time.Now()
	
	gso.logger.Debug("Updated scaling factor", 
		"name", name, 
		"old_value", gso.ScalingFactors[name].Value, 
		"new_value", newValue,
		"observed_impact", observedImpact,
		"expected_impact", expectedImpact)
	
	return newValue
}

// GetScalingFactor returns a scaling factor value
func (gso *GlobalStateObject) GetScalingFactor(name string) float64 {
	gso.mu.RLock()
	defer gso.mu.RUnlock()
	
	sf, exists := gso.ScalingFactors[name]
	if !exists {
		return 1.0 // Default value
	}
	return sf.Value
}

// DecrementRewardPool reduces the network reward pool by the specified amount
func (gso *GlobalStateObject) DecrementRewardPool(amount uint64) error {
	gso.mu.Lock()
	defer gso.mu.Unlock()
	
	if gso.NetworkRewardPool < amount {
		return ErrInsufficientRewardPool
	}
	
	gso.NetworkRewardPool -= amount
	gso.LastUpdated = time.Now()
	
	gso.logger.Debug("Decremented reward pool",
		"amount", amount, 
		"remaining", gso.NetworkRewardPool)
	
	return nil
}

// AddValidator adds or updates a validator
func (gso *GlobalStateObject) AddValidator(address string, stake uint64, performance float64) error {
	gso.mu.Lock()
	defer gso.mu.Unlock()
	
	// Check minimum stake requirement (100,000 tokens)
	if stake < 100000 {
		return ErrInsufficientValidatorStake
	}
	
	gso.Validators[address] = ValidatorInfo{
		Stake:       stake,
		Performance: utils.Clamp(performance, 0, 1),
		LastPing:    time.Now(),
	}
	
	gso.LastUpdated = time.Now()
	
	gso.logger.Debug("Added/updated validator",
		"address", address,
		"stake", stake,
		"performance", performance)
	
	return nil
}

// UpdateValidatorPerformance updates a validator's performance score
func (gso *GlobalStateObject) UpdateValidatorPerformance(address string, delta float64) error {
	gso.mu.Lock()
	defer gso.mu.Unlock()
	
	validator, exists := gso.Validators[address]
	if !exists {
		return ErrValidatorNotFound
	}
	
	validator.Performance = utils.Clamp(validator.Performance+delta, 0, 1)
	validator.LastPing = time.Now()
	gso.Validators[address] = validator
	gso.LastUpdated = time.Now()
	
	gso.logger.Debug("Updated validator performance",
		"address", address,
		"delta", delta,
		"new_performance", validator.Performance)
	
	return nil
}

// PingValidator updates a validator's last ping time
func (gso *GlobalStateObject) PingValidator(address string) error {
	gso.mu.Lock()
	defer gso.mu.Unlock()
	
	validator, exists := gso.Validators[address]
	if !exists {
		return ErrValidatorNotFound
	}
	
	validator.LastPing = time.Now()
	gso.Validators[address] = validator
	
	return nil
}

// GetOnlineValidators returns validators that have pinged within the threshold
func (gso *GlobalStateObject) GetOnlineValidators(pingThreshold time.Duration) map[string]ValidatorInfo {
	gso.mu.RLock()
	defer gso.mu.RUnlock()
	
	result := make(map[string]ValidatorInfo)
	now := time.Now()
	
	for addr, validator := range gso.Validators {
		if now.Sub(validator.LastPing) <= pingThreshold {
			result[addr] = validator
		}
	}
	
	return result
}

// AddAssociationMapping adds a mapping from L1 address to L2 wallet
func (gso *GlobalStateObject) AddAssociationMapping(chainID uint64, l1Address string, l2Address string) {
	gso.mu.Lock()
	defer gso.mu.Unlock()
	
	key := makeAssociationKey(chainID, l1Address)
	gso.AssociationMapping[key] = l2Address
	gso.LastUpdated = time.Now()
	
	gso.logger.Debug("Added association mapping",
		"chain_id", chainID,
		"l1_address", l1Address,
		"l2_address", l2Address)
}

// GetL2AddressForL1 returns the L2 address associated with an L1 address
func (gso *GlobalStateObject) GetL2AddressForL1(chainID uint64, l1Address string) (string, bool) {
	gso.mu.RLock()
	defer gso.mu.RUnlock()
	
	key := makeAssociationKey(chainID, l1Address)
	l2Address, exists := gso.AssociationMapping[key]
	return l2Address, exists
}

// Lock/mint tokens from L2 to L1
func (gso *GlobalStateObject) LockTokensForL1(chainID uint64, amount uint64) error {
	gso.mu.Lock()
	defer gso.mu.Unlock()
	
	if gso.TotalTokenSupplyL2 < amount {
		return ErrInsufficientL2Supply
	}
	
	gso.TotalTokenSupplyL2 -= amount
	gso.BridgedTokens[chainID] += amount
	gso.LastUpdated = time.Now()
	
	gso.logger.Debug("Locked tokens for L1",
		"chain_id", chainID,
		"amount", amount,
		"total_bridged", gso.BridgedTokens[chainID],
		"remaining_l2_supply", gso.TotalTokenSupplyL2)
	
	return nil
}

// Unlock/redeem tokens from L1 to L2
func (gso *GlobalStateObject) UnlockTokensFromL1(chainID uint64, amount uint64) error {
	gso.mu.Lock()
	defer gso.mu.Unlock()
	
	current, exists := gso.BridgedTokens[chainID]
	if !exists || current < amount {
		return ErrInsufficientBridgedTokens
	}
	
	gso.BridgedTokens[chainID] -= amount
	gso.TotalTokenSupplyL2 += amount
	gso.LastUpdated = time.Now()
	
	gso.logger.Debug("Unlocked tokens from L1",
		"chain_id", chainID,
		"amount", amount,
		"remaining_bridged", gso.BridgedTokens[chainID],
		"new_l2_supply", gso.TotalTokenSupplyL2)
	
	return nil
}

// VerifySupplyInvariant checks that the total supply equals 1 billion
func (gso *GlobalStateObject) VerifySupplyInvariant() bool {
	gso.mu.RLock()
	defer gso.mu.RUnlock()
	
	totalBridged := uint64(0)
	for _, amount := range gso.BridgedTokens {
		totalBridged += amount
	}
	
	return (gso.TotalTokenSupplyL2 + totalBridged) == 1000000000
}

// GetHash returns a hash of the global state for committing to L2 blocks
func (gso *GlobalStateObject) GetHash() string {
	gso.mu.RLock()
	defer gso.mu.RUnlock()
	
	// Serialize to JSON for hashing
	data, err := json.Marshal(gso)
	if err != nil {
		gso.logger.Error("Failed to marshal GSO for hashing", "error", err)
		return ""
	}
	
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// Helper function to create a key for the association mapping
func makeAssociationKey(chainID uint64, l1Address string) string {
	return fmt.Sprintf("%d:%s", chainID, l1Address)
}

// Custom errors
var (
	ErrInsufficientRewardPool    = fmt.Errorf("insufficient reward pool")
	ErrInsufficientValidatorStake = fmt.Errorf("insufficient validator stake (minimum 100,000 required)")
	ErrValidatorNotFound         = fmt.Errorf("validator not found")
	ErrInsufficientL2Supply      = fmt.Errorf("insufficient L2 token supply")
	ErrInsufficientBridgedTokens = fmt.Errorf("insufficient bridged tokens")
)