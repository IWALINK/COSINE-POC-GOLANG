package fraud

import (
	"math"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/creditscore"
	"github.com/IWALINK/cosine/internal/utils"
)

// BlacklistManager manages partial blacklisting of wallets
// based on association risk and prevents blacklisted wallets from
// participating in votes for wallets they're associated with
type BlacklistManager struct {
	associationAnalyzer *AssociationRiskAnalyzer
	creditUpdater       *creditscore.CreditUpdateManager
	blacklistedWallets  map[string]*BlacklistEntry
	voteRestrictions    map[string][]string // Mapping from wallet to wallets it cannot vote on
	config              *utils.ConfigManager
	logger              *utils.Logger
	metrics             *utils.MetricsManager
	mu                  sync.RWMutex
}

// BlacklistEntry represents a wallet that has been partially blacklisted
type BlacklistEntry struct {
	WalletAddress     string                 `json:"wallet_address"`
	BlacklistTimestamp time.Time             `json:"blacklist_timestamp"`
	AssociationScore   float64               `json:"association_score"`
	Reason             string                `json:"reason"`
	AssociatedWith     []string              `json:"associated_with"`
	Paths              []AssociationPath     `json:"paths"`
	CreditPenalty      float64               `json:"credit_penalty"`
	Status             BlacklistStatus       `json:"status"`
}

// BlacklistStatus represents the current status of a blacklisted wallet
type BlacklistStatus string

// Blacklist status constants
const (
	PartialBlacklist BlacklistStatus = "partial_blacklist"
	FullBlacklist    BlacklistStatus = "full_blacklist"
	Rehabilitated    BlacklistStatus = "rehabilitated"
)

// NewBlacklistManager creates a new blacklist manager
func NewBlacklistManager(
	associationAnalyzer *AssociationRiskAnalyzer, 
	creditUpdater *creditscore.CreditUpdateManager,
	config *utils.ConfigManager,
	logger *utils.Logger,
	metrics *utils.MetricsManager,
) *BlacklistManager {
	return &BlacklistManager{
		associationAnalyzer: associationAnalyzer,
		creditUpdater:       creditUpdater,
		blacklistedWallets:  make(map[string]*BlacklistEntry),
		voteRestrictions:    make(map[string][]string),
		config:              config,
		logger:              logger.WithComponent("BlacklistManager"),
		metrics:             metrics,
	}
}

// ProcessAssociationPenalty processes an association penalty for a wallet
// This is triggered when a wallet is found to have associations with malicious wallets
func (bm *BlacklistManager) ProcessAssociationPenalty(
	walletAddress string, 
	triggerTimestamp time.Time, 
	gamma float64,
) (float64, *BlacklistEntry) {
	// Get association risk score and paths
	riskScore, paths := bm.associationAnalyzer.ProcessAssociationAnalysis(
		walletAddress, 
		triggerTimestamp,
	)
	
	// If no risk, no penalty needed
	if riskScore <= 0 || len(paths) == 0 {
		return 0, nil
	}
	
	// Get association risk stats (mean, stdDev) for standardized deviation calculation
	mean, stdDev := bm.associationAnalyzer.GetAssociationRiskStats()
	
	// Process credit score update through dynamic scaling
	result := bm.creditUpdater.ProcessAssociationUpdate(walletAddress, riskScore, gamma)
	
	// Create blacklist entry
	entry := &BlacklistEntry{
		WalletAddress:     walletAddress,
		BlacklistTimestamp: time.Now(),
		AssociationScore:   riskScore,
		Reason:             "Association with malicious wallets",
		CreditPenalty:      result.Delta,
		Status:             PartialBlacklist,
	}
	
	// Extract associated malicious wallets from paths
	associatedWith := make(map[string]bool)
	for _, path := range paths {
		if len(path.Transactions) > 0 {
			// First transaction's from address is the potentially malicious wallet
			maliciousWallet := path.Transactions[0].FromAddress
			associatedWith[maliciousWallet] = true
			
			// Add to vote restrictions
			bm.addVoteRestriction(walletAddress, maliciousWallet)
		}
	}
	
	// Convert map to slice for storage
	for wallet := range associatedWith {
		entry.AssociatedWith = append(entry.AssociatedWith, wallet)
	}
	
	// Store paths for reference
	entry.Paths = paths
	
	// Add to blacklisted wallets
	bm.mu.Lock()
	bm.blacklistedWallets[walletAddress] = entry
	bm.mu.Unlock()
	
	// Log the penalty
	bm.logger.Info("Applied association penalty",
		"wallet", walletAddress,
		"risk_score", riskScore,
		"standardized_risk", (riskScore-mean)/stdDev,
		"credit_penalty", result.Delta,
		"associated_wallets", len(entry.AssociatedWith))
	
	return result.Delta, entry
}

// IsExchangeWallet determines if a wallet is likely an exchange wallet
// based on transaction patterns, to protect exchanges from false penalties
func (bm *BlacklistManager) IsExchangeWallet(walletAddress string) bool {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	
	// This is a simplified check - in a real implementation, this would be more sophisticated
	// and would analyze transaction volumes, frequencies, number of unique counterparties, etc.
	
	// Example criteria (simplified):
	// 1. High transaction volume
	// 2. Many unique counterparties
	// 3. Frequent incoming/outgoing transactions
	// 4. Consistent transaction pattern
	
	// For this MVP implementation, we'll use a simple check based on transaction count
	txCount := 0
	if txs, ok := bm.associationAnalyzer.transactions[walletAddress]; ok {
		txCount = len(txs)
	}
	
	// If transaction count exceeds a threshold, consider it an exchange wallet
	exchangeThreshold := 1000 // Example threshold
	return txCount > exchangeThreshold
}

// CanVoteOn checks if a wallet is allowed to vote on another wallet
// Returns true if allowed, false if restricted
func (bm *BlacklistManager) CanVoteOn(voterWallet, targetWallet string) bool {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	
	// Check if voter has any restrictions
	if restrictions, ok := bm.voteRestrictions[voterWallet]; ok {
		// Check if target wallet is in the restrictions
		for _, restricted := range restrictions {
			if restricted == targetWallet {
				return false
			}
		}
	}
	
	// No restrictions found
	return true
}

// addVoteRestriction adds a vote restriction for a wallet
// The wallet cannot vote on the target wallet
func (bm *BlacklistManager) addVoteRestriction(wallet, targetWallet string) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	
	// Initialize slice if needed
	if _, ok := bm.voteRestrictions[wallet]; !ok {
		bm.voteRestrictions[wallet] = []string{}
	}
	
	// Check if restriction already exists
	for _, existing := range bm.voteRestrictions[wallet] {
		if existing == targetWallet {
			return
		}
	}
	
	// Add restriction
	bm.voteRestrictions[wallet] = append(bm.voteRestrictions[wallet], targetWallet)
	
	bm.logger.Debug("Added vote restriction",
		"wallet", wallet,
		"cannot_vote_on", targetWallet)
}

// IsPartiallyBlacklisted checks if a wallet is partially blacklisted
func (bm *BlacklistManager) IsPartiallyBlacklisted(walletAddress string) bool {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	
	entry, exists := bm.blacklistedWallets[walletAddress]
	return exists && entry.Status == PartialBlacklist
}

// GetBlacklistEntry returns the blacklist entry for a wallet if it exists
func (bm *BlacklistManager) GetBlacklistEntry(walletAddress string) *BlacklistEntry {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	
	if entry, exists := bm.blacklistedWallets[walletAddress]; exists {
		return entry
	}
	
	return nil
}

// GetAllBlacklistedWallets returns all blacklisted wallets
func (bm *BlacklistManager) GetAllBlacklistedWallets() []*BlacklistEntry {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	
	entries := make([]*BlacklistEntry, 0, len(bm.blacklistedWallets))
	for _, entry := range bm.blacklistedWallets {
		entries = append(entries, entry)
	}
	
	return entries
}

// CalculateEffectiveBlacklistFactor calculates the effective blacklist factor
// for a wallet based on its association penalty and credit score
// This can be used by other systems to adjust their behavior based on
// how strongly blacklisted a wallet is
func (bm *BlacklistManager) CalculateEffectiveBlacklistFactor(walletAddress string) float64 {
	entry := bm.GetBlacklistEntry(walletAddress)
	if entry == nil {
		return 0 // Not blacklisted
	}
	
	// Get current credit score
	currentScore := bm.creditUpdater.GetCreditScoreManager().GetCreditScore(walletAddress)
	
	// Calculate effective factor based on penalty magnitude and current score
	// Higher penalty and lower score means stronger blacklisting effect
	penaltyMagnitude := math.Abs(entry.CreditPenalty)
	
	// If score is negative, the effect is stronger
	if currentScore < 0 {
		return math.Min(1.0, penaltyMagnitude/100.0 * (1.0 - currentScore/1000.0))
	}
	
	// If score is positive but there was a penalty, effect is moderate
	return math.Min(0.5, penaltyMagnitude/200.0)
}

// UpdateBlacklistStatus updates the status of a blacklisted wallet
// This can be called when a wallet's situation changes (rehabilitation or full blacklisting)
func (bm *BlacklistManager) UpdateBlacklistStatus(walletAddress string, status BlacklistStatus) bool {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	
	entry, exists := bm.blacklistedWallets[walletAddress]
	if !exists {
		return false
	}
	
	entry.Status = status
	
	bm.logger.Info("Updated blacklist status",
		"wallet", walletAddress,
		"status", status)
	
	return true
}

// GetVoteRestrictions returns all wallets that a given wallet cannot vote on
func (bm *BlacklistManager) GetVoteRestrictions(walletAddress string) []string {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	
	if restrictions, ok := bm.voteRestrictions[walletAddress]; ok {
		return restrictions
	}
	
	return []string{}
}