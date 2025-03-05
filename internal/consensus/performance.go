package consensus

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
	"github.com/libp2p/go-libp2p/core/peer"
)

// OutlierStatus represents whether a proposal is an inlier or outlier
type OutlierStatus bool

const (
	// StatusInlier indicates the proposal is within acceptable bounds
	StatusInlier OutlierStatus = true
	
	// StatusOutlier indicates the proposal is outside acceptable bounds
	StatusOutlier OutlierStatus = false
)

// ProposalData represents a credit score proposal from a validator
type ProposalData struct {
	// Validator ID that made the proposal
	ValidatorID peer.ID
	
	// Proposed credit score value
	ProposedValue float64
	
	// Previous credit score value
	PreviousValue float64
	
	// Proposed change (delta)
	Delta float64
	
	// Timestamp of proposal
	Timestamp time.Time
	
	// Whether this proposal is an inlier or outlier
	Status OutlierStatus
}

// AggregationResult represents the result of proposal aggregation
type AggregationResult struct {
	// Final aggregated score
	AggregatedValue float64
	
	// Mean of all proposals
	Mean float64
	
	// Standard deviation of all proposals
	StdDev float64
	
	// Median of all proposals
	Median float64
	
	// Median Absolute Deviation (normalized)
	MAD float64
	
	// Tau threshold parameter (mean-based)
	Tau float64
	
	// K threshold parameter (median-based)
	K float64
	
	// Total inlier proposals
	InlierCount int
	
	// Total outlier proposals
	OutlierCount int
	
	// All proposals with their status
	Proposals []ProposalData
	
	// Target wallet address
	TargetWallet string
	
	// Timestamp of aggregation
	Timestamp time.Time
}

// PerformanceManager handles validator performance scoring and credit score aggregation
type PerformanceManager struct {
	validatorSet    *ValidatorSet
	selector        *SubsetSelector
	logger          *utils.Logger
	config          *utils.ConfigManager
	gso             *storage.GlobalStateObject
	deltaPlus       float64
	deltaMinus      float64
	rewardPool      uint64
	mu              sync.RWMutex
	lastAggregation *AggregationResult
	// Kalman filter state for threshold parameters
	tauKalman struct {
		tau float64 // Current tau value
		p   float64 // Error variance
	}
	kKalman struct {
		k float64 // Current k value
		p float64 // Error variance
	}
}

// NewPerformanceManager creates a new validator performance manager
func NewPerformanceManager(
	validatorSet *ValidatorSet,
	selector *SubsetSelector,
	gso *storage.GlobalStateObject,
	config *utils.ConfigManager,
	logger *utils.Logger,
) *PerformanceManager {
	// Get delta+ parameter (default to 0.01)
	deltaPlus := config.GetFloat64("consensus.deltaPlus")
	if deltaPlus <= 0 {
		deltaPlus = 0.01
	}
	
	// Get delta- parameter (default to 0.05)
	deltaMinus := config.GetFloat64("consensus.deltaMinus")
	if deltaMinus <= 0 {
		deltaMinus = 0.05
	}
	
	// Initialize Kalman filter parameters
	tauInit, kInit, pInit, qTau, qK, rTau, rK := config.GetKalmanFilterParams("consensus.kalman")
	
	// Use defaults if not specified
	if tauInit <= 0 {
		tauInit = 2.5 // Default tau
	}
	if kInit <= 0 {
		kInit = 3.0 // Default k
	}
	if pInit <= 0 {
		pInit = 1.0 // Default initial error variance
	}
	
	pm := &PerformanceManager{
		validatorSet: validatorSet,
		selector:     selector,
		logger:       logger.WithComponent("PerformanceManager"),
		config:       config,
		gso:          gso,
		deltaPlus:    deltaPlus,
		deltaMinus:   deltaMinus,
	}
	
	// Initialize Kalman filter state
	pm.tauKalman.tau = tauInit
	pm.tauKalman.p = pInit
	pm.kKalman.k = kInit
	pm.kKalman.p = pInit
	
	// Log initialization parameters
	logger.Info("Initialized performance manager",
		"tau", tauInit,
		"k", kInit,
		"delta_plus", deltaPlus,
		"delta_minus", deltaMinus,
		"q_tau", qTau,
		"q_k", qK,
		"r_tau", rTau,
		"r_k", rK)
	
	return pm
}

// AggregateProposals aggregates credit score proposals using hybrid outlier filtering
func (pm *PerformanceManager) AggregateProposals(
	targetWallet string,
	proposals []ProposalData,
) (*AggregationResult, error) {
	if len(proposals) == 0 {
		return nil, fmt.Errorf("no proposals to aggregate")
	}
	
	// Extract values for statistical processing
	values := make([]float64, len(proposals))
	for i, prop := range proposals {
		values[i] = prop.ProposedValue
	}
	
	// Calculate mean-based measures
	mean := utils.Mean(values)
	stdDev := utils.StandardDeviation(values)
	
	// Calculate median-based measures
	median := utils.Median(values)
	normalizedMAD := utils.NormalizedMAD(values)
	
	// Get current threshold parameters
	pm.mu.RLock()
	tau := pm.tauKalman.tau
	k := pm.kKalman.k
	pm.mu.RUnlock()
	
	// Apply outlier filtering using both approaches
	inlierCount := 0
	for i := range proposals {
		// Calculate deviations
		meanDev := math.Abs(proposals[i].ProposedValue - mean)
		medianDev := math.Abs(proposals[i].ProposedValue - median)
		
		// Check against both thresholds
		meanInlier := meanDev <= tau*stdDev
		medianInlier := medianDev <= k*normalizedMAD
		
		// Mark as inlier only if it passes both criteria
		proposals[i].Status = StatusOutlier
		if meanInlier && medianInlier {
			proposals[i].Status = StatusInlier
		}
	}
	
	// Compute aggregated value from inliers only
	var aggregatedValue float64
	if inlierCount > 0 {
		sum := 0.0
		for _, prop := range proposals {
			if prop.Status == StatusInlier {
				sum += prop.ProposedValue
			}
		}
		aggregatedValue = sum / float64(inlierCount)
	} else {
		// No inliers, fall back to median
		aggregatedValue = median
		pm.logger.Warn("No inliers found, using median as aggregated value",
			"target_wallet", targetWallet,
			"proposal_count", len(proposals),
			"median", median)
	}
	
	// Create aggregation result
	result := &AggregationResult{
		AggregatedValue: aggregatedValue,
		Mean:            mean,
		StdDev:          stdDev,
		Median:          median,
		MAD:             normalizedMAD,
		Tau:             tau,
		K:               k,
		InlierCount:     inlierCount,
		OutlierCount:    len(proposals) - inlierCount,
		Proposals:       proposals,
		TargetWallet:    targetWallet,
		Timestamp:       time.Now(),
	}
	
	// Update threshold parameters using Kalman filters
	pm.updateThresholdParameters(mean, stdDev, median, normalizedMAD, proposals)
	
	// Store last aggregation result
	pm.mu.Lock()
	pm.lastAggregation = result
	pm.mu.Unlock()
	
	// Update validator performance scores
	pm.updateValidatorPerformance(proposals)
	
	// Distribute rewards to inlier validators
	pm.distributeRewards(proposals)
	
	pm.logger.Info("Aggregated proposals",
		"target_wallet", targetWallet,
		"proposal_count", len(proposals),
		"inlier_count", inlierCount,
		"outlier_count", len(proposals) - inlierCount,
		"aggregated_value", aggregatedValue)
	
	return result, nil
}

// Update threshold parameters using Kalman filters
func (pm *PerformanceManager) updateThresholdParameters(
	mean float64,
	stdDev float64,
	median float64,
	mad float64,
	proposals []ProposalData,
) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	
	// Skip update if there are too few proposals
	if len(proposals) < 3 {
		return
	}
	
	// Calculate maximum normalized deviations as measurements
	maxMeanDev := 0.0
	maxMedianDev := 0.0
	
	for _, prop := range proposals {
		meanDev := math.Abs(prop.ProposedValue - mean) / stdDev
		medianDev := math.Abs(prop.ProposedValue - median) / mad
		
		if meanDev > maxMeanDev {
			maxMeanDev = meanDev
		}
		if medianDev > maxMedianDev {
			maxMedianDev = medianDev
		}
	}
	
	// Get Kalman filter parameters
	_, _, _, qTau, qK, rTau, rK := pm.config.GetKalmanFilterParams("consensus.kalman")
	
	// Prediction step for tau
	pTauPred := pm.tauKalman.p + qTau
	
	// Update step for tau
	gTau := pTauPred / (pTauPred + rTau)
	pm.tauKalman.tau = pm.tauKalman.tau + gTau*(maxMeanDev - pm.tauKalman.tau)
	pm.tauKalman.p = (1 - gTau) * pTauPred
	
	// Prediction step for k
	pKPred := pm.kKalman.p + qK
	
	// Update step for k
	gK := pKPred / (pKPred + rK)
	pm.kKalman.k = pm.kKalman.k + gK*(maxMedianDev - pm.kKalman.k)
	pm.kKalman.p = (1 - gK) * pKPred
	
	// Ensure thresholds are in reasonable range
	if pm.tauKalman.tau < 1.0 {
		pm.tauKalman.tau = 1.0
	} else if pm.tauKalman.tau > 5.0 {
		pm.tauKalman.tau = 5.0
	}
	
	if pm.kKalman.k < 1.0 {
		pm.kKalman.k = 1.0
	} else if pm.kKalman.k > 5.0 {
		pm.kKalman.k = 5.0
	}
	
	pm.logger.Debug("Updated threshold parameters",
		"tau", pm.tauKalman.tau,
		"k", pm.kKalman.k,
		"max_mean_dev", maxMeanDev,
		"max_median_dev", maxMedianDev)
}

// Update validator performance scores based on proposal status
func (pm *PerformanceManager) updateValidatorPerformance(proposals []ProposalData) {
	for _, prop := range proposals {
		var delta float64
		if prop.Status == StatusInlier {
			delta = pm.deltaPlus
		} else {
			delta = -pm.deltaMinus
		}
		
		err := pm.validatorSet.UpdateValidatorPerformance(prop.ValidatorID, delta)
		if err != nil {
			pm.logger.Error("Failed to update validator performance",
				"validator", prop.ValidatorID.String(),
				"error", err)
		}
	}
}

// Distribute rewards to validators with inlier proposals
func (pm *PerformanceManager) distributeRewards(proposals []ProposalData) {
	// Count inliers
	inlierCount := 0
	for _, prop := range proposals {
		if prop.Status == StatusInlier {
			inlierCount++
		}
	}
	
	// Skip if no inliers
	if inlierCount == 0 {
		pm.logger.Warn("No inliers to reward")
		return
	}
	
	// Get reward per inlier
	rewardPerInlier := pm.config.GetFloat64("consensus.rewardPerEvent")
	if rewardPerInlier <= 0 {
		rewardPerInlier = 10.0 // Default 10 tokens per event
	}
	
	// Calculate total reward for this event
	totalReward := uint64(rewardPerInlier * float64(inlierCount))
	
	// Check if enough in pool
	if pm.gso != nil {
		err := pm.gso.DecrementRewardPool(totalReward)
		if err != nil {
			pm.logger.Error("Failed to decrement reward pool", "error", err)
			return
		}
		
		// Update accumulated cost for target wallet in storage
		// This would be handled by the calling code in a real implementation
	}
	
	// Log reward distribution
	pm.logger.Debug("Distributed rewards",
		"inlier_count", inlierCount,
		"reward_per_inlier", rewardPerInlier,
		"total_reward", totalReward)
}

// GetLastAggregation returns the most recent aggregation result
func (pm *PerformanceManager) GetLastAggregation() *AggregationResult {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.lastAggregation
}

// GetThresholdParameters returns the current threshold parameters
func (pm *PerformanceManager) GetThresholdParameters() (tau float64, k float64) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.tauKalman.tau, pm.kKalman.k
}

// IsInlier checks if a proposed value would be considered an inlier
func (pm *PerformanceManager) IsInlier(proposedValue float64, values []float64) bool {
	// Calculate mean-based measures
	mean := utils.Mean(values)
	stdDev := utils.StandardDeviation(values)
	
	// Calculate median-based measures
	median := utils.Median(values)
	normalizedMAD := utils.NormalizedMAD(values)
	
	// Get current threshold parameters
	pm.mu.RLock()
	tau := pm.tauKalman.tau
	k := pm.kKalman.k
	pm.mu.RUnlock()
	
	// Calculate deviations
	meanDev := math.Abs(proposedValue - mean)
	medianDev := math.Abs(proposedValue - median)
	
	// Check against both thresholds
	meanInlier := meanDev <= tau*stdDev
	medianInlier := medianDev <= k*normalizedMAD
	
	// Return true only if it passes both criteria
	return meanInlier && medianInlier
}