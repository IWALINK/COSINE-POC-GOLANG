package creditscore

import (
	"math"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
)

// CreditMetrics holds statistics about credit scores across the network
type CreditMetrics struct {
	Mean      float64   `json:"mean"`
	StdDev    float64   `json:"std_dev"`
	Min       float64   `json:"min"`
	Max       float64   `json:"max"`
	Count     int       `json:"count"`
	Timestamp time.Time `json:"timestamp"`
}

// CreditScoreManager manages wallet credit scores
type CreditScoreManager struct {
	walletTrie    *storage.WalletStateTrie
	dynamicScaling *DynamicScaling
	config        *utils.ConfigManager
	logger        *utils.Logger
	metrics       *utils.MetricsManager
	creditMetrics CreditMetrics
	mu            sync.RWMutex
}

// NewCreditScoreManager creates a new credit score manager
func NewCreditScoreManager(walletTrie *storage.WalletStateTrie, dynamicScaling *DynamicScaling,
	config *utils.ConfigManager, logger *utils.Logger, metrics *utils.MetricsManager) *CreditScoreManager {
	
	csm := &CreditScoreManager{
		walletTrie:    walletTrie,
		dynamicScaling: dynamicScaling,
		config:        config,
		logger:        logger.WithComponent("CreditScoreManager"),
		metrics:       metrics,
		creditMetrics: CreditMetrics{
			Mean:      0,
			StdDev:    1,
			Min:       0,
			Max:       0,
			Count:     0,
			Timestamp: time.Now(),
		},
	}
	
	// Initial calculation of credit metrics
	csm.updateCreditMetrics()
	
	return csm
}

// GetCreditScore returns a wallet's credit score
func (csm *CreditScoreManager) GetCreditScore(walletAddress string) float64 {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	return wallet.CreditScore
}

// GetAssociationRiskScore returns a wallet's association risk score
func (csm *CreditScoreManager) GetAssociationRiskScore(walletAddress string) float64 {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	return wallet.AssociationRiskScore
}

// UpdateCreditScore updates a wallet's credit score with the given delta
func (csm *CreditScoreManager) UpdateCreditScore(walletAddress string, delta float64) float64 {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	wallet.UpdateCreditScore(delta)
	csm.walletTrie.UpdateWallet(wallet)
	
	// Update metrics if available
	if csm.metrics != nil {
		updateType := "other"
		if delta > 0 {
			updateType = "positive"
		} else if delta < 0 {
			updateType = "negative"
		}
		
		csm.metrics.IncCounter("credit_score_updates_total", updateType)
		csm.metrics.ObserveHistogram("credit_score_shift_magnitude", math.Abs(delta), updateType)
	}
	
	csm.logger.Debug("Updated credit score", 
		"wallet", walletAddress, 
		"delta", delta, 
		"new_score", wallet.CreditScore)
	
	// Schedule deferred update of credit metrics (not every single update)
	if csm.shouldUpdateMetrics() {
		go csm.updateCreditMetrics()
	}
	
	return wallet.CreditScore
}

// UpdateAssociationRiskScore updates a wallet's association risk score
func (csm *CreditScoreManager) UpdateAssociationRiskScore(walletAddress string, delta float64) float64 {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	wallet.UpdateAssociationRisk(delta)
	csm.walletTrie.UpdateWallet(wallet)
	
	// Update metrics if available
	if csm.metrics != nil && delta > 0 {
		csm.metrics.IncCounter("association_penalties_total")
		csm.metrics.ObserveHistogram("association_risk_score", wallet.AssociationRiskScore)
	}
	
	csm.logger.Debug("Updated association risk score", 
		"wallet", walletAddress, 
		"delta", delta, 
		"new_score", wallet.AssociationRiskScore)
	
	return wallet.AssociationRiskScore
}

// GetWalletReputationScore returns a wallet's reputation score for voting
func (csm *CreditScoreManager) GetWalletReputationScore(walletAddress string) float64 {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	return wallet.ReputationScore
}

// UpdateWalletReputationScore updates a wallet's reputation score
func (csm *CreditScoreManager) UpdateWalletReputationScore(walletAddress string, delta float64) float64 {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	
	// Clamp reputation between 0 and 1
	newReputation := utils.Clamp(wallet.ReputationScore+delta, 0, 1)
	wallet.ReputationScore = newReputation
	csm.walletTrie.UpdateWallet(wallet)
	
	csm.logger.Debug("Updated reputation score", 
		"wallet", walletAddress, 
		"delta", delta, 
		"new_score", newReputation)
	
	return newReputation
}

// GetRollingStatistic returns a wallet's rolling statistic
func (csm *CreditScoreManager) GetRollingStatistic(walletAddress, statName string) (float64, time.Time) {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	
	if stat, exists := wallet.RollingStats[statName]; exists {
		return stat.Value, stat.LastUpdated
	}
	
	// Default values if not found
	if statName == "μ_V" || statName == "μ_R" {
		return 0, time.Now()
	} else if statName == "σ_V" || statName == "σ_R" {
		return 1, time.Now()
	}
	
	return 0, time.Now()
}

// UpdateRollingStatistic updates a wallet's rolling statistic
func (csm *CreditScoreManager) UpdateRollingStatistic(walletAddress, statName string, newValue float64, alpha float64) {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	wallet.UpdateRollingStatistic(statName, newValue, alpha)
	csm.walletTrie.UpdateWallet(wallet)
	
	csm.logger.Debug("Updated rolling statistic", 
		"wallet", walletAddress, 
		"stat", statName, 
		"new_value", newValue, 
		"alpha", alpha)
}

// GetCreditMetrics returns network-wide credit score metrics
func (csm *CreditScoreManager) GetCreditMetrics() CreditMetrics {
	csm.mu.RLock()
	defer csm.mu.RUnlock()
	return csm.creditMetrics
}

// updateCreditMetrics calculates network-wide credit score metrics
func (csm *CreditScoreManager) updateCreditMetrics() {
	csm.mu.Lock()
	defer csm.mu.Unlock()
	
	wallets := csm.walletTrie.GetAllWallets()
	if len(wallets) == 0 {
		csm.creditMetrics = CreditMetrics{
			Mean:      0,
			StdDev:    1,
			Min:       0,
			Max:       0,
			Count:     0,
			Timestamp: time.Now(),
		}
		return
	}
	
	// Extract credit scores
	scores := make([]float64, len(wallets))
	min := math.MaxFloat64
	max := -math.MaxFloat64
	
	for i, wallet := range wallets {
		score := wallet.CreditScore
		scores[i] = score
		
		if score < min {
			min = score
		}
		if score > max {
			max = score
		}
	}
	
	// Calculate mean and standard deviation
	mean := utils.Mean(scores)
	stdDev := utils.StandardDeviation(scores)
	
	// Ensure non-zero standard deviation
	if stdDev == 0 {
		stdDev = 1
	}
	
	csm.creditMetrics = CreditMetrics{
		Mean:      mean,
		StdDev:    stdDev,
		Min:       min,
		Max:       max,
		Count:     len(wallets),
		Timestamp: time.Now(),
	}
	
	csm.logger.Info("Updated credit metrics", 
		"count", csm.creditMetrics.Count, 
		"mean", csm.creditMetrics.Mean, 
		"std_dev", csm.creditMetrics.StdDev,
		"min", csm.creditMetrics.Min,
		"max", csm.creditMetrics.Max)
}

// shouldUpdateMetrics determines if metrics should be updated
// This avoids updating too frequently
func (csm *CreditScoreManager) shouldUpdateMetrics() bool {
	csm.mu.RLock()
	defer csm.mu.RUnlock()
	
	// Update if more than 5 minutes has passed or no metrics calculated yet
	if csm.creditMetrics.Count == 0 || time.Since(csm.creditMetrics.Timestamp) > 5*time.Minute {
		return true
	}
	
	return false
}

// NormalizeScore normalizes a raw credit score using network-wide metrics
func (csm *CreditScoreManager) NormalizeScore(score float64) float64 {
	csm.mu.RLock()
	defer csm.mu.RUnlock()
	
	return (score - csm.creditMetrics.Mean) / csm.creditMetrics.StdDev
}

// GetNormalizedCreditVector returns a 2D vector for cosine similarity calculations
func (csm *CreditScoreManager) GetNormalizedCreditVector(walletAddress string) utils.Vector {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	
	csm.mu.RLock()
	mean := csm.creditMetrics.Mean
	stdDev := csm.creditMetrics.StdDev
	csm.mu.RUnlock()
	
	return wallet.GetNormalizedCreditVector(mean, stdDev)
}

// GetCreditScoreVerificationResult returns whether a wallet's credit score meets or exceeds a threshold
func (csm *CreditScoreManager) GetCreditScoreVerificationResult(walletAddress string, threshold float64) (bool, float64) {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	
	// Direct comparison with threshold
	if wallet.CreditScore >= threshold {
		return true, wallet.CreditScore
	}
	
	return false, wallet.CreditScore
}

// AccumulateCost increments the accumulated cost of computation for a wallet
func (csm *CreditScoreManager) AccumulateCost(walletAddress string, amount float64) float64 {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	wallet.IncrementAccumulatedCost(amount)
	csm.walletTrie.UpdateWallet(wallet)
	
	return wallet.AccumulatedCost
}

// GetAccumulatedCost returns a wallet's accumulated cost
func (csm *CreditScoreManager) GetAccumulatedCost(walletAddress string) float64 {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	return wallet.AccumulatedCost
}

// ResetAccumulatedCost resets a wallet's accumulated cost (after verification fee payment)
func (csm *CreditScoreManager) ResetAccumulatedCost(walletAddress string) {
	wallet := csm.walletTrie.GetWallet(walletAddress)
	wallet.ResetAccumulatedCost()
	csm.walletTrie.UpdateWallet(wallet)
	
	csm.logger.Debug("Reset accumulated cost", "wallet", walletAddress)
}