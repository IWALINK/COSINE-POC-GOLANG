package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
)

// LinkedL1Address represents a blockchain address on an L1 chain linked to an L2 wallet
type LinkedL1Address struct {
	ChainID uint64 `json:"chain_id"`
	Address string `json:"address"`
}

// VoteEvent represents a single vote cast by a wallet
type VoteEvent struct {
	VoterAddress  string    `json:"voter_address"`
	VoteValue     int       `json:"vote_value"` // -1 or +1
	Timestamp     time.Time `json:"timestamp"`
	EffectiveVote float64   `json:"effective_vote"` // R_U * vote_value
}

// RollingStatistic represents an exponential moving average with its last updated time
type RollingStatistic struct {
	Value        float64   `json:"value"`
	LastUpdated  time.Time `json:"last_updated"`
}

// WalletState represents a wallet's complete state in the COSINE L2
type WalletState struct {
	// L2 wallet address (not stored directly as it's used as key in the trie)
	Address string `json:"-"`

	// Linked L1 addresses (e.g., Ethereum, Bitcoin)
	LinkedL1Addresses []LinkedL1Address `json:"linked_l1_addresses"`

	// Accumulated cost of computation since last verification
	AccumulatedCost float64 `json:"c_acc"`

	// Reputation score in [0, 1] (for voting)
	ReputationScore float64 `json:"r_u"`

	// Rolling statistics for EMAs
	RollingStats map[string]RollingStatistic `json:"rolling_stats"`

	// Last update timestamp
	LastUpdated time.Time `json:"last_updated"`

	// Credit score (unbounded)
	CreditScore float64 `json:"c_w"`

	// Association risk score
	AssociationRiskScore float64 `json:"r_assoc"`

	// Vote history
	VoteHistory []VoteEvent `json:"vote_history"`

	// Token balance
	TokenBalance uint64 `json:"token_balance"`
}

// NewWalletState creates a new wallet state with default values
func NewWalletState(address string) *WalletState {
	now := time.Now()
	return &WalletState{
		Address:            address,
		LinkedL1Addresses:  []LinkedL1Address{},
		AccumulatedCost:    0,
		ReputationScore:    0.1, // Default initial reputation
		LastUpdated:        now,
		CreditScore:        0,   // Default initial credit score
		AssociationRiskScore: 0,
		VoteHistory:        []VoteEvent{},
		TokenBalance:       0,
		RollingStats: map[string]RollingStatistic{
			"μ_V": {Value: 0, LastUpdated: now},
			"σ_V": {Value: 1, LastUpdated: now},
			"μ_R": {Value: 0, LastUpdated: now},
			"σ_R": {Value: 1, LastUpdated: now},
		},
	}
}

// AddLinkedL1Address adds an L1 address to this wallet's links
func (ws *WalletState) AddLinkedL1Address(chainID uint64, address string) {
	// Check if already linked
	for _, linked := range ws.LinkedL1Addresses {
		if linked.ChainID == chainID && linked.Address == address {
			return
		}
	}
	
	ws.LinkedL1Addresses = append(ws.LinkedL1Addresses, LinkedL1Address{
		ChainID: chainID,
		Address: address,
	})
	ws.LastUpdated = time.Now()
}

// AddVoteEvent records a new vote event to the wallet's history
func (ws *WalletState) AddVoteEvent(vote VoteEvent) {
	ws.VoteHistory = append(ws.VoteHistory, vote)
	ws.LastUpdated = time.Now()
}

// UpdateCreditScore updates the wallet's credit score
func (ws *WalletState) UpdateCreditScore(delta float64) {
	ws.CreditScore += delta
	ws.LastUpdated = time.Now()
}

// UpdateRollingStatistic updates a rolling statistic using EMA
func (ws *WalletState) UpdateRollingStatistic(name string, newValue float64, alpha float64) {
	stat, exists := ws.RollingStats[name]
	if !exists {
		stat = RollingStatistic{Value: newValue, LastUpdated: time.Now()}
	} else {
		stat.Value = utils.ExponentialMovingAverage(stat.Value, newValue, alpha)
		stat.LastUpdated = time.Now()
	}
	ws.RollingStats[name] = stat
}

// IncrementAccumulatedCost adds to the accumulated computation cost
func (ws *WalletState) IncrementAccumulatedCost(amount float64) {
	ws.AccumulatedCost += amount
	ws.LastUpdated = time.Now()
}

// ResetAccumulatedCost resets the accumulated cost (after verification fee payment)
func (ws *WalletState) ResetAccumulatedCost() {
	ws.AccumulatedCost = 0
	ws.LastUpdated = time.Now()
}

// TransferTokens transfers tokens from this wallet to another
func (ws *WalletState) TransferTokens(amount uint64) error {
	if ws.TokenBalance < amount {
		return fmt.Errorf("insufficient balance: have %d, want %d", ws.TokenBalance, amount)
	}
	ws.TokenBalance -= amount
	ws.LastUpdated = time.Now()
	return nil
}

// ReceiveTokens adds tokens to this wallet
func (ws *WalletState) ReceiveTokens(amount uint64) {
	ws.TokenBalance += amount
	ws.LastUpdated = time.Now()
}

// UpdateAssociationRisk updates the association risk score
func (ws *WalletState) UpdateAssociationRisk(delta float64) {
	ws.AssociationRiskScore += delta
	ws.LastUpdated = time.Now()
}

// GetNormalizedCreditVector returns a 2D vector for cosine similarity calculation
func (ws *WalletState) GetNormalizedCreditVector(mean, stdDev float64) utils.Vector {
	return utils.VectorizeScore(ws.CreditScore, mean, stdDev)
}

// WalletStateTrie represents a Merkle Patricia Trie of wallet states
type WalletStateTrie struct {
	wallets map[string]*WalletState
	mu      sync.RWMutex
	logger  *utils.Logger
	root    string // Root hash of the trie
}

// NewWalletStateTrie creates a new wallet state trie
func NewWalletStateTrie(logger *utils.Logger) *WalletStateTrie {
	return &WalletStateTrie{
		wallets: make(map[string]*WalletState),
		logger:  logger.WithComponent("WalletStateTrie"),
		root:    "",
	}
}

// GetWallet returns a wallet state by address, creating one if it doesn't exist
func (wst *WalletStateTrie) GetWallet(address string) *WalletState {
	wst.mu.RLock()
	wallet, exists := wst.wallets[address]
	wst.mu.RUnlock()
	
	if !exists {
		wst.mu.Lock()
		// Double-check to handle race condition
		wallet, exists = wst.wallets[address]
		if !exists {
			wallet = NewWalletState(address)
			wst.wallets[address] = wallet
			wst.logger.Debug("Created new wallet state", "address", address)
		}
		wst.mu.Unlock()
	}
	
	return wallet
}

// UpdateWallet saves a wallet state and updates the root hash
func (wst *WalletStateTrie) UpdateWallet(wallet *WalletState) {
	wst.mu.Lock()
	defer wst.mu.Unlock()
	
	wst.wallets[wallet.Address] = wallet
	wst.updateRoot()
	wst.logger.Debug("Updated wallet state", "address", wallet.Address, "new_root", wst.root)
}

// DeleteWallet removes a wallet from the trie
func (wst *WalletStateTrie) DeleteWallet(address string) {
	wst.mu.Lock()
	defer wst.mu.Unlock()
	
	delete(wst.wallets, address)
	wst.updateRoot()
	wst.logger.Debug("Deleted wallet state", "address", address, "new_root", wst.root)
}

// GetRoot returns the current root hash of the trie
func (wst *WalletStateTrie) GetRoot() string {
	wst.mu.RLock()
	defer wst.mu.RUnlock()
	return wst.root
}

// updateRoot recalculates the Merkle root of the trie
// This is a simplified implementation - in production, you'd use a real Merkle Patricia Trie
func (wst *WalletStateTrie) updateRoot() {
	// Simple concatenation of all addresses for demo
	// In production: implement actual Merkle Patricia Trie logic
	var concatAddresses string
	for addr := range wst.wallets {
		concatAddresses += addr
	}
	
	// Hash the concatenation to get a root
	hash := sha256.Sum256([]byte(concatAddresses))
	wst.root = hex.EncodeToString(hash[:])
}

// GetAllWallets returns all wallets in the trie
func (wst *WalletStateTrie) GetAllWallets() []*WalletState {
	wst.mu.RLock()
	defer wst.mu.RUnlock()
	
	wallets := make([]*WalletState, 0, len(wst.wallets))
	for _, wallet := range wst.wallets {
		wallets = append(wallets, wallet)
	}
	return wallets
}

// GetWalletCount returns the number of wallets in the trie
func (wst *WalletStateTrie) GetWalletCount() int {
	wst.mu.RLock()
	defer wst.mu.RUnlock()
	return len(wst.wallets)
}