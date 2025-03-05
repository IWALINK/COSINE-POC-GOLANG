// Package p2p provides peer-to-peer networking functionality for COSINE validators.
package p2p

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/multiformats/go-multiaddr"
)

// NetworkManager represents the P2P network manager for COSINE validators.
type NetworkManager struct {
	host      host.Host
	ctx       context.Context
	cancel    context.CancelFunc
	logger    *utils.Logger
	config    *utils.ConfigManager
	metrics   *utils.MetricsManager
	mu        sync.RWMutex
	pingTimes map[peer.ID]time.Time // Last ping time for each peer
}

// Options represents configuration options for the NetworkManager.
type Options struct {
	ListenAddrs []multiaddr.Multiaddr
	PrivKeyPath string
	MaxPeers    int
}

// DefaultOptions returns default network manager options.
func DefaultOptions(config *utils.ConfigManager) (*Options, error) {
	port := config.GetInt("p2p.port")
	if port == 0 {
		port = 9000 // Default port if not specified
	}

	addr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to create listen address: %w", err)
	}

	maxPeers := config.GetInt("p2p.maxPeers")
	if maxPeers == 0 {
		maxPeers = 50 // Default if not specified
	}

	// Create a private key path in the node's data directory
	dataDir := config.GetString("node.datadir")
	if dataDir == "" {
		dataDir = "./data" // Default data directory
	}
	
	nodeID := config.GetString("node.id")
	if nodeID == "" {
		nodeID = "node" // Default node identifier
	}
	
	privKeyPath := filepath.Join(dataDir, fmt.Sprintf("%s_key", nodeID))

	return &Options{
		ListenAddrs: []multiaddr.Multiaddr{addr},
		PrivKeyPath: privKeyPath,
		MaxPeers:    maxPeers,
	}, nil
}

// NewNetworkManager creates a new P2P network manager.
func NewNetworkManager(config *utils.ConfigManager, logger *utils.Logger) (*NetworkManager, error) {
	// Validate required dependencies
	if config == nil {
		return nil, fmt.Errorf("configuration manager is required")
	}
	if logger == nil {
		logger = utils.GetGlobalLogger()
	}

	// Get options with defaults
	opts, err := DefaultOptions(config)
	if err != nil {
		return nil, err
	}

	// Create context
	ctx, cancel := context.WithCancel(context.Background())

	// Create metrics collector
	metrics := utils.GetGlobalMetrics()

	// Create NetworkManager instance
	nm := &NetworkManager{
		ctx:       ctx,
		cancel:    cancel,
		logger:    logger.WithComponent("p2p-network"),
		config:    config,
		metrics:   metrics,
		pingTimes: make(map[peer.ID]time.Time),
	}

	// Initialize libp2p host
	if err := nm.initHost(opts); err != nil {
		cancel()
		return nil, err
	}

	// Log successful initialization
	nm.logger.Info("P2P network manager initialized",
		"id", nm.host.ID().String(),
		"addresses", nm.host.Addrs())

	// Start background tasks
	go nm.startPingTicker(ctx)

	return nm, nil
}

// initHost initializes the libp2p host with the given options.
func (nm *NetworkManager) initHost(opts *Options) error {
    // Generate or load private key
    priv, err := nm.getPrivateKey(opts.PrivKeyPath)
    if err != nil {
        return fmt.Errorf("failed to get private key: %w", err)
    }

    // Setup connection manager
    connManager, err := connmgr.NewConnManager(
        opts.MaxPeers/2, // Low watermark
        opts.MaxPeers,   // High watermark
        connmgr.WithGracePeriod(time.Minute),
    )
    if err != nil {
        return fmt.Errorf("failed to create connection manager: %w", err)
    }

    // Create libp2p host
    host, err := libp2p.New(
        libp2p.ListenAddrs(opts.ListenAddrs...),
        libp2p.Identity(priv),
        libp2p.ConnectionManager(connManager),
        libp2p.NATPortMap(),
        libp2p.EnableRelay(),
    )
    if err != nil {
        return fmt.Errorf("failed to create libp2p host: %w", err)
    }

    nm.host = host

	host.Network().Notify(&network.NotifyBundle{
    ConnectedF: func(net network.Network, conn network.Conn) {
        remotePeer := conn.RemotePeer()
        nm.logger.Debug("New connection established",
            "peer", remotePeer.String(),
            "direction", conn.Stat().Direction.String())
        nm.metrics.IncCounter("p2p_connections_total", conn.Stat().Direction.String())
        nm.metrics.SetGauge("p2p_peers_connected", float64(len(host.Network().Peers())))
    },
    // You can implement other callback functions (DisconnectedF, etc.) as needed.
    })

    return nil
}

// getPrivateKey gets the private key for the node, generating a new one if needed.
func (nm *NetworkManager) getPrivateKey(keyPath string) (crypto.PrivKey, error) {
	// Try to load the key from file first
	if keyPath != "" {
		if err := os.MkdirAll(filepath.Dir(keyPath), 0755); err != nil {
			return nil, fmt.Errorf("failed to create directory for key: %w", err)
		}

		if _, err := os.Stat(keyPath); err == nil {
			// Key file exists, try to load it
			keyData, err := ioutil.ReadFile(keyPath)
			if err != nil {
				return nil, fmt.Errorf("failed to read key file: %w", err)
			}

			// Decode the key from base64
			keyBytes, err := base64.StdEncoding.DecodeString(string(keyData))
			if err != nil {
				return nil, fmt.Errorf("failed to decode key from base64: %w", err)
			}

			// Unmarshal private key
			priv, err := crypto.UnmarshalPrivateKey(keyBytes)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal private key: %w", err)
			}

			nm.logger.Info("Loaded existing private key", "path", keyPath)
			return priv, nil
		}
	}

	// No key found or error reading it, generate a new one
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, -1, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	// Save the key if a path is provided
	if keyPath != "" {
		// Marshal the private key
		keyBytes, err := crypto.MarshalPrivateKey(priv)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal private key: %w", err)
		}

		// Encode to base64 for storage
		keyData := []byte(base64.StdEncoding.EncodeToString(keyBytes))

		// Write to file
		if err := ioutil.WriteFile(keyPath, keyData, 0600); err != nil {
			return nil, fmt.Errorf("failed to write key file: %w", err)
		}

		nm.logger.Info("Generated and saved new private key", "path", keyPath)
	}

	return priv, nil
}

// startPingTicker starts a ticker to periodically ping peers.
func (nm *NetworkManager) startPingTicker(ctx context.Context) {
	// Get ping interval from config
	interval := nm.config.GetDuration("p2p.pingInterval")
	if interval == 0 {
		interval = 5 * time.Minute // Default to 5 minutes
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			nm.pingAllPeers()
		case <-ctx.Done():
			return
		}
	}
}

// pingAllPeers pings all connected peers.
func (nm *NetworkManager) pingAllPeers() {
	peers := nm.host.Network().Peers()
	nm.logger.Debug("Pinging peers", "count", len(peers))

	for _, p := range peers {
		go func(id peer.ID) {
			if err := nm.PingPeer(id); err != nil {
				nm.logger.Debug("Failed to ping peer",
					"peer", id.String(),
					"error", err)
			}
		}(p)
	}
}

// PingPeer pings a specific peer and updates the ping time.
func (nm *NetworkManager) PingPeer(id peer.ID) error {
	// Implement actual ping using libp2p ping service
	// For MVP, we'll just update the ping time

	nm.mu.Lock()
	nm.pingTimes[id] = time.Now()
	nm.mu.Unlock()

	nm.logger.Debug("Peer pinged", "peer", id.String())
	return nil
}

// Connect connects to a peer using its multiaddress.
func (nm *NetworkManager) Connect(addr multiaddr.Multiaddr) error {
	peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return fmt.Errorf("invalid peer address: %w", err)
	}

	if err := nm.host.Connect(nm.ctx, *peerInfo); err != nil {
		return fmt.Errorf("failed to connect to peer: %w", err)
	}

	nm.logger.Info("Connected to peer", "peer", peerInfo.ID.String())
	return nil
}

// Disconnect disconnects from a peer.
func (nm *NetworkManager) Disconnect(id peer.ID) error {
	if nm.host.Network().Connectedness(id) != network.Connected {
		return fmt.Errorf("not connected to peer %s", id.String())
	}

	if err := nm.host.Network().ClosePeer(id); err != nil {
		return fmt.Errorf("failed to close connection to peer: %w", err)
	}

	nm.logger.Info("Disconnected from peer", "peer", id.String())
	return nil
}

// GetPeers returns the list of connected peers.
func (nm *NetworkManager) GetPeers() []peer.ID {
	return nm.host.Network().Peers()
}

// IsPeerOnline checks if a peer is online based on recent pings.
func (nm *NetworkManager) IsPeerOnline(id peer.ID) bool {
	if nm.host.Network().Connectedness(id) != network.Connected {
		return false
	}

	nm.mu.RLock()
	pingTime, ok := nm.pingTimes[id]
	nm.mu.RUnlock()

	if !ok {
		return false
	}

	// Get ping timeout from config (default to 10 minutes)
	pingTimeout := nm.config.GetDuration("p2p.pingTimeout")
	if pingTimeout == 0 {
		pingTimeout = 10 * time.Minute
	}

	return time.Since(pingTime) < pingTimeout
}

// GetHost returns the libp2p host instance.
func (nm *NetworkManager) GetHost() host.Host {
	return nm.host
}

// Close shuts down the network manager and closes all connections.
func (nm *NetworkManager) Close() error {
	nm.cancel()
	return nm.host.Close()
}

// GetMultiaddrs returns the full multiaddresses of this node, including peer ID
func (nm *NetworkManager) GetMultiaddrs() []string {
	hostID := nm.host.ID()
	addrs := nm.host.Addrs()
	
	var fullAddrs []string
	for _, addr := range addrs {
		fullAddr := fmt.Sprintf("%s/p2p/%s", addr.String(), hostID.String())
		fullAddrs = append(fullAddrs, fullAddr)
	}
	
	return fullAddrs
}