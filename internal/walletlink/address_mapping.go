package walletlink

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
)

// Common errors
var (
	ErrInvalidL2Wallet      = errors.New("invalid L2 wallet format")
	ErrMappingExists        = errors.New("mapping already exists")
	ErrMappingNotFound      = errors.New("mapping not found")
	ErrProofVerificationFailed = errors.New("proof verification failed")
	ErrNoCrossChainLinks    = errors.New("no cross-chain links found for this L2 wallet")
)

// ChainMappingInfo contains information about an L1 address linked to an L2 wallet
type ChainMappingInfo struct {
	ChainID       uint64    `json:"chain_id"`
	ChainName     string    `json:"chain_name"`
	L1Address     string    `json:"l1_address"`
	LinkedAt      time.Time `json:"linked_at"`
	LastVerified  time.Time `json:"last_verified"`
}

// L2WalletInfo contains information about an L2 wallet and its linked L1 addresses
type L2WalletInfo struct {
	L2Wallet      string             `json:"l2_wallet"`
	CreatedAt     time.Time          `json:"created_at"`
	LinkedChains  []ChainMappingInfo `json:"linked_chains"`
}

// AddressMappingManager manages the mapping between L1 addresses and L2 wallets
type AddressMappingManager struct {
	l1Trie        *storage.LinkedL1Trie
	proofVerifier *ProofVerifier
	logger        *utils.Logger
	metrics       *utils.MetricsManager
	
	// Cache for quick lookups (to reduce storage access)
	l1ToL2Cache   map[string]string      // chainID:L1Address -> L2Wallet
	l2ToL1Cache   map[string]map[uint64]string // L2Wallet -> chainID -> L1Address
	
	mu            sync.RWMutex
}

// NewAddressMappingManager creates a new address mapping manager
func NewAddressMappingManager(
	l1Trie *storage.LinkedL1Trie,
	proofVerifier *ProofVerifier,
	logger *utils.Logger,
	metrics *utils.MetricsManager,
) *AddressMappingManager {
	return &AddressMappingManager{
		l1Trie:        l1Trie,
		proofVerifier: proofVerifier,
		logger:        logger.WithComponent("AddressMappingManager"),
		metrics:       metrics,
		l1ToL2Cache:   make(map[string]string),
		l2ToL1Cache:   make(map[string]map[uint64]string),
	}
}

// makeL1CacheKey creates a cache key from chain ID and L1 address
func makeL1CacheKey(chainID uint64, l1Address string) string {
	return fmt.Sprintf("%d:%s", chainID, l1Address)
}

// LinkAddressWithVerification links an L1 address to an L2 wallet after verifying ownership
func (m *AddressMappingManager) LinkAddressWithVerification(
	ctx context.Context,
	chainType ChainType,
	chainID uint64,
	l1Address string,
	l2WalletHex string,
	signature string,
) error {
	// Validate the L2 wallet
	if err := m.validateL2Wallet(l2WalletHex); err != nil {
		return err
	}
	
	startTime := time.Now()
	m.logger.Info("Linking address with verification", 
		"chain", chainType, 
		"chain_id", chainID,
		"l1_address", l1Address, 
		"l2_wallet", l2WalletHex)
	
	// First verify the proof using the ProofVerifier
	isValid, err := m.proofVerifier.VerifyProof(ctx, chainType, l1Address, l2WalletHex, signature)
	if err != nil {
		m.logger.Error("Proof verification failed", "error", err)
		return fmt.Errorf("failed to verify proof: %w", err)
	}
	
	if !isValid {
		return ErrProofVerificationFailed
	}
	
	// Check if the L1 address is already linked
	isLinked, err := m.IsAddressLinked(ctx, chainID, l1Address)
	if err != nil {
		return fmt.Errorf("failed to check if address is linked: %w", err)
	}
	
	if isLinked {
		return ErrMappingExists
	}
	
	// Create a new mapping in the L1 trie
	err = m.l1Trie.LinkAddress(chainID, l1Address, l2WalletHex)
	if err != nil {
		m.logger.Error("Failed to link address in trie", "error", err)
		return fmt.Errorf("failed to link address in storage: %w", err)
	}
	
	// Update the cache
	m.mu.Lock()
	m.l1ToL2Cache[makeL1CacheKey(chainID, l1Address)] = l2WalletHex
	
	// Update the reverse mapping cache
	if _, exists := m.l2ToL1Cache[l2WalletHex]; !exists {
		m.l2ToL1Cache[l2WalletHex] = make(map[uint64]string)
	}
	m.l2ToL1Cache[l2WalletHex][chainID] = l1Address
	m.mu.Unlock()
	
	// Record metrics
	if m.metrics != nil {
		m.metrics.IncCounter("addresses_linked_total", "chain", string(chainType))
		m.metrics.ObserveHistogram("address_linking_duration_seconds", time.Since(startTime).Seconds())
	}
	
	m.logger.Info("Address linked successfully", 
		"chain", chainType, 
		"chain_id", chainID,
		"l1_address", l1Address, 
		"l2_wallet", l2WalletHex,
		"duration", time.Since(startTime))
	
	return nil
}

// validateL2Wallet validates the format of an L2 wallet
func (m *AddressMappingManager) validateL2Wallet(l2WalletHex string) error {
	// Ensure the L2 wallet is a valid 32-byte hex string (with or without 0x prefix)
	l2WalletHex = strings.TrimPrefix(l2WalletHex, "0x")
	if len(l2WalletHex) != 64 {
		return fmt.Errorf("%w: must be a 32-byte hex string", ErrInvalidL2Wallet)
	}
	
	for _, c := range l2WalletHex {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return fmt.Errorf("%w: contains invalid hex characters", ErrInvalidL2Wallet)
		}
	}
	
	return nil
}

// IsAddressLinked checks if an L1 address is already linked to an L2 wallet
func (m *AddressMappingManager) IsAddressLinked(ctx context.Context, chainID uint64, l1Address string) (bool, error) {
	// First check the cache
	m.mu.RLock()
	cacheKey := makeL1CacheKey(chainID, l1Address)
	l2Wallet, exists := m.l1ToL2Cache[cacheKey]
	m.mu.RUnlock()
	
	if exists && l2Wallet != "" {
		return true, nil
	}
	
	// If not in cache, check the storage
	isLinked := m.l1Trie.IsLinked(chainID, l1Address)
	
	return isLinked, nil
}

// IsL2WalletLinked checks if an L2 wallet is linked to any L1 address
func (m *AddressMappingManager) IsL2WalletLinked(ctx context.Context, l2WalletHex string) (bool, error) {
	// Validate the L2 wallet
	if err := m.validateL2Wallet(l2WalletHex); err != nil {
		return false, err
	}
	
	// First check the cache
	m.mu.RLock()
	chainMappings, exists := m.l2ToL1Cache[l2WalletHex]
	hasLinks := exists && len(chainMappings) > 0
	m.mu.RUnlock()
	
	if hasLinks {
		return true, nil
	}
	
	// If not in cache, check the storage
	links := m.l1Trie.GetAllLinksForL2Wallet(l2WalletHex)
	hasLinks = len(links) > 0
	
	return hasLinks, nil
}

// GetLinkedL2Wallet gets the L2 wallet linked to an L1 address
func (m *AddressMappingManager) GetLinkedL2Wallet(ctx context.Context, chainID uint64, l1Address string) (string, error) {
	// First check the cache
	m.mu.RLock()
	cacheKey := makeL1CacheKey(chainID, l1Address)
	l2Wallet, exists := m.l1ToL2Cache[cacheKey]
	m.mu.RUnlock()
	
	if exists && l2Wallet != "" {
		return l2Wallet, nil
	}
	
	// If not in cache, check the storage
	l2Wallet, err := m.l1Trie.GetL2Address(chainID, l1Address)
	if err != nil {
		if strings.Contains(err.Error(), "no L2 wallet") {
			return "", ErrMappingNotFound
		}
		return "", fmt.Errorf("failed to get linked L2 wallet: %w", err)
	}
	
	// Update the cache
	m.mu.Lock()
	m.l1ToL2Cache[cacheKey] = l2Wallet
	m.mu.Unlock()
	
	return l2Wallet, nil
}

// GetLinkedL1Address gets the L1 address linked to an L2 wallet for a specific chain
func (m *AddressMappingManager) GetLinkedL1Address(ctx context.Context, chainID uint64, l2WalletHex string) (string, error) {
	// Validate the L2 wallet
	if err := m.validateL2Wallet(l2WalletHex); err != nil {
		return "", err
	}
	
	// First check the cache
	m.mu.RLock()
	chainMappings, exists := m.l2ToL1Cache[l2WalletHex]
	if exists {
		if l1Address, found := chainMappings[chainID]; found && l1Address != "" {
			m.mu.RUnlock()
			return l1Address, nil
		}
	}
	m.mu.RUnlock()
	
	// If not in cache, check the storage
	// Get all links and find the one for the specified chain
	links := m.l1Trie.GetAllLinksForL2Wallet(l2WalletHex)
	for _, link := range links {
		if link.ChainID == chainID {
			// Update the cache
			m.mu.Lock()
			if _, exists := m.l2ToL1Cache[l2WalletHex]; !exists {
				m.l2ToL1Cache[l2WalletHex] = make(map[uint64]string)
			}
			m.l2ToL1Cache[l2WalletHex][chainID] = link.L1Address
			m.mu.Unlock()
			
			return link.L1Address, nil
		}
	}
	
	return "", ErrMappingNotFound
}

// GetAllLinkedAddresses gets all L1 addresses linked to an L2 wallet
func (m *AddressMappingManager) GetAllLinkedAddresses(ctx context.Context, l2WalletHex string) ([]ChainMappingInfo, error) {
	// Validate the L2 wallet
	if err := m.validateL2Wallet(l2WalletHex); err != nil {
		return nil, err
	}
	
	// Get all links from storage
	links := m.l1Trie.GetAllLinksForL2Wallet(l2WalletHex)
	if len(links) == 0 {
		return nil, ErrNoCrossChainLinks
	}
	
	// Convert to ChainMappingInfo
	mappings := make([]ChainMappingInfo, 0, len(links))
	for _, link := range links {
		// Determine chain name from chain ID
		chainName := getChainName(link.ChainID)
		
		mapping := ChainMappingInfo{
			ChainID:      link.ChainID,
			ChainName:    chainName,
			L1Address:    link.L1Address,
			LinkedAt:     link.LinkedAt,
			LastVerified: link.LinkedAt, // Initially the same as LinkedAt
		}
		
		mappings = append(mappings, mapping)
	}
	
	// Update the cache
	m.mu.Lock()
	if _, exists := m.l2ToL1Cache[l2WalletHex]; !exists {
		m.l2ToL1Cache[l2WalletHex] = make(map[uint64]string)
	}
	
	for _, mapping := range mappings {
		m.l2ToL1Cache[l2WalletHex][mapping.ChainID] = mapping.L1Address
		m.l1ToL2Cache[makeL1CacheKey(mapping.ChainID, mapping.L1Address)] = l2WalletHex
	}
	m.mu.Unlock()
	
	return mappings, nil
}

// GetL2WalletInfo gets comprehensive information about an L2 wallet and its linked addresses
func (m *AddressMappingManager) GetL2WalletInfo(ctx context.Context, l2WalletHex string) (*L2WalletInfo, error) {
	// Validate the L2 wallet
	if err := m.validateL2Wallet(l2WalletHex); err != nil {
		return nil, err
	}
	
	// Get all linked addresses
	mappings, err := m.GetAllLinkedAddresses(ctx, l2WalletHex)
	if err != nil {
		if errors.Is(err, ErrNoCrossChainLinks) {
			// If no links found, return basic info
			return &L2WalletInfo{
				L2Wallet:     l2WalletHex,
				CreatedAt:    time.Now(), // We don't have actual creation time
				LinkedChains: []ChainMappingInfo{},
			}, nil
		}
		return nil, err
	}
	
	// Use the earliest LinkedAt time as the CreatedAt time
	createdAt := time.Now()
	for _, mapping := range mappings {
		if mapping.LinkedAt.Before(createdAt) {
			createdAt = mapping.LinkedAt
		}
	}
	
	return &L2WalletInfo{
		L2Wallet:     l2WalletHex,
		CreatedAt:    createdAt,
		LinkedChains: mappings,
	}, nil
}

// ClearCache clears the mapping cache, forcing future lookups to use storage
func (m *AddressMappingManager) ClearCache() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.l1ToL2Cache = make(map[string]string)
	m.l2ToL1Cache = make(map[string]map[uint64]string)
	
	m.logger.Info("Address mapping cache cleared")
}

// VerifyConsistency checks that the storage mappings are consistent
func (m *AddressMappingManager) VerifyConsistency(ctx context.Context) (bool, []string) {
	m.logger.Info("Verifying address mapping consistency")
	
	// Delegate to L1Trie to verify consistency
	isConsistent, inconsistencies := m.l1Trie.VerifyConsistency()
	
	if !isConsistent {
		m.logger.Warn("Address mapping inconsistencies found", 
			"count", len(inconsistencies))
		
		// Log each inconsistency
		for i, inconsistency := range inconsistencies {
			m.logger.Warn("Inconsistency", "index", i, "details", inconsistency)
		}
	} else {
		m.logger.Info("Address mappings are consistent")
	}
	
	return isConsistent, inconsistencies
}

// getChainName returns a human-readable name for a chain ID
func getChainName(chainID uint64) string {
	switch chainID {
	case 1:
		return "Ethereum"
	case 10:
		return "Optimism"
	case 56:
		return "BNB Chain"
	case 137:
		return "Polygon"
	case 42161:
		return "Arbitrum"
	case 43114:
		return "Avalanche"
	case 502:
		return "Bitcoin"
	default:
		return fmt.Sprintf("Chain %d", chainID)
	}
}