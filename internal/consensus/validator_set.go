package consensus

import (
	"fmt"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
	"github.com/IWALINK/cosine/pkg/p2p"
	"github.com/libp2p/go-libp2p/core/peer"
)

// Validator represents a COSINE network validator
type Validator struct {
	// Peer ID of the validator
	ID peer.ID
	
	// Staked tokens (minimum 100,000)
	Stake uint64
	
	// Performance score in [0,1]
	Performance float64
	
	// Last ping timestamp
	LastPing time.Time
	
	// Node multiaddress for network communication
	Multiaddress string
}

// ValidatorStatus represents the validator's current status
type ValidatorStatus string

const (
	// StatusOnline indicates validator is online and actively participating
	StatusOnline ValidatorStatus = "online"
	
	// StatusOffline indicates validator hasn't pinged recently
	StatusOffline ValidatorStatus = "offline"
	
	// StatusSlashed indicates validator has been slashed for misbehavior
	StatusSlashed ValidatorStatus = "slashed"
)

// ValidatorSet manages the set of validators in the COSINE network
type ValidatorSet struct {
	validators     map[peer.ID]*Validator
	messaging      *p2p.MessagingService
	logger         *utils.Logger
	config         *utils.ConfigManager
	mu             sync.RWMutex
	pingThreshold  time.Duration
	minStake       uint64
	changeHandlers []func(peer.ID, *Validator, ValidatorStatus)
}

// NewValidatorSet creates a new validator set
func NewValidatorSet(messagingService *p2p.MessagingService, config *utils.ConfigManager, logger *utils.Logger) *ValidatorSet {
	// Get ping threshold from config (default to 5 minutes)
	pingThreshold := config.GetDuration("consensus.pingThreshold")
	if pingThreshold == 0 {
		pingThreshold = 5 * time.Minute
	}
	
	// Get minimum stake (default to 100,000 tokens)
	minStake := uint64(config.GetInt("consensus.minStake"))
	if minStake == 0 {
		minStake = 100000
	}
	
	vs := &ValidatorSet{
		validators:    make(map[peer.ID]*Validator),
		messaging:     messagingService,
		logger:        logger.WithComponent("ValidatorSet"),
		config:        config,
		pingThreshold: pingThreshold,
		minStake:      minStake,
		changeHandlers: []func(peer.ID, *Validator, ValidatorStatus){},
	}
	
	// Start background monitoring
	go vs.startMonitoring()
	
	// Register handler for ping messages
	if messagingService != nil {
		messagingService.RegisterHandler(p2p.TypePing, vs.handlePingMessage)
	}
	
	return vs
}

// AddValidator adds a new validator to the set
func (vs *ValidatorSet) AddValidator(id peer.ID, stake uint64, performance float64, multiaddr string) error {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	
	// Check if stake meets minimum requirement
	if stake < vs.minStake {
		return fmt.Errorf("insufficient stake: %d (minimum: %d)", stake, vs.minStake)
	}
	
	// Clamp performance score to [0,1]
	if performance < 0 {
		performance = 0
	} else if performance > 1 {
		performance = 1
	}
	
	// Add or update validator
	_, exists := vs.validators[id]
	vs.validators[id] = &Validator{
		ID:           id,
		Stake:        stake,
		Performance:  performance,
		LastPing:     time.Now(),
		Multiaddress: multiaddr,
	}
	
	vs.logger.Info("Validator added/updated",
		"id", id.String(),
		"stake", stake,
		"performance", performance,
		"multiaddress", multiaddr)
	
	// Notify handlers
	status := StatusOnline
	for _, handler := range vs.changeHandlers {
		handler(id, vs.validators[id], status)
	}
	
	// Log whether this was an add or update
	if exists {
		vs.logger.Debug("Updated existing validator", "id", id.String())
	} else {
		vs.logger.Debug("Added new validator", "id", id.String())
	}
	
	return nil
}

// RemoveValidator removes a validator from the set
func (vs *ValidatorSet) RemoveValidator(id peer.ID) bool {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	
	validator, exists := vs.validators[id]
	if !exists {
		return false
	}
	
	delete(vs.validators, id)
	
	vs.logger.Info("Validator removed", "id", id.String())
	
	// Notify handlers
	for _, handler := range vs.changeHandlers {
		handler(id, validator, StatusSlashed)
	}
	
	return true
}

// UpdateValidatorPerformance updates a validator's performance score
func (vs *ValidatorSet) UpdateValidatorPerformance(id peer.ID, delta float64) error {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	
	validator, exists := vs.validators[id]
	if !exists {
		return fmt.Errorf("validator not found: %s", id.String())
	}
	
	// Update performance score with bounds checking
	newPerformance := validator.Performance + delta
	if newPerformance < 0 {
		newPerformance = 0
	} else if newPerformance > 1 {
		newPerformance = 1
	}
	
	validator.Performance = newPerformance
	
	vs.logger.Debug("Updated validator performance",
		"id", id.String(),
		"delta", delta,
		"new_performance", newPerformance)
	
	return nil
}

// UpdateValidatorStake updates a validator's stake
func (vs *ValidatorSet) UpdateValidatorStake(id peer.ID, newStake uint64) error {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	
	validator, exists := vs.validators[id]
	if !exists {
		return fmt.Errorf("validator not found: %s", id.String())
	}
	
	// Check minimum stake
	if newStake < vs.minStake {
		return fmt.Errorf("insufficient stake: %d (minimum: %d)", newStake, vs.minStake)
	}
	
	validator.Stake = newStake
	
	vs.logger.Debug("Updated validator stake",
		"id", id.String(),
		"new_stake", newStake)
	
	return nil
}

// GetValidator returns a validator by ID
func (vs *ValidatorSet) GetValidator(id peer.ID) (*Validator, bool) {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	
	validator, exists := vs.validators[id]
	return validator, exists
}

// GetOnlineValidators returns all validators that have pinged within the threshold
func (vs *ValidatorSet) GetOnlineValidators() map[peer.ID]*Validator {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	
	result := make(map[peer.ID]*Validator)
	cutoff := time.Now().Add(-vs.pingThreshold)
	
	for id, validator := range vs.validators {
		if validator.LastPing.After(cutoff) {
			result[id] = validator
		}
	}
	
	return result
}

// GetAllValidators returns all validators regardless of online status
func (vs *ValidatorSet) GetAllValidators() map[peer.ID]*Validator {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	
	// Create a copy to prevent external modification
	result := make(map[peer.ID]*Validator, len(vs.validators))
	for id, validator := range vs.validators {
		validatorCopy := *validator
		result[id] = &validatorCopy
	}
	
	return result
}

// IsOnline checks if a validator is online
func (vs *ValidatorSet) IsOnline(id peer.ID) bool {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	
	validator, exists := vs.validators[id]
	if !exists {
		return false
	}
	
	return time.Since(validator.LastPing) <= vs.pingThreshold
}

// RegisterChangeHandler registers a handler for validator status changes
func (vs *ValidatorSet) RegisterChangeHandler(handler func(peer.ID, *Validator, ValidatorStatus)) {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	
	vs.changeHandlers = append(vs.changeHandlers, handler)
}

// UpdateLastPing updates a validator's last ping time
func (vs *ValidatorSet) UpdateLastPing(id peer.ID) error {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	
	validator, exists := vs.validators[id]
	if !exists {
		return fmt.Errorf("validator not found: %s", id.String())
	}
	
	wasOffline := !vs.isOnlineNoLock(validator)
	validator.LastPing = time.Now()
	
	// If validator was offline but is now online, notify handlers
	if wasOffline {
		for _, handler := range vs.changeHandlers {
			handler(id, validator, StatusOnline)
		}
	}
	
	return nil
}

// Helper method: check if validator is online without lock
func (vs *ValidatorSet) isOnlineNoLock(validator *Validator) bool {
	return time.Since(validator.LastPing) <= vs.pingThreshold
}

// GetOnlineCount returns the number of online validators
func (vs *ValidatorSet) GetOnlineCount() int {
	return len(vs.GetOnlineValidators())
}

// GetTotalCount returns the total number of validators
func (vs *ValidatorSet) GetTotalCount() int {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	
	return len(vs.validators)
}

// GetTotalStake returns the total stake of all validators
func (vs *ValidatorSet) GetTotalStake() uint64 {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	
	var total uint64
	for _, validator := range vs.validators {
		total += validator.Stake
	}
	
	return total
}

// GetOnlineStake returns the total stake of online validators
func (vs *ValidatorSet) GetOnlineStake() uint64 {
	vs.mu.RLock()
	defer vs.mu.RUnlock()
	
	var total uint64
	cutoff := time.Now().Add(-vs.pingThreshold)
	
	for _, validator := range vs.validators {
		if validator.LastPing.After(cutoff) {
			total += validator.Stake
		}
	}
	
	return total
}

// Start background monitoring for offline validators
func (vs *ValidatorSet) startMonitoring() {
	// Check interval - half of the ping threshold to be proactive
	interval := vs.pingThreshold / 2
	if interval < time.Minute {
		interval = time.Minute
	}
	
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	for range ticker.C {
		vs.checkOfflineValidators()
	}
}

// Check for validators that have gone offline
func (vs *ValidatorSet) checkOfflineValidators() {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	
	cutoff := time.Now().Add(-vs.pingThreshold)
	
	for id, validator := range vs.validators {
		wasOnline := validator.LastPing.After(cutoff)
		isOnline := time.Now().Sub(validator.LastPing) <= vs.pingThreshold
		
		// Status changed from online to offline
		if wasOnline && !isOnline {
			vs.logger.Debug("Validator went offline",
				"id", id.String(),
				"last_ping", validator.LastPing.Format(time.RFC3339))
			
			// Notify handlers
			for _, handler := range vs.changeHandlers {
				handler(id, validator, StatusOffline)
			}
		}
	}
}

// Handler for ping messages from validators
func (vs *ValidatorSet) handlePingMessage(msg *p2p.Message, from peer.ID) error {
	err := vs.UpdateLastPing(from)
	if err != nil {
		vs.logger.Debug("Received ping from unknown validator", 
			"peer", from.String())
		return nil
	}
	
	// Send pong response if this is a ping
	if msg.Type == p2p.TypePing {
		pongMsg := &p2p.Message{
			Type:      p2p.TypePong,
			Timestamp: time.Now().UnixNano(),
			Sender:    "", // Will be filled by messaging service
			Target:    from.String(),
			Payload:   map[string]interface{}{},
		}
		
		err := vs.messaging.SendDirect(from, pongMsg)
		if err != nil {
			vs.logger.Error("Failed to send pong response",
				"peer", from.String(),
				"error", err)
		}
	}
	
	return nil
}