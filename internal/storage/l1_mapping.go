package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
)

// L1MappingEntry represents a single mapping from an L1 address to an L2 wallet
type L1MappingEntry struct {
	ChainID    uint64    `json:"chain_id"`
	L1Address  string    `json:"l1_address"`
	L2Address  string    `json:"l2_address"`
	LinkedAt   time.Time `json:"linked_at"`
}

// LinkedL1Trie represents a Merkle Trie ensuring one-to-one mapping from L1 addresses to L2 wallets
type LinkedL1Trie struct {
	mappings map[string]L1MappingEntry
	mu       sync.RWMutex
	logger   *utils.Logger
	root     string
	wst      *WalletStateTrie // Reference to wallet state trie
}

// NewLinkedL1Trie creates a new L1 mapping trie
func NewLinkedL1Trie(logger *utils.Logger, wst *WalletStateTrie) *LinkedL1Trie {
	return &LinkedL1Trie{
		mappings: make(map[string]L1MappingEntry),
		logger:   logger.WithComponent("LinkedL1Trie"),
		root:     "",
		wst:      wst,
	}
}

// makeKey generates a key from chain ID and L1 address
func makeKey(chainID uint64, l1Address string) string {
	// In a real implementation, this should be a cryptographic hash
	// Here we use a simple string format for demonstration
	data := fmt.Sprintf("%d:%s", chainID, l1Address)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// LinkAddress creates a new mapping from an L1 address to an L2 wallet
func (lt *LinkedL1Trie) LinkAddress(chainID uint64, l1Address, l2Address string) error {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	
	key := makeKey(chainID, l1Address)
	
	// Check if this L1 address is already linked
	if entry, exists := lt.mappings[key]; exists {
		return fmt.Errorf("L1 address already linked to L2 wallet %s since %s", 
			entry.L2Address, entry.LinkedAt.Format(time.RFC3339))
	}
	
	// Create the mapping
	lt.mappings[key] = L1MappingEntry{
		ChainID:    chainID,
		L1Address:  l1Address,
		L2Address:  l2Address,
		LinkedAt:   time.Now(),
	}
	
	// Update the Merkle root
	lt.updateRoot()
	
	// Update the wallet state if needed
	wallet := lt.wst.GetWallet(l2Address)
	wallet.AddLinkedL1Address(chainID, l1Address)
	lt.wst.UpdateWallet(wallet)
	
	lt.logger.Info("Linked L1 address to L2 wallet", 
		"chain_id", chainID,
		"l1_address", l1Address,
		"l2_address", l2Address)
	
	return nil
}

// GetL2Address returns the L2 wallet address linked to an L1 address
func (lt *LinkedL1Trie) GetL2Address(chainID uint64, l1Address string) (string, error) {
	lt.mu.RLock()
	defer lt.mu.RUnlock()
	
	key := makeKey(chainID, l1Address)
	entry, exists := lt.mappings[key]
	if !exists {
		return "", fmt.Errorf("no L2 wallet linked to this L1 address")
	}
	
	return entry.L2Address, nil
}

// IsLinked checks if an L1 address is already linked
func (lt *LinkedL1Trie) IsLinked(chainID uint64, l1Address string) bool {
	lt.mu.RLock()
	defer lt.mu.RUnlock()
	
	key := makeKey(chainID, l1Address)
	_, exists := lt.mappings[key]
	return exists
}

// GetLinkInfo returns detailed information about a link
func (lt *LinkedL1Trie) GetLinkInfo(chainID uint64, l1Address string) (*L1MappingEntry, error) {
	lt.mu.RLock()
	defer lt.mu.RUnlock()
	
	key := makeKey(chainID, l1Address)
	entry, exists := lt.mappings[key]
	if !exists {
		return nil, fmt.Errorf("no L2 wallet linked to this L1 address")
	}
	
	// Return a copy to prevent modification
	entryCopy := entry
	return &entryCopy, nil
}

// GetAllLinksForL2Wallet returns all L1 addresses linked to an L2 wallet
func (lt *LinkedL1Trie) GetAllLinksForL2Wallet(l2Address string) []L1MappingEntry {
	lt.mu.RLock()
	defer lt.mu.RUnlock()
	
	result := []L1MappingEntry{}
	for _, entry := range lt.mappings {
		if entry.L2Address == l2Address {
			result = append(result, entry)
		}
	}
	
	return result
}

// RemoveLink removes a link between an L1 address and L2 wallet
func (lt *LinkedL1Trie) RemoveLink(chainID uint64, l1Address string) error {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	
	key := makeKey(chainID, l1Address)
	entry, exists := lt.mappings[key]
	if !exists {
		return fmt.Errorf("no L2 wallet linked to this L1 address")
	}
	
	// Remove the mapping
	delete(lt.mappings, key)
	
	// Update the Merkle root
	lt.updateRoot()
	
	lt.logger.Info("Removed link between L1 address and L2 wallet", 
		"chain_id", chainID,
		"l1_address", l1Address,
		"l2_address", entry.L2Address)
	
	return nil
}

// GetRoot returns the current Merkle root of the trie
func (lt *LinkedL1Trie) GetRoot() string {
	lt.mu.RLock()
	defer lt.mu.RUnlock()
	return lt.root
}

// updateRoot recalculates the Merkle root of the trie
// This is a simplified implementation - in production, you'd use a real Merkle Patricia Trie
func (lt *LinkedL1Trie) updateRoot() {
	// Simple concatenation of all keys for demo
	// In production: implement actual Merkle Patricia Trie logic
	var concatKeys string
	for key := range lt.mappings {
		concatKeys += key
	}
	
	// Hash the concatenation to get a root
	hash := sha256.Sum256([]byte(concatKeys))
	lt.root = hex.EncodeToString(hash[:])
}

// GetLinkCount returns the total number of links in the trie
func (lt *LinkedL1Trie) GetLinkCount() int {
	lt.mu.RLock()
	defer lt.mu.RUnlock()
	return len(lt.mappings)
}

// VerifyConsistency checks that the L1T and WST mappings are consistent
func (lt *LinkedL1Trie) VerifyConsistency() (bool, []string) {
	lt.mu.RLock()
	defer lt.mu.RUnlock()
	
	inconsistencies := []string{}
	
	// Check each mapping in L1T against WST
	for _, entry := range lt.mappings {
		wallet := lt.wst.GetWallet(entry.L2Address)
		found := false
		
		for _, linked := range wallet.LinkedL1Addresses {
			if linked.ChainID == entry.ChainID && linked.Address == entry.L1Address {
				found = true
				break
			}
		}
		
		if !found {
			inconsistencies = append(inconsistencies, 
				fmt.Sprintf("L1 address (chain: %d, addr: %s) linked to L2 wallet %s in L1T but not in WST", 
					entry.ChainID, entry.L1Address, entry.L2Address))
		}
	}
	
	// Check each wallet in WST against L1T
	for _, wallet := range lt.wst.GetAllWallets() {
		for _, linked := range wallet.LinkedL1Addresses {
			key := makeKey(linked.ChainID, linked.Address)
			entry, exists := lt.mappings[key]
			
			if !exists {
				inconsistencies = append(inconsistencies, 
					fmt.Sprintf("L1 address (chain: %d, addr: %s) linked to L2 wallet %s in WST but not in L1T", 
						linked.ChainID, linked.Address, wallet.Address))
			} else if entry.L2Address != wallet.Address {
				inconsistencies = append(inconsistencies, 
					fmt.Sprintf("L1 address (chain: %d, addr: %s) linked to different L2 wallets: %s in L1T, %s in WST", 
						linked.ChainID, linked.Address, entry.L2Address, wallet.Address))
			}
		}
	}
	
	return len(inconsistencies) == 0, inconsistencies
}