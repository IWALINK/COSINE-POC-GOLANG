package p2p

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
	 drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-kad-dht/dual"
	"github.com/multiformats/go-multiaddr"
)

// DiscoveryService manages peer discovery for COSINE validators.
type DiscoveryService struct {
	network       *NetworkManager
	logger        *utils.Logger
	config        *utils.ConfigManager
	metrics       *utils.MetricsManager
	dht           *dual.DHT
	ctx           context.Context
	cancel        context.CancelFunc
	bootstrapLock sync.Mutex
	bootstrapped  bool
}

// NewDiscoveryService creates a new peer discovery service.
func NewDiscoveryService(
	network *NetworkManager,
	config *utils.ConfigManager,
	logger *utils.Logger,
) (*DiscoveryService, error) {
	if network == nil {
		return nil, fmt.Errorf("network manager is required")
	}
	if config == nil {
		return nil, fmt.Errorf("configuration manager is required")
	}
	if logger == nil {
		logger = utils.GetGlobalLogger()
	}

	ctx, cancel := context.WithCancel(context.Background())

	// Create discovery service instance
	service := &DiscoveryService{
		network:       network,
		logger:        logger.WithComponent("p2p-discovery"),
		config:        config,
		metrics:       utils.GetGlobalMetrics(),
		ctx:           ctx,
		cancel:        cancel,
		bootstrapped:  false,
	}

	// Initialize the DHT
	if err := service.initDHT(); err != nil {
		cancel()
		return nil, err
	}

	return service, nil
}

// initDHT initializes the Kademlia DHT for peer discovery.
func (ds *DiscoveryService) initDHT() error {
	// Get host from network manager
	host := ds.network.GetHost()

	// Create DHT instance
	d, err := dual.New(ds.ctx, host,
		dual.DHTOption(dht.Mode(dht.ModeServer)),
		dual.DHTOption(dht.BootstrapPeers(ds.getBootstrapNodes()...)),
	)
	if err != nil {
		return fmt.Errorf("failed to create DHT: %w", err)
	}

	ds.dht = d
	ds.logger.Info("DHT initialized")

	return nil
}

// getBootstrapNodes gets bootstrap nodes from configuration.
func (ds *DiscoveryService) getBootstrapNodes() []peer.AddrInfo {
	bootstrapNodesStr := ds.config.GetStringSlice("p2p.bootstrapNodes")
	if len(bootstrapNodesStr) == 0 {
		return nil
	}

	var bootstrapNodes []peer.AddrInfo
	for _, addrStr := range bootstrapNodesStr {
		addr, err := multiaddr.NewMultiaddr(addrStr)
		if err != nil {
			ds.logger.Warn("Invalid bootstrap node address", "address", addrStr, "error", err)
			continue
		}

		peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			ds.logger.Warn("Failed to parse peer info from address", "address", addrStr, "error", err)
			continue
		}

		bootstrapNodes = append(bootstrapNodes, *peerInfo)
	}

	ds.logger.Info("Loaded bootstrap nodes", "count", len(bootstrapNodes))
	return bootstrapNodes
}

// Bootstrap initiates the peer discovery process.
func (ds *DiscoveryService) Bootstrap() error {
	ds.bootstrapLock.Lock()
	defer ds.bootstrapLock.Unlock()

	if ds.bootstrapped {
		return nil
	}

	ds.logger.Info("Starting bootstrap process")
	
	// Connect to bootstrap nodes
	bootstrapNodes := ds.getBootstrapNodes()
	if len(bootstrapNodes) > 0 {
		var wg sync.WaitGroup
		
		for _, peerInfo := range bootstrapNodes {
			wg.Add(1)
			go func(pi peer.AddrInfo) {
				defer wg.Done()
				
				if err := ds.network.GetHost().Connect(ds.ctx, pi); err != nil {
					ds.logger.Warn("Failed to connect to bootstrap node", 
						"peer", pi.ID.String(),
						"error", err)
				} else {
					ds.logger.Info("Connected to bootstrap node", "peer", pi.ID.String())
				}
			}(peerInfo)
		}
		
		// Wait for connection attempts to complete (with timeout)
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()
		
		select {
		case <-done:
			ds.logger.Info("Bootstrap node connections complete")
		case <-time.After(30 * time.Second):
			ds.logger.Warn("Bootstrap connection timeout")
		}
	}

	// Bootstrap the DHT
	if err := ds.dht.Bootstrap(ds.ctx); err != nil {
		return fmt.Errorf("failed to bootstrap DHT: %w", err)
	}

	ds.bootstrapped = true
	ds.logger.Info("Bootstrap process completed")
	
	// Start periodic peer discovery
	go ds.startDiscoveryLoop()
	
	return nil
}

// startDiscoveryLoop periodically discovers new peers.
func (ds *DiscoveryService) startDiscoveryLoop() {
	// Get discovery interval from config (default to 5 minutes)
	interval := ds.config.GetDuration("p2p.discoveryInterval")
	if interval == 0 {
		interval = 5 * time.Minute
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ds.discoverPeers()
		case <-ds.ctx.Done():
			return
		}
	}
}

// discoverPeers performs a single round of peer discovery.
func (ds *DiscoveryService) discoverPeers() {
	start := time.Now()
	ds.logger.Debug("Starting peer discovery round")

	// Find peers in the network
	routingDiscovery := drouting.NewRoutingDiscovery(ds.dht.WAN)
	
	// Use the validator-node as the discovery namespace
	routingDiscovery.Advertise(ds.ctx, "/cosine/validator-node")
	
	// Find other peers advertising the same service
	peerChan, err := routingDiscovery.FindPeers(ds.ctx, "/cosine/validator-node")
	if err != nil {
		ds.logger.Error("Failed to find peers", "error", err)
		return
	}

	count := 0
	maxPeers := ds.config.GetInt("p2p.maxPeers")
	currentPeerCount := len(ds.network.GetPeers())
	roomForNewPeers := maxPeers - currentPeerCount
	
	if roomForNewPeers <= 0 {
		ds.logger.Debug("Max peer count reached, skipping peer discovery")
		return
	}

	// Connect to discovered peers
	for p := range peerChan {
		// Skip self or already connected peers
		if p.ID == ds.network.GetHost().ID() || ds.network.GetHost().Network().Connectedness(p.ID) == 1 {
			continue
		}

		// Connect to the peer
		if err := ds.network.GetHost().Connect(ds.ctx, p); err != nil {
			ds.logger.Debug("Failed to connect to discovered peer",
				"peer", p.ID.String(),
				"error", err)
			continue
		}

		ds.logger.Debug("Connected to discovered peer", "peer", p.ID.String())
		count++
		
		// Limit the number of new connections
		if count >= roomForNewPeers {
			break
		}
	}

	duration := time.Since(start)
	ds.logger.Debug("Peer discovery round completed",
		"duration", duration,
		"new_connections", count)
	
	// Update metrics
	ds.metrics.SetGauge("p2p_discovery_duration_seconds", duration.Seconds())
	ds.metrics.SetGauge("p2p_peers_connected", float64(len(ds.network.GetPeers())))
}

// FindValidator tries to find a specific validator node by ID.
func (ds *DiscoveryService) FindValidator(validatorID string) (peer.ID, error) {
	// Convert validator ID to peer ID (this could be a custom mapping)
	// For MVP, assume validator ID is the peer ID
	peerID, err := peer.Decode(validatorID)
	if err != nil {
		return "", fmt.Errorf("invalid validator ID: %w", err)
	}

	// Check if peer is already connected
	if ds.network.GetHost().Network().Connectedness(peerID) == 1 {
		return peerID, nil
	}

	// Try to find the peer in the DHT
	ctx, cancel := context.WithTimeout(ds.ctx, 30*time.Second)
	defer cancel()

	peerInfo, err := ds.dht.WAN.FindPeer(ctx, peerID)
	if err != nil {
		return "", fmt.Errorf("validator not found: %w", err)
	}

	// Connect to the peer
	if err := ds.network.GetHost().Connect(ctx, peerInfo); err != nil {
		return "", fmt.Errorf("failed to connect to validator: %w", err)
	}

	return peerID, nil
}

// AdvertiseValidator announces this node as a validator.
func (ds *DiscoveryService) AdvertiseValidator() error {
	routingDiscovery := drouting.NewRoutingDiscovery(ds.dht.WAN)
	
	// Use the validator-node as the discovery namespace
	ttl, err := routingDiscovery.Advertise(ds.ctx, "/cosine/validator-node")
	if err != nil {
		return fmt.Errorf("failed to advertise as validator: %w", err)
	}

	ds.logger.Info("Advertised as validator node", "ttl", ttl)
	return nil
}

// GetDHT returns the dual DHT instance.
func (ds *DiscoveryService) GetDHT() *dual.DHT {
	return ds.dht
}

// Close shuts down the discovery service.
func (ds *DiscoveryService) Close() error {
	ds.cancel()
	if ds.dht != nil {
		return ds.dht.Close()
	}
	return nil
}
