package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
    "github.com/libp2p/go-libp2p/core/peer"
	"github.com/IWALINK/cosine/internal/utils"
	"github.com/IWALINK/cosine/pkg/p2p"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	// Initialize configuration
	if err := utils.InitGlobalConfig("config.yaml", "development"); err != nil {
		fmt.Printf("Failed to initialize configuration: %v\n", err)
		os.Exit(1)
	}
	config := utils.GetGlobalConfig()

	// Setup logging
	logOptions := &utils.LoggerOptions{
		Level:       utils.DebugLevel, // Using debug level for more detailed logs
		OutputPaths: []string{"stdout", "../../logs/cosine.log"},
		Encoding:    "console",
		Development: true,
	}
	if err := utils.InitGlobalLogger(logOptions); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	logger := utils.GetGlobalLogger()

	// Setup metrics
	metricsOptions := &utils.MetricsOptions{
		Address: ":9090",
		Path:    "/metrics",
	}
	if err := utils.InitGlobalMetrics(metricsOptions); err != nil {
		logger.Error("Failed to initialize metrics", "error", err)
		os.Exit(1)
	}
	metrics := utils.GetGlobalMetrics()

	// Log startup information
	logger.Info("COSINE P2P Test CLI starting up",
		"version", "0.1.0",
		"environment", config.GetString("node.environment"),
	)

	// Initialize P2P Network Manager
	networkManager, err := p2p.NewNetworkManager(config, logger)
	if err != nil {
		logger.Error("Failed to initialize network manager", "error", err)
		os.Exit(1)
	}

	// Initialize Discovery Service
	discoveryService, err := p2p.NewDiscoveryService(networkManager, config, logger)
	if err != nil {
		logger.Error("Failed to initialize discovery service", "error", err)
		os.Exit(1)
	}

	// Bootstrap the discovery service
	if err := discoveryService.Bootstrap(); err != nil {
		logger.Error("Failed to bootstrap discovery service", "error", err)
		os.Exit(1)
	}

	// Initialize Messaging Service
	messagingService, err := p2p.NewMessagingService(networkManager, config, logger)
	if err != nil {
		logger.Error("Failed to initialize messaging service", "error", err)
		os.Exit(1)
	}

	// Register message handlers
	registerMessageHandlers(messagingService, logger)

	// Advertise as validator if enabled
	if err := discoveryService.AdvertiseValidator(); err != nil {
		logger.Warn("Failed to advertise as validator", "error", err)
	} else {
		logger.Info("Successfully advertised as validator")
	}

	// Connect to bootstrap nodes (if any are configured)
	connectToBootstrapNodes(networkManager, config, logger)

	// Start periodic tasks (in a separate goroutine)
	go runPeriodicTasks(networkManager, messagingService, logger, metrics)

	// Record some metrics
	metrics.IncCounter("votes_total", "positive")
	metrics.SetGauge("peers_count", float64(len(networkManager.GetPeers())))

	// Setup graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Keep running until signal
	logger.Info("COSINE P2P Test CLI running. Press Ctrl+C to stop.")
	<-sigCh

	// Graceful shutdown
	logger.Info("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Close P2P services
	if err := messagingService.Close(); err != nil {
		logger.Error("Error closing messaging service", "error", err)
	}

	if err := discoveryService.Close(); err != nil {
		logger.Error("Error closing discovery service", "error", err)
	}

	if err := networkManager.Close(); err != nil {
		logger.Error("Error closing network manager", "error", err)
	}

	if err := metrics.StopServer(ctx); err != nil {
		logger.Error("Error stopping metrics server", "error", err)
	}

	if err := logger.Close(); err != nil {
		fmt.Printf("Error closing logger: %v\n", err)
	}
}

// registerMessageHandlers registers handlers for different message types
func registerMessageHandlers(ms *p2p.MessagingService, logger *utils.Logger) {
	// Handler for vote messages
	ms.RegisterHandler(p2p.TypeVote, func(msg *p2p.Message, from peer.ID) error {
	  logger.Info("Received vote message",
		"from", from.String(),
		"target", msg.Target,
		"value", msg.Payload["value"],
		"reputation", msg.Payload["reputation"])
	  return nil
    })

    // Handler for credit score messages
	ms.RegisterHandler(p2p.TypeCreditScore, func(msg *p2p.Message, from peer.ID) error {
	   logger.Info("Received credit score message",
		"from", from.String(),
		"target", msg.Target,
		"score", msg.Payload["score"],
		"reason", msg.Payload["reason"])
	   return nil
    })



	// Handler for validator info messages
	 ms.RegisterHandler(p2p.TypeValidatorInfo, func(msg *p2p.Message, from peer.ID) error {
	   logger.Info("Received validator info message",
		"from", from.String(),
		"validator_id", msg.Payload["validator_id"],
		"stake", msg.Payload["stake"],
		"performance", msg.Payload["performance"])
	   return nil
    })


	// Handler for ping messages
	ms.RegisterHandler(p2p.TypePing, func(msg *p2p.Message, from peer.ID) error {
		logger.Info("Received ping message", "from", from.String())

		// Send pong response
		pongMsg := &p2p.Message{
			Type:      p2p.TypePong,
			Timestamp: time.Now().UnixNano(),
			Payload:   map[string]interface{}{},
		}

		// Send direct message without conversion since 'from' is already a peer.ID
		if err := ms.SendDirect(from, pongMsg); err != nil {
			logger.Error("Failed to send pong response", "error", err)
			return err
		}

		logger.Info("Sent pong response", "to", from.String())
		return nil
    })


	// Handler for pong messages
	ms.RegisterHandler(p2p.TypePong, func(msg *p2p.Message, from peer.ID) error {
		logger.Info("Received pong message", "from", from.String())
		return nil
    })

}

// connectToBootstrapNodes attempts to connect to configured bootstrap nodes
func connectToBootstrapNodes(nm *p2p.NetworkManager, config *utils.ConfigManager, logger *utils.Logger) {
	bootstrapNodes := config.GetStringSlice("p2p.bootstrapNodes")

	if len(bootstrapNodes) == 0 {
		logger.Info("No bootstrap nodes configured")
		return
	}

	for _, node := range bootstrapNodes {
		addr, err := multiaddr.NewMultiaddr(node)
		if err != nil {
			logger.Error("Invalid bootstrap node address", "address", node, "error", err)
			continue
		}

		logger.Info("Attempting to connect to bootstrap node", "address", node)
		if err := nm.Connect(addr); err != nil {
			logger.Warn("Failed to connect to bootstrap node", "address", node, "error", err)
		} else {
			logger.Info("Connected to bootstrap node", "address", node)
		}
	}
}

// runPeriodicTasks performs periodic P2P tasks
func runPeriodicTasks(nm *p2p.NetworkManager, ms *p2p.MessagingService, logger *utils.Logger, metrics *utils.MetricsManager) {
	// Create ticker for periodic broadcasts (every 30 seconds)
	broadcastTicker := time.NewTicker(30 * time.Second)
	defer broadcastTicker.Stop()

	// Create ticker for peer stats (every 10 seconds)
	statsTicker := time.NewTicker(10 * time.Second)
	defer statsTicker.Stop()

	for {
		select {
		case <-broadcastTicker.C:
			// Broadcast a test vote message
			err := ms.BroadcastVote("0xTestWallet123", 1, 0.85)
			if err != nil {
				logger.Error("Failed to broadcast vote", "error", err)
			} else {
				logger.Info("Broadcast test vote message")
				metrics.IncCounter("votes_total", "positive")
			}

			// Broadcast a test credit score message
			err = ms.BroadcastCreditScore("0xTestWallet456", 75.5, "periodic update")
			if err != nil {
				logger.Error("Failed to broadcast credit score", "error", err)
			} else {
				logger.Info("Broadcast test credit score message")
				metrics.IncCounter("credit_score_updates_total", "periodic")
			}

		case <-statsTicker.C:
			// Log connected peers
			peers := nm.GetPeers()
			logger.Info("Connected peers", "count", len(peers))
			metrics.SetGauge("peers_count", float64(len(peers)))

			// Try to ping a peer if we have any
			if len(peers) > 0 {
				// Pick the first peer
				peer := peers[0]

				// Try to ping the peer
				latency, err := ms.Ping(peer)
				if err != nil {
					logger.Warn("Failed to ping peer", "peer", peer.String(), "error", err)
				} else {
					logger.Info("Pinged peer", "peer", peer.String(), "latency", latency)
					metrics.ObserveHistogram("validator_ping_latency_seconds", latency.Seconds(), peer.String())
				}
			}
		}
	}
}
