package consensus

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

// ValidatorScore combines a validator's ID with their VRF-derived score
type ValidatorScore struct {
	ID        peer.ID
	Score     float64
	Validator *Validator
	VRFOutput *VRFOutput
}

// SelectionResult represents the outcome of a validator selection round
type SelectionResult struct {
	// Selected validators subset
	SelectedValidators []peer.ID
	
	// The seed that was used for selection
	Seed []byte
	
	// Timestamp of the selection
	Timestamp time.Time
	
	// ID of the selection round (derived from seed)
	RoundID string
	
	// Complete selection data with scores
	ScoreData []ValidatorScore
}

// SubsetSelector handles VRF-based validator subset selection
type SubsetSelector struct {
	validatorSet     *ValidatorSet
	vrfManager       *VRFManager
	logger           *utils.Logger
	config           *utils.ConfigManager
	alpha            float64
	beta             float64
	gamma            float64
	maxValidators    int
	selectionTimeout time.Duration
	mu               sync.RWMutex
	lastSelection    *SelectionResult
}

// NewSubsetSelector creates a new validator subset selector
func NewSubsetSelector(
	validatorSet *ValidatorSet,
	vrfManager *VRFManager,
	config *utils.ConfigManager,
	logger *utils.Logger,
) *SubsetSelector {
	// Get alpha parameter (default to 0.5)
	alpha := config.GetFloat64("consensus.alpha")
	if alpha <= 0 {
		alpha = 0.5
	}
	
	// Get beta parameter (default to 0.2)
	beta := config.GetFloat64("consensus.beta")
	if beta <= 0 {
		beta = 0.2
	}
	
	// Get gamma parameter (default to 0.1)
	gamma := config.GetFloat64("consensus.gamma")
	if gamma <= 0 {
		gamma = 0.1
	}
	
	// Get maximum validators (default to 100)
	maxValidators := config.GetInt("consensus.maxValidators")
	if maxValidators <= 0 {
		maxValidators = 100
	}
	
	// Get selection timeout (default to 30 seconds)
	selectionTimeout := config.GetDuration("consensus.selectionTimeout")
	if selectionTimeout == 0 {
		selectionTimeout = 30 * time.Second
	}
	
	return &SubsetSelector{
		validatorSet:     validatorSet,
		vrfManager:       vrfManager,
		logger:           logger.WithComponent("SubsetSelector"),
		config:           config,
		alpha:            alpha,
		beta:             beta, 
		gamma:            gamma,
		maxValidators:    maxValidators,
		selectionTimeout: selectionTimeout,
	}
}

// SelectValidatorSubset selects a subset of validators based on VRF output
func (ss *SubsetSelector) SelectValidatorSubset(seed []byte) (*SelectionResult, error) {
	// Get online validators
	onlineValidators := ss.validatorSet.GetOnlineValidators()
	if len(onlineValidators) == 0 {
		return nil, fmt.Errorf("no online validators available")
	}
	
	// Calculate how many validators to select
	numSelected := int(math.Ceil(ss.gamma * float64(len(onlineValidators))))
	if numSelected > ss.maxValidators {
		numSelected = ss.maxValidators
	}
	if numSelected <= 0 {
		numSelected = 1 // At least one validator must be selected
	}
	
	// Create context with timeout for VRF computation
	ctx, cancel := context.WithTimeout(context.Background(), ss.selectionTimeout)
	defer cancel()
	
	// Calculate scores for all online validators
	scoresChan := make(chan ValidatorScore, len(onlineValidators))
	errChan := make(chan error, len(onlineValidators))
	
	// We need to compute our own VRF score in a separate goroutine
	ourVRFComputed := false
	if ss.vrfManager != nil {
		ownPeerID, err := peer.IDFromPrivateKey(ss.vrfManager.privateKey)
		if err == nil && onlineValidators[ownPeerID] != nil {
			ourVRFComputed = true
			go func() {
				// Compute our own VRF output
				vrfOutput, err := ss.vrfManager.ComputeVRF(seed)
				if err != nil {
					errChan <- err
					return
				}
				
				validator := onlineValidators[ownPeerID]
				score := ComputeScore(
					vrfOutput.Value,
					float64(validator.Stake),
					validator.Performance,
					ss.alpha,
					ss.beta,
				)
				
				scoresChan <- ValidatorScore{
					ID:        ownPeerID,
					Score:     score,
					Validator: validator,
					VRFOutput: vrfOutput,
				}
			}()
		}
	}
	
	// For other validators, we'll use precalculated VRF outputs or simulate
	// In a real system, these would be received from the network
	remaining := len(onlineValidators)
	if ourVRFComputed {
		remaining--
	}
	
	for id, validator := range onlineValidators {
		// Skip our own validator if we already started computing it
		if ourVRFComputed && ss.vrfManager != nil {
			ownPeerID, _ := peer.IDFromPrivateKey(ss.vrfManager.privateKey)
			if id == ownPeerID {
				continue
			}
		}
		
		go func(id peer.ID, validator *Validator) {
			// In a real system, these would be VRF outputs from the validators
			// Here we're simulating them for demonstration purposes
			
			// Create a deterministic seed for this validator
			hasher := sha256.New()
			hasher.Write(seed)
			hasher.Write([]byte(id.String()))
			deterministicSeed := hasher.Sum(nil)
			
			// Generate a simulated VRF output
			pubKey, err := id.ExtractPublicKey()
			if err != nil {
				errChan <- fmt.Errorf("failed to extract public key for %s: %w", id, err)
				return
			}
			
			// Simulate VRF computation (this is deterministic for the same validator and seed)
			vrfOutput, err := simulateVRFOutput(pubKey, deterministicSeed)
			if err != nil {
				errChan <- err
				return
			}
			
			// Compute score
			score := ComputeScore(
				vrfOutput.Value,
				float64(validator.Stake),
				validator.Performance,
				ss.alpha,
				ss.beta,
			)
			
			scoresChan <- ValidatorScore{
				ID:        id,
				Score:     score,
				Validator: validator,
				VRFOutput: vrfOutput,
			}
		}(id, validator)
	}
	
	// Collect results
	scores := make([]ValidatorScore, 0, len(onlineValidators))
	collected := 0
	
	for collected < len(onlineValidators) {
		select {
		case score := <-scoresChan:
			scores = append(scores, score)
			collected++
		case err := <-errChan:
			ss.logger.Error("Error computing VRF score", "error", err)
			collected++
		case <-ctx.Done():
			return nil, fmt.Errorf("selection timed out")
		}
	}
	
	// Sort by score (lowest first)
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score < scores[j].Score
	})
	
	// Select top N validators
	if numSelected > len(scores) {
		numSelected = len(scores)
	}
	
	selected := make([]peer.ID, numSelected)
	for i := 0; i < numSelected; i++ {
		selected[i] = scores[i].ID
	}
	
	// Create round ID from hash of seed
	roundHash := sha256.Sum256(seed)
	roundID := fmt.Sprintf("%x", roundHash[:8])
	
	// Create selection result
	result := &SelectionResult{
		SelectedValidators: selected,
		Seed:               seed,
		Timestamp:          time.Now(),
		RoundID:            roundID,
		ScoreData:          scores,
	}
	
	// Store the last selection
	ss.mu.Lock()
	ss.lastSelection = result
	ss.mu.Unlock()
	
	ss.logger.Info("Selected validator subset",
		"round_id", roundID,
		"selected_count", len(selected),
		"online_count", len(onlineValidators),
		"gamma", ss.gamma)
	
	return result, nil
}

// GetLastSelection returns the most recent selection result
func (ss *SubsetSelector) GetLastSelection() *SelectionResult {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	return ss.lastSelection
}

// IsSelectedValidator checks if the given validator ID was selected in the most recent round
func (ss *SubsetSelector) IsSelectedValidator(id peer.ID) bool {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	
	if ss.lastSelection == nil {
		return false
	}
	
	for _, selectedID := range ss.lastSelection.SelectedValidators {
		if selectedID == id {
			return true
		}
	}
	
	return false
}

// GetSelectedValidators returns the list of validators selected in the most recent round
func (ss *SubsetSelector) GetSelectedValidators() []peer.ID {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	
	if ss.lastSelection == nil {
		return nil
	}
	
	// Return a copy to prevent external modification
	result := make([]peer.ID, len(ss.lastSelection.SelectedValidators))
	copy(result, ss.lastSelection.SelectedValidators)
	
	return result
}

// Simulate VRF output (for testing/demo purposes)
func simulateVRFOutput(pubKey crypto.PubKey, seed []byte) (*VRFOutput, error) {
	// In a real system, validators would compute their VRF outputs and share them
	// Here we're creating a deterministic but unpredictable value based on the public key and seed
	
	// Combine public key and seed
	pubKeyBytes, err := pubKey.Raw()
	if err != nil {
		return nil, fmt.Errorf("failed to serialize public key: %w", err)
	}
	
	hasher := sha256.New()
	hasher.Write(pubKeyBytes)
	hasher.Write(seed)
	hash := hasher.Sum(nil)
	
	// First 8 bytes for the VRF value
	value := float64(0)
	for i := 0; i < 8 && i < len(hash); i++ {
		value += float64(hash[i]) / (256.0 * math.Pow(256.0, float64(i)))
	}
	
	return &VRFOutput{
		Output:    hash,
		Proof:     hash, // In a real system, this would be a ZK proof
		Value:     value,
		Timestamp: time.Now(),
	}, nil
}