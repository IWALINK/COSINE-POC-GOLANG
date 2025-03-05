package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/gofrs/flock"
)

// PeerRegistry manages a list of known peer addresses in a shared file.
type PeerRegistry struct {
	filePath string
	peers    map[string][]string // nodeID -> list of multiaddresses
	mu       sync.RWMutex
	logger   *Logger
	fileLock *flock.Flock
}

// NewPeerRegistry creates a new peer registry using the shared file at filePath.
func NewPeerRegistry(filePath string, logger *Logger) (*PeerRegistry, error) {
	// Ensure the registry directory exists.
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory for peer registry: %w", err)
	}

	// Initialize a file lock (using an external lock file).
	fl := flock.New(filePath + ".lock")

	registry := &PeerRegistry{
		filePath: filePath,
		peers:    make(map[string][]string),
		logger:   logger,
		fileLock: fl,
	}

	// Load existing registry if available.
	if _, err := os.Stat(filePath); err == nil {
		if err := registry.load(); err != nil {
			logger.Warn("Failed to load peer registry, starting with empty registry", "error", err)
		}
	}

	return registry, nil
}

// RegisterPeer adds or merges a node's addresses into the registry.
func (r *PeerRegistry) RegisterPeer(nodeID string, addrs []string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Acquire the file lock to prevent concurrent writes.
	locked, err := r.fileLock.TryLock()
	if err != nil {
		return fmt.Errorf("failed to acquire file lock: %w", err)
	}
	if !locked {
		return fmt.Errorf("could not acquire file lock")
	}
	defer r.fileLock.Unlock()

	// Load the latest registry from disk to merge updates from other nodes.
	if err := r.load(); err != nil {
		r.logger.Warn("Failed to load registry for merging", "error", err)
	}

	// Merge: if this node already exists, combine addresses without duplicates.
	existing, exists := r.peers[nodeID]
	if exists {
		r.peers[nodeID] = mergeStringSlices(existing, addrs)
	} else {
		r.peers[nodeID] = addrs
	}
	return r.save()
}

// mergeStringSlices merges two slices of strings, removing duplicates.
func mergeStringSlices(slice1, slice2 []string) []string {
	m := make(map[string]bool)
	for _, s := range slice1 {
		m[s] = true
	}
	for _, s := range slice2 {
		m[s] = true
	}
	var result []string
	for s := range m {
		result = append(result, s)
	}
	return result
}

// GetPeers returns a copy of all peer entries.
func (r *PeerRegistry) GetPeers() map[string][]string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Load latest changes before returning.
	if err := r.load(); err != nil {
		r.logger.Warn("Failed to load registry in GetPeers", "error", err)
	}
	peersCopy := make(map[string][]string)
	for id, addrs := range r.peers {
		cp := make([]string, len(addrs))
		copy(cp, addrs)
		peersCopy[id] = cp
	}
	return peersCopy
}

// GetPeerAddrs returns the addresses for a specific node.
func (r *PeerRegistry) GetPeerAddrs(nodeID string) []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if err := r.load(); err != nil {
		r.logger.Warn("Failed to load registry in GetPeerAddrs", "error", err)
	}
	addrs, exists := r.peers[nodeID]
	if !exists {
		return nil
	}
	cp := make([]string, len(addrs))
	copy(cp, addrs)
	return cp
}

// RemovePeer removes a peer from the registry.
func (r *PeerRegistry) RemovePeer(nodeID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	locked, err := r.fileLock.TryLock()
	if err != nil {
		return fmt.Errorf("failed to acquire file lock: %w", err)
	}
	if !locked {
		return fmt.Errorf("could not acquire file lock")
	}
	defer r.fileLock.Unlock()

	if err := r.load(); err != nil {
		r.logger.Warn("Failed to load registry for removing peer", "error", err)
	}

	delete(r.peers, nodeID)
	return r.save()
}

// load reads the registry from the shared file.
func (r *PeerRegistry) load() error {
	data, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		return fmt.Errorf("failed to read peer registry file: %w", err)
	}
	var loaded map[string][]string
	if err := json.Unmarshal(data, &loaded); err != nil {
		return fmt.Errorf("failed to unmarshal peer registry: %w", err)
	}
	r.peers = loaded
	return nil
}

// save writes the current registry to the shared file.
func (r *PeerRegistry) save() error {
	data, err := json.MarshalIndent(r.peers, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal peer registry: %w", err)
	}
	return ioutil.WriteFile(r.filePath, data, 0644)
}

// GetBootstrapAddrs returns all peer addresses except for the specified node.
func (r *PeerRegistry) GetBootstrapAddrs(excludeNodeID string) []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if err := r.load(); err != nil {
		r.logger.Warn("Failed to load registry in GetBootstrapAddrs", "error", err)
	}
	var addrs []string
	for id, nodeAddrs := range r.peers {
		if id == excludeNodeID {
			continue
		}
		addrs = append(addrs, nodeAddrs...)
	}
	return addrs
}

// Clear empties the registry.
func (r *PeerRegistry) Clear() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	locked, err := r.fileLock.TryLock()
	if err != nil {
		return fmt.Errorf("failed to acquire file lock: %w", err)
	}
	if !locked {
		return fmt.Errorf("could not acquire file lock")
	}
	defer r.fileLock.Unlock()

	r.peers = make(map[string][]string)
	return r.save()
}
