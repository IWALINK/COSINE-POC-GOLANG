package fraud

import (
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
)

// Transaction represents a single transaction between wallets
type Transaction struct {
	FromAddress      string    `json:"from_address"`      // Sender wallet address
	ToAddress        string    `json:"to_address"`        // Receiver wallet address
	Amount           float64   `json:"amount"`            // Transaction amount
	Timestamp        time.Time `json:"timestamp"`         // Transaction timestamp
	ClusteringCount  int       `json:"clustering_count"`  // Number of transactions in a short interval
	TxHash           string    `json:"tx_hash"`           // Transaction hash for unique identification
	ConfirmationTime time.Time `json:"confirmation_time"` // When the transaction was confirmed
}

// AssociationPath represents a path of transactions from a malicious wallet to a target wallet
type AssociationPath struct {
	Transactions []Transaction `json:"transactions"` // List of transactions in the path
	HopCount     int           `json:"hop_count"`    // Number of hops (transactions) in the path
	TotalWeight  float64       `json:"total_weight"` // Aggregate weight of the path
}

// MaliciousWallet represents a wallet identified as malicious
type MaliciousWallet struct {
	Address          string    `json:"address"`           // Wallet address
	FlaggedTimestamp time.Time `json:"flagged_timestamp"` // When the wallet was flagged
	FlagReason       string    `json:"flag_reason"`       // Reason for flagging
	ConfidenceScore  float64   `json:"confidence_score"`  // Confidence in the malicious classification
}

// AssociationRiskAnalyzer performs association risk analysis on transactions
// to identify wallets linked to malicious actors
type AssociationRiskAnalyzer struct {
	randomWindows   *RandomWindows
	walletTrie      *storage.WalletStateTrie
	transactions    map[string][]Transaction      // Map from wallet address to transactions
	maliciousWallets map[string]MaliciousWallet   // Map of known malicious wallets
	pathCache       map[string][]AssociationPath  // Cache of computed paths to avoid repeated computation
	config          *utils.ConfigManager
	logger          *utils.Logger
	metrics         *utils.MetricsManager
	mu              sync.RWMutex
	
	// Transaction weight parameters
	beta  float64 // Parameter for clustering weight
	delta float64 // Parameter for time decay factor
}

// NewAssociationRiskAnalyzer creates a new association risk analyzer
func NewAssociationRiskAnalyzer(
	randomWindows *RandomWindows,
	walletTrie *storage.WalletStateTrie,
	config *utils.ConfigManager,
	logger *utils.Logger,
	metrics *utils.MetricsManager,
) *AssociationRiskAnalyzer {
	
	// Default parameters from the whitepaper
	beta := 0.5
	delta := 0.01
	
	// Override with configuration if available
	if config != nil {
		beta = config.GetFloat64("fraud.association.beta")
		delta = config.GetFloat64("fraud.association.delta")
	}
	
	ara := &AssociationRiskAnalyzer{
		randomWindows:    randomWindows,
		walletTrie:       walletTrie,
		transactions:     make(map[string][]Transaction),
		maliciousWallets: make(map[string]MaliciousWallet),
		pathCache:        make(map[string][]AssociationPath),
		config:           config,
		logger:           logger.WithComponent("AssociationRiskAnalyzer"),
		metrics:          metrics,
		beta:             beta,
		delta:            delta,
	}
	
	ara.logger.Info("Initialized association risk analyzer",
		"beta", ara.beta,
		"delta", ara.delta)
	
	return ara
}

// AddTransaction adds a new transaction to the analyzer's database
func (ara *AssociationRiskAnalyzer) AddTransaction(tx Transaction) {
	ara.mu.Lock()
	defer ara.mu.Unlock()
	
	// Add to sender's transactions
	ara.transactions[tx.FromAddress] = append(ara.transactions[tx.FromAddress], tx)
	
	// Add to receiver's transactions
	ara.transactions[tx.ToAddress] = append(ara.transactions[tx.ToAddress], tx)
	
	// Calculate clustering count based on previous transactions in a short time window
	// For simplicity, we're just incrementing the existing count
	// In a full implementation, this would analyze the transaction frequency pattern
	if len(ara.transactions[tx.FromAddress]) > 1 {
		lastTx := ara.transactions[tx.FromAddress][len(ara.transactions[tx.FromAddress])-2]
		timeDiff := tx.Timestamp.Sub(lastTx.Timestamp)
		
		// If transactions occurred within close proximity (e.g., 1 hour)
		if timeDiff < time.Hour {
			tx.ClusteringCount = lastTx.ClusteringCount + 1
		} else {
			tx.ClusteringCount = 1
		}
	} else {
		tx.ClusteringCount = 1
	}
	
	// Clear path cache since new transactions could change paths
	ara.pathCache = make(map[string][]AssociationPath)
	
	ara.logger.Debug("Added transaction",
		"from", tx.FromAddress,
		"to", tx.ToAddress,
		"amount", tx.Amount,
		"cluster_count", tx.ClusteringCount)
}

// AddMaliciousWallet adds a wallet to the list of known malicious wallets
func (ara *AssociationRiskAnalyzer) AddMaliciousWallet(
	address string,
	reason string,
	confidenceScore float64,
) {
	ara.mu.Lock()
	defer ara.mu.Unlock()
	
	ara.maliciousWallets[address] = MaliciousWallet{
		Address:          address,
		FlaggedTimestamp: time.Now(),
		FlagReason:       reason,
		ConfidenceScore:  confidenceScore,
	}
	
	// Clear path cache since the set of malicious wallets changed
	ara.pathCache = make(map[string][]AssociationPath)
	
	ara.logger.Info("Added malicious wallet",
		"address", address,
		"reason", reason,
		"confidence", confidenceScore)
}

// IsMaliciousWallet checks if a wallet is in the malicious list
func (ara *AssociationRiskAnalyzer) IsMaliciousWallet(address string) bool {
	ara.mu.RLock()
	defer ara.mu.RUnlock()
	
	_, isMalicious := ara.maliciousWallets[address]
	return isMalicious
}

// CalculateTransactionWeight calculates the weight of a transaction
// based on amount, clustering, and time as per the whitepaper formula:
// w_i = log(1 + A_i) × [1 + β × (n_cluster - 1)] × e^(-δ × (t_current - t_i))
func (ara *AssociationRiskAnalyzer) CalculateTransactionWeight(tx Transaction, currentTime time.Time) float64 {
	return utils.TransactionWeight(
		tx.Amount,
		tx.ClusteringCount,
		tx.Timestamp,
		currentTime,
		ara.beta,
		ara.delta,
	)
}

// CalculateAssociationRisk calculates the association risk for a wallet
// based on its connections to known malicious wallets
func (ara *AssociationRiskAnalyzer) CalculateAssociationRisk(
	walletAddress string,
	triggerTimestamp time.Time,
) (float64, []AssociationPath) {
	ara.mu.RLock()
	defer ara.mu.RUnlock()
	
	// Generate random time window and hop limit for this analysis
	params := ara.randomWindows.GenerateForWallet(walletAddress, triggerTimestamp)
	
	// Check for cached results
	cacheKey := walletAddress + ":" + triggerTimestamp.String()
	if paths, exists := ara.pathCache[cacheKey]; exists {
		// Calculate total risk from cached paths
		totalRisk := 0.0
		for _, path := range paths {
			totalRisk += path.TotalWeight
		}
		return totalRisk, paths
	}
	
	// Find all paths from malicious wallets to the target wallet
	// within the random hop limit and time window
	paths := ara.findAssociationPaths(walletAddress, params)
	
	// Calculate total risk as sum of path weights
	totalRisk := 0.0
	for _, path := range paths {
		totalRisk += path.TotalWeight
	}
	
	// Cache the results
	ara.pathCache[cacheKey] = paths
	
	// Update metrics if available
	if ara.metrics != nil && len(paths) > 0 {
		ara.metrics.ObserveHistogram("association_risk_score", totalRisk)
	}
	
	ara.logger.Debug("Calculated association risk",
		"wallet", walletAddress,
		"risk_score", totalRisk,
		"path_count", len(paths),
		"time_window_days", params.TimeLimit.Hours()/24,
		"hop_limit", params.HopLimit)
	
	return totalRisk, paths
}

// findAssociationPaths finds all paths from malicious wallets to the target wallet
// within the specified hop limit and time window
func (ara *AssociationRiskAnalyzer) findAssociationPaths(
	targetWallet string,
	params RandomWindowsParams,
) []AssociationPath {
	paths := []AssociationPath{}
	
	// Current time for weight calculation
	currentTime := time.Now()
	
	// Starting point: check all malicious wallets
	for maliciousAddr := range ara.maliciousWallets {
		// Check direct transaction from malicious to target (1-hop)
		// and recursively discover multi-hop paths
		visited := make(map[string]bool)
		path := []Transaction{}
		
		ara.findPathsDFS(
			maliciousAddr,
			targetWallet,
			params.HopLimit,
			params.TimeLimit,
			ara.maliciousWallets[maliciousAddr].FlaggedTimestamp,
			visited,
			path,
			&paths,
			currentTime,
		)
	}
	
	// Calculate total weight for each path
	for i := range paths {
		totalWeight := 0.0
		for _, tx := range paths[i].Transactions {
			weight := ara.CalculateTransactionWeight(tx, currentTime)
			totalWeight += weight
		}
		paths[i].TotalWeight = totalWeight
	}
	
	return paths
}

// findPathsDFS performs depth-first search to find all paths
// from a malicious wallet to the target within hop limit and time window
func (ara *AssociationRiskAnalyzer) findPathsDFS(
	currentAddr string,
	targetAddr string,
	remainingHops int,
	timeLimit time.Duration,
	triggerTimestamp time.Time,
	visited map[string]bool,
	currentPath []Transaction,
	allPaths *[]AssociationPath,
	currentTime time.Time,
) {
	// If we already visited this wallet or exceeded hop limit, stop
	if visited[currentAddr] || remainingHops <= 0 {
		return
	}
	
	// Mark as visited
	visited[currentAddr] = true
	
	// Check if we've reached the target
	if currentAddr == targetAddr && len(currentPath) > 0 {
		// We found a path, add it to results
		*allPaths = append(*allPaths, AssociationPath{
			Transactions: append([]Transaction{}, currentPath...),
			HopCount:     len(currentPath),
			TotalWeight:  0, // Will be calculated after DFS
		})
		
		// Don't return here, continue to explore other paths through this node
	}
	
	// Get transactions from this wallet
	transactions := ara.transactions[currentAddr]
	
	// Try all outgoing transactions
	for _, tx := range transactions {
		// Skip if transaction is not within time window
		if !ara.randomWindows.IsWithinTimeWindow(tx.Timestamp, triggerTimestamp, RandomWindowsParams{
			TimeLimit: timeLimit,
		}) {
			continue
		}
		
		// Skip if transaction is not from current wallet
		if tx.FromAddress != currentAddr {
			continue
		}
		
		// Add transaction to current path
		newPath := append(currentPath, tx)
		
		// Continue DFS from the recipient wallet
		visitedCopy := make(map[string]bool)
		for k, v := range visited {
			visitedCopy[k] = v
		}
		
		ara.findPathsDFS(
			tx.ToAddress,
			targetAddr,
			remainingHops-1,
			timeLimit,
			triggerTimestamp,
			visitedCopy,
			newPath,
			allPaths,
			currentTime,
		)
	}
}

// UpdateWalletAssociationRisk updates a wallet's association risk score in storage
func (ara *AssociationRiskAnalyzer) UpdateWalletAssociationRisk(walletAddress string, riskScore float64) {
	wallet := ara.walletTrie.GetWallet(walletAddress)
	wallet.UpdateAssociationRisk(riskScore)
	ara.walletTrie.UpdateWallet(wallet)
	
	ara.logger.Debug("Updated wallet association risk",
		"wallet", walletAddress,
		"risk_score", riskScore)
}

// GetAssociationRiskStats returns the statistics for association risk scores
// to be used in calculating standardized deviations for penalties
func (ara *AssociationRiskAnalyzer) GetAssociationRiskStats() (mean, stdDev float64) {
	ara.mu.RLock()
	defer ara.mu.RUnlock()
	
	// Collect all risk scores
	scores := []float64{}
	wallets := ara.walletTrie.GetAllWallets()
	
	for _, wallet := range wallets {
		scores = append(scores, wallet.AssociationRiskScore)
	}
	
	// Calculate mean and standard deviation
	mean = utils.Mean(scores)
	stdDev = utils.StandardDeviation(scores)
	
	// Ensure non-zero standard deviation
	if stdDev == 0 {
		stdDev = 1.0
	}
	
	return mean, stdDev
}

// ProcessAssociationAnalysis runs association risk analysis for a wallet
// and returns the risk score and detailed paths
func (ara *AssociationRiskAnalyzer) ProcessAssociationAnalysis(
	walletAddress string,
	triggerTimestamp time.Time,
) (float64, []AssociationPath) {
	// Calculate association risk
	riskScore, paths := ara.CalculateAssociationRisk(walletAddress, triggerTimestamp)
	
	// Update the wallet's association risk score in storage
	ara.UpdateWalletAssociationRisk(walletAddress, riskScore)
	
	// Update counter metrics
	if ara.metrics != nil && riskScore > 0 {
		ara.metrics.IncCounter("association_penalties_total")
	}
	
	return riskScore, paths
}