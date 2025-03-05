package creditscore

import (
	"math"
	"sync"
	"time"
	"github.com/IWALINK/cosine/internal/utils"
)

// CreditUpdateType identifies the type of credit score update
type CreditUpdateType string

// Credit update types
const (
	NegativeVoteUpdate  CreditUpdateType = "negative_vote"
	PositiveVoteUpdate  CreditUpdateType = "positive_vote"
	AssociationUpdate   CreditUpdateType = "association"
	RehabilitationUpdate CreditUpdateType = "rehabilitation"
)

// CreditUpdateEntry represents a single credit score update
type CreditUpdateEntry struct {
	WalletAddress  string          `json:"wallet_address"`
	UpdateType     CreditUpdateType `json:"update_type"`
	Delta          float64         `json:"delta"`
	ExpectedDelta  float64         `json:"expected_delta"` // For scaling factor feedback
	Timestamp      time.Time       `json:"timestamp"`
	ValidatorCount int             `json:"validator_count"`
	InlierCount    int             `json:"inlier_count"`
	OutlierCount   int             `json:"outlier_count"`
}

// CreditUpdateResult contains the result of a credit score update
type CreditUpdateResult struct {
	WalletAddress    string          `json:"wallet_address"`
	UpdateType       CreditUpdateType `json:"update_type"`
	OldScore         float64         `json:"old_score"`
	NewScore         float64         `json:"new_score"`
	Delta            float64         `json:"delta"`
	FilterResult     OutlierFilterResult `json:"filter_result"`
	ScalingFactor    float64         `json:"scaling_factor"`
	AccumulatedCost  float64         `json:"accumulated_cost"`
	Timestamp        time.Time       `json:"timestamp"`
}

// CreditUpdateManager manages credit score updates
type CreditUpdateManager struct {
	scoreManager   *CreditScoreManager
	outlierFilter  *OutlierFilter
	dynamicScaling *DynamicScaling
	kalmanManager  *KalmanManager
	config         *utils.ConfigManager
	logger         *utils.Logger
	metrics        *utils.MetricsManager
	mu             sync.RWMutex
	updateHistory  []CreditUpdateEntry
}

// NewCreditUpdateManager creates a new credit update manager
func NewCreditUpdateManager(scoreManager *CreditScoreManager, outlierFilter *OutlierFilter, 
	dynamicScaling *DynamicScaling, kalmanManager *KalmanManager, 
	config *utils.ConfigManager, logger *utils.Logger, metrics *utils.MetricsManager) *CreditUpdateManager {
	
	return &CreditUpdateManager{
		scoreManager:   scoreManager,
		outlierFilter:  outlierFilter,
		dynamicScaling: dynamicScaling,
		kalmanManager:  kalmanManager,
		config:         config,
		logger:         logger.WithComponent("CreditUpdateManager"),
		metrics:        metrics,
		updateHistory:  make([]CreditUpdateEntry, 0, 1000), // Cap history at 1000 entries
	}
}

// ProcessVoteUpdates processes validator proposals for vote-based credit score updates
func (cum *CreditUpdateManager) ProcessVoteUpdates(walletAddress string, proposals []ProposalEntry, 
	updateType CreditUpdateType, gamma float64) CreditUpdateResult {
	
	// Get current wallet credit score
	oldScore := cum.scoreManager.GetCreditScore(walletAddress)
	
	// Step 1: Apply hybrid outlier filtering to the proposals
	filterResult := cum.outlierFilter.FilterOutliers(proposals)
	
	// Step 2: Get the effective vote value from the aggregated score
	// This will be compared to the mean of existing votes to determine shift direction and magnitude
	effectiveVote := filterResult.AggregatedScore
	
	// Step 3: Get vote statistics for this wallet
	voteStatMean, _ := cum.scoreManager.GetRollingStatistic(walletAddress, "μ_V")
	voteStatStdDev, _ := cum.scoreManager.GetRollingStatistic(walletAddress, "σ_V")
	
	// Avoid division by zero
	if voteStatStdDev == 0 {
		voteStatStdDev = 1.0
	}
	
	// Step 4: Calculate the deviation from the mean
	deltaV := effectiveVote - voteStatMean
	
	// Step 5: Determine if update should proceed based on lambda threshold
	lambda := 2.5 // Standard deviations from mean to trigger an update
	if cum.config != nil {
		lambda = cum.config.GetFloat64("voting.lambda")
	}
	
	var delta float64 = 0
	var scalingType ScalingFactorType
	
	// Only proceed if deviation exceeds threshold
	if math.Abs(deltaV) > lambda*voteStatStdDev {
		// Step 6: Calculate credit score shift using dynamic scaling
		if updateType == NegativeVoteUpdate || deltaV < 0 {
			scalingType = NegativeVoteScaling
			delta = cum.dynamicScaling.CalculateVoteShift(deltaV, voteStatStdDev, gamma)
		} else {
			scalingType = PositiveVoteScaling
			delta = cum.dynamicScaling.CalculateVoteShift(deltaV, voteStatStdDev, gamma)
		}
		
		// Step 7: Update the wallet's credit score
		cum.scoreManager.UpdateCreditScore(walletAddress, delta)
		
		// Step 8: Update the scaling factor based on the actual credit score shift
		expectedDelta := delta // For initial feedback, expected = actual
		cum.dynamicScaling.UpdateScalingFactorFromShift(scalingType, delta, expectedDelta)
		
		// Step 9: Update the rolling vote statistics
		alpha := 0.1 // EMA weight for vote statistics
		cum.scoreManager.UpdateRollingStatistic(walletAddress, "μ_V", effectiveVote, alpha)
		newDev := math.Abs(deltaV)
		cum.scoreManager.UpdateRollingStatistic(walletAddress, "σ_V", newDev, alpha)
		
		// Record the update for history
		cum.addUpdateToHistory(CreditUpdateEntry{
			WalletAddress:  walletAddress,
			UpdateType:     updateType,
			Delta:          delta,
			ExpectedDelta:  expectedDelta,
			Timestamp:      time.Now(),
			ValidatorCount: len(proposals),
			InlierCount:    filterResult.InlierCount,
			OutlierCount:   filterResult.OutlierCount,
		})
	} else {
		cum.logger.Debug("Vote update skipped: deviation below threshold", 
			"wallet", walletAddress, 
			"delta_v", deltaV, 
			"threshold", lambda*voteStatStdDev)
	}
	
	// Calculate accumulated cost for this update
	// Typically this would be based on validator rewards
	rewardPerValidator := 0.01 // Example value
	if cum.config != nil {
		rewardPerValidator = cum.config.GetFloat64("rewards.per_validator")
	}
	accumulatedCost := cum.scoreManager.AccumulateCost(walletAddress, float64(filterResult.InlierCount)*rewardPerValidator)
	
	newScore := cum.scoreManager.GetCreditScore(walletAddress)
	
	return CreditUpdateResult{
		WalletAddress:   walletAddress,
		UpdateType:      updateType,
		OldScore:        oldScore,
		NewScore:        newScore,
		Delta:           delta,
		FilterResult:    filterResult,
		ScalingFactor:   cum.dynamicScaling.GetScalingFactor(scalingType),
		AccumulatedCost: accumulatedCost,
		Timestamp:       time.Now(),
	}
}

// ProcessAssociationUpdate processes association risk-based credit score updates
func (cum *CreditUpdateManager) ProcessAssociationUpdate(walletAddress string, riskScore float64, gamma float64) CreditUpdateResult {
	// Get current wallet credit score
	oldScore := cum.scoreManager.GetCreditScore(walletAddress)
	
	// Get association risk statistics
	riskStatMean, _ := cum.scoreManager.GetRollingStatistic(walletAddress, "μ_R")
	riskStatStdDev, _ := cum.scoreManager.GetRollingStatistic(walletAddress, "σ_R")
	
	// Calculate credit score shift using dynamic scaling
	delta := cum.dynamicScaling.CalculateAssociationShift(riskScore, riskStatMean, riskStatStdDev, gamma)
	
	if delta != 0 {
		// Update the wallet's credit score
		cum.scoreManager.UpdateCreditScore(walletAddress, delta)
		
		// Update the wallet's association risk score
		cum.scoreManager.UpdateAssociationRiskScore(walletAddress, riskScore)
		
		// Update the rolling risk statistics
		alpha := 0.1 // EMA weight for risk statistics
		cum.scoreManager.UpdateRollingStatistic(walletAddress, "μ_R", riskScore, alpha)
		newDev := math.Abs(riskScore - riskStatMean)
		cum.scoreManager.UpdateRollingStatistic(walletAddress, "σ_R", newDev, alpha)
		
		// Update the association scaling factor
		expectedDelta := delta // For initial feedback
		cum.dynamicScaling.UpdateScalingFactorFromShift(AssociationScaling, delta, expectedDelta)
		
		// Record the update for history
		cum.addUpdateToHistory(CreditUpdateEntry{
			WalletAddress: walletAddress,
			UpdateType:    AssociationUpdate,
			Delta:         delta,
			ExpectedDelta: expectedDelta,
			Timestamp:     time.Now(),
		})
	}
	
	// Calculate accumulated cost for this update
	rewardForAssociation := 0.05 // Example value
	if cum.config != nil {
		rewardForAssociation = cum.config.GetFloat64("rewards.association")
	}
	accumulatedCost := cum.scoreManager.AccumulateCost(walletAddress, rewardForAssociation)
	
	newScore := cum.scoreManager.GetCreditScore(walletAddress)
	
	// Create an empty filter result as association updates don't use outlier filtering
	emptyFilterResult := OutlierFilterResult{
		InlierCount:  0,
		OutlierCount: 0,
	}
	
	return CreditUpdateResult{
		WalletAddress:   walletAddress,
		UpdateType:      AssociationUpdate,
		OldScore:        oldScore,
		NewScore:        newScore,
		Delta:           delta,
		FilterResult:    emptyFilterResult,
		ScalingFactor:   cum.dynamicScaling.GetScalingFactor(AssociationScaling),
		AccumulatedCost: accumulatedCost,
		Timestamp:       time.Now(),
	}
}

// ProcessRehabilitationUpdate processes rehabilitation-based credit score updates
func (cum *CreditUpdateManager) ProcessRehabilitationUpdate(walletAddress string, rehabVote float64, gamma float64) CreditUpdateResult {
	// Get current wallet credit score
	oldScore := cum.scoreManager.GetCreditScore(walletAddress)
	
	// Get rehab statistics (can use same stats as vote statistics)
	rehabStatMean, _ := cum.scoreManager.GetRollingStatistic(walletAddress, "μ_V")
	rehabStatStdDev, _ := cum.scoreManager.GetRollingStatistic(walletAddress, "σ_V")
	
	// Calculate credit score shift using dynamic scaling
	delta := cum.dynamicScaling.CalculateRehabShift(rehabVote, rehabStatMean, rehabStatStdDev, gamma)
	
	if delta != 0 {
		// Update the wallet's credit score
		cum.scoreManager.UpdateCreditScore(walletAddress, delta)
		
		// Update the rolling statistics
		alpha := 0.1 // EMA weight for rehab statistics
		cum.scoreManager.UpdateRollingStatistic(walletAddress, "μ_V", rehabVote, alpha)
		newDev := math.Abs(rehabVote - rehabStatMean)
		cum.scoreManager.UpdateRollingStatistic(walletAddress, "σ_V", newDev, alpha)
		
		// Update the rehab scaling factor
		expectedDelta := delta // For initial feedback
		cum.dynamicScaling.UpdateScalingFactorFromShift(RehabilitationScaling, delta, expectedDelta)
		
		// Record the update for history
		cum.addUpdateToHistory(CreditUpdateEntry{
			WalletAddress: walletAddress,
			UpdateType:    RehabilitationUpdate,
			Delta:         delta,
			ExpectedDelta: expectedDelta,
			Timestamp:     time.Now(),
		})
	}
	
	// Calculate accumulated cost for this update
	rewardForRehab := 0.03 // Example value
	if cum.config != nil {
		rewardForRehab = cum.config.GetFloat64("rewards.rehabilitation")
	}
	accumulatedCost := cum.scoreManager.AccumulateCost(walletAddress, rewardForRehab)
	
	newScore := cum.scoreManager.GetCreditScore(walletAddress)
	
	// Create an empty filter result as rehab updates don't use outlier filtering
	emptyFilterResult := OutlierFilterResult{
		InlierCount:  0,
		OutlierCount: 0,
	}
	
	return CreditUpdateResult{
		WalletAddress:   walletAddress,
		UpdateType:      RehabilitationUpdate,
		OldScore:        oldScore,
		NewScore:        newScore,
		Delta:           delta,
		FilterResult:    emptyFilterResult,
		ScalingFactor:   cum.dynamicScaling.GetScalingFactor(RehabilitationScaling),
		AccumulatedCost: accumulatedCost,
		Timestamp:       time.Now(),
	}
}

// ProcessUnifiedUpdate combines multiple update types into a single update
func (cum *CreditUpdateManager) ProcessUnifiedUpdate(walletAddress string, 
	voteProposals []ProposalEntry, voteType CreditUpdateType,
	associationRisk float64, rehabVote float64, gamma float64) CreditUpdateResult {
	
	// Get current wallet credit score
	oldScore := cum.scoreManager.GetCreditScore(walletAddress)
	
	// Process each update type and accumulate deltas
	var totalDelta float64
	var filterResult OutlierFilterResult
	
	// Process vote update if proposals exist
	if len(voteProposals) > 0 {
		voteResult := cum.ProcessVoteUpdates(walletAddress, voteProposals, voteType, gamma)
		totalDelta += voteResult.Delta
		filterResult = voteResult.FilterResult
	}
	
	// Process association update if risk score is non-zero
	if associationRisk > 0 {
		assocResult := cum.ProcessAssociationUpdate(walletAddress, associationRisk, gamma)
		totalDelta += assocResult.Delta
	}
	
	// Process rehabilitation update if rehab vote is non-zero
	if rehabVote > 0 {
		rehabResult := cum.ProcessRehabilitationUpdate(walletAddress, rehabVote, gamma)
		totalDelta += rehabResult.Delta
	}
	
	newScore := cum.scoreManager.GetCreditScore(walletAddress)
	accumulatedCost := cum.scoreManager.GetAccumulatedCost(walletAddress)
	
	return CreditUpdateResult{
		WalletAddress:   walletAddress,
		UpdateType:      "unified",
		OldScore:        oldScore,
		NewScore:        newScore,
		Delta:           totalDelta,
		FilterResult:    filterResult,
		AccumulatedCost: accumulatedCost,
		Timestamp:       time.Now(),
	}
}

// GetVerificationFee calculates the fee for verifying a wallet's credit score
func (cum *CreditUpdateManager) GetVerificationFee(walletAddress string) float64 {
	accumulatedCost := cum.scoreManager.GetAccumulatedCost(walletAddress)
	
	// Apply fee multiplier
	kappa := 1.2 // Default multiplier
	if cum.config != nil {
		kappa = cum.config.GetFloat64("fees.verification.kappa")
	}
	
	return kappa * accumulatedCost
}

// ProcessVerification handles the verification fee payment and resets accumulated cost
func (cum *CreditUpdateManager) ProcessVerification(walletAddress string) float64 {
	fee := cum.GetVerificationFee(walletAddress)
	
	// Reset accumulated cost
	cum.scoreManager.ResetAccumulatedCost(walletAddress)
	
	cum.logger.Info("Processed verification", 
		"wallet", walletAddress, 
		"fee", fee)
	
	return fee
}

// AddUpdateToHistory adds a credit update entry to history
func (cum *CreditUpdateManager) addUpdateToHistory(entry CreditUpdateEntry) {
	cum.mu.Lock()
	defer cum.mu.Unlock()
	
	// Add to history, maintaining max size
	if len(cum.updateHistory) >= 1000 {
		// Remove oldest entry
		cum.updateHistory = cum.updateHistory[1:]
	}
	
	cum.updateHistory = append(cum.updateHistory, entry)
}

// GetUpdateHistory returns credit update history for a wallet
func (cum *CreditUpdateManager) GetUpdateHistory(walletAddress string) []CreditUpdateEntry {
	cum.mu.RLock()
	defer cum.mu.RUnlock()
	
	// Filter history for specified wallet
	result := make([]CreditUpdateEntry, 0)
	for _, entry := range cum.updateHistory {
		if entry.WalletAddress == walletAddress {
			result = append(result, entry)
		}
	}
	
	return result
}

// GetCosineSimilarity calculates cosine similarity between wallet's score vector and threshold/target vector
func (cum *CreditUpdateManager) GetCosineSimilarity(walletAddress string, threshold float64) float64 {
	// Get wallet's normalized vector
	walletVector := cum.scoreManager.GetNormalizedCreditVector(walletAddress)
	
	// Create threshold vector
	metrics := cum.scoreManager.GetCreditMetrics()
	normalizedThreshold := (threshold - metrics.Mean) / metrics.StdDev
	thresholdVector := utils.VectorFromValues(normalizedThreshold, 1)
	
	// Calculate cosine similarity
	return utils.CosineSimilarity(walletVector, thresholdVector)
}

// GetWalletCosineSimilarity calculates cosine similarity between two wallets
func (cum *CreditUpdateManager) GetWalletCosineSimilarity(wallet1Address, wallet2Address string) float64 {
	// Get both wallet vectors
	wallet1Vector := cum.scoreManager.GetNormalizedCreditVector(wallet1Address)
	wallet2Vector := cum.scoreManager.GetNormalizedCreditVector(wallet2Address)
	
	// Calculate cosine similarity
	return utils.CosineSimilarity(wallet1Vector, wallet2Vector)
}

// IsAboveThreshold determines if a wallet's score meets or exceeds a threshold with cosine similarity
func (cum *CreditUpdateManager) IsAboveThreshold(walletAddress string, threshold float64, similarityThreshold float64) bool {
	similarity := cum.GetCosineSimilarity(walletAddress, threshold)
	return similarity >= similarityThreshold
}

// GetCreditScoreManager returns the credit score manager
func (cum *CreditUpdateManager) GetCreditScoreManager() *CreditScoreManager {
	return cum.scoreManager
}