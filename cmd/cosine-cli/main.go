
//###############################################################################################################################################
/*TEST UTILS*/

/*package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
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
		Level:       utils.InfoLevel,
		OutputPaths: []string{"stdout", "./logs/cosine.log"},
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
	logger.Info("COSINE CLI starting up",
		"version", "0.1.0",
		"environment", config.GetString("node.environment"),
	)

	// Test math utilities
	v1 := utils.VectorFromValues(1.0, 2.0, 3.0)
	v2 := utils.VectorFromValues(4.0, 5.0, 6.0)
	similarity := utils.CosineSimilarity(v1, v2)
	logger.Info("Vector similarity test", "similarity", similarity)

	// Test Kalman filter
	kf := utils.NewKalmanFilter(0.0, 1.0, 0.01, 1.0)
	measurements := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	for _, m := range measurements {
		filtered := kf.Update(m)
		logger.Info("Kalman filter test", "measurement", m, "filtered", filtered)
	}

	// Record some metrics
	metrics.IncCounter("votes_total", "positive")
	metrics.SetGauge("peers_count", 5)
	metrics.ObserveHistogram("api_request_duration_seconds", 0.42, "creditCheck", "GET")

	// Setup graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Keep running until signal
	logger.Info("COSINE CLI running. Press Ctrl+C to stop.")
	<-sigCh

	// Graceful shutdown
	logger.Info("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := metrics.StopServer(ctx); err != nil {
		logger.Error("Error stopping metrics server", "error", err)
	}

	if err := logger.Close(); err != nil {
		fmt.Printf("Error closing logger: %v\n", err)
	}
}*/

//###############################################################################################################################################
/*TEST DATA STRUCTURES*/
/*package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
    "math"
	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
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
		Level:       utils.InfoLevel,
		OutputPaths: []string{"stdout", "./logs/cosine.log"},
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
	logger.Info("COSINE CLI starting up",
		"version", "0.1.0",
		"environment", config.GetString("node.environment"),
	)

	// Initialize storage systems
	wst := storage.NewWalletStateTrie(logger)
	gso := storage.NewGlobalStateObject(logger)
	l1t := storage.NewLinkedL1Trie(logger, wst)

	logger.Info("Storage systems initialized",
		"wst_root", wst.GetRoot(),
		"gso_hash", gso.GetHash(),
		"l1t_root", l1t.GetRoot())

	// Test wallet state operations
	testWalletAddr := "0x1234567890abcdef1234567890abcdef12345678"
	wallet := wst.GetWallet(testWalletAddr)
	wallet.ReputationScore = 0.75
	wallet.CreditScore = 450.0
	wallet.AddLinkedL1Address(1, "0xEthereumAddress")
	wallet.AddLinkedL1Address(501, "SolanaAddress")
	wallet.ReceiveTokens(1000)
	wst.UpdateWallet(wallet)

	logger.Info("Created test wallet",
		"address", testWalletAddr,
		"reputation", wallet.ReputationScore,
		"credit_score", wallet.CreditScore,
		"token_balance", wallet.TokenBalance,
		"linked_addresses", len(wallet.LinkedL1Addresses))

	// Test L1 mapping
	err := l1t.LinkAddress(1, "0xEthereumAddress", testWalletAddr)
	if err != nil {
		logger.Error("Error linking address", "error", err)
	} else {
		logger.Info("Successfully linked Ethereum address",
			"eth_address", "0xEthereumAddress",
			"l2_address", testWalletAddr)
	}

	// Test Global State Object operations
	// Initialize validator
	validatorAddr := "0xValidator1"
	gso.AddValidator(validatorAddr, 150000, 0.9)

	// Test dynamic scaling factor updates
	gso.UpdateScalingFactor("K_neg", 0.95, 1.0, 0.01, 1.0)
	gso.UpdateScalingFactor("K_pos", 1.05, 1.0, 0.01, 1.0)

	logger.Info("Updated scaling factors",
		"k_neg", gso.GetScalingFactor("K_neg"),
		"k_pos", gso.GetScalingFactor("K_pos"))

	// Test credit score updates using dynamic scaling
	kNeg := gso.GetScalingFactor("K_neg")
	deltaVote := -5.0 // Negative vote effect
	wallet.UpdateCreditScore(deltaVote * kNeg)
	wst.UpdateWallet(wallet)

	logger.Info("Updated wallet credit score",
		"new_score", wallet.CreditScore,
		"delta", deltaVote * kNeg)

	// Test association risk
	wallet.UpdateAssociationRisk(1.5)
	wst.UpdateWallet(wallet)

	logger.Info("Updated association risk",
		"association_risk", wallet.AssociationRiskScore)

	// Test cosine similarity for credit verification
	creditVector := wallet.GetNormalizedCreditVector(0, 100)
	thresholdVector := utils.VectorFromValues(3.5, 1)
	similarity := utils.CosineSimilarity(creditVector, thresholdVector)

	logger.Info("Credit verification using cosine similarity",
		"similarity", similarity,
		"threshold", 0.95,
		"verified", similarity >= 0.95)

	// Record some metrics
	metrics.IncCounter("wallet_count")
	metrics.SetGauge("peers_count", 5)
	metrics.IncCounter("credit_score_updates_total", "vote")
	metrics.ObserveHistogram("credit_score_shift_magnitude", math.Abs(deltaVote * kNeg), "vote")

	// Verify consistency between L1T and WST
	consistent, inconsistencies := l1t.VerifyConsistency()
	if !consistent {
		logger.Error("Storage inconsistency detected", "issues", inconsistencies)
	} else {
		logger.Info("Storage consistency verified")
	}

	// Setup graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Keep running until signal
	logger.Info("COSINE CLI running. Press Ctrl+C to stop.")
	<-sigCh

	// Graceful shutdown
	logger.Info("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := metrics.StopServer(ctx); err != nil {
		logger.Error("Error stopping metrics server", "error", err)
	}

	if err := logger.Close(); err != nil {
		fmt.Printf("Error closing logger: %v\n", err)
	}
}*/



//###############################################################################################################################################
/*TEST P2P NETWORKING*/
/*package main

import (
	"context"
	"fmt"
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/IWALINK/cosine/internal/utils"
	"github.com/IWALINK/cosine/pkg/p2p"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	// Parse command line arguments
	configPath := flag.String("config", "config.yaml", "Path to configuration file")
	clearRegistry := flag.Bool("clear-registry", false, "Clear the peer registry on startup")
	flag.Parse()

	// Initialize configuration
	if err := utils.InitGlobalConfig(*configPath, "development"); err != nil {
		fmt.Printf("Failed to initialize configuration: %v\n", err)
		os.Exit(1)
	}
	config := utils.GetGlobalConfig()

	// Set up node ID suffix for logs
	nodeID := config.GetString("node.id")
	if nodeID == "" {
		fmt.Println("Node ID not specified in config, using default 'node'")
		nodeID = "node"
	}

	// Setup logging
	logOptions := &utils.LoggerOptions{
		Level:       utils.DebugLevel,
		OutputPaths: []string{"stdout", fmt.Sprintf("./logs/cosine-%s.log", nodeID)},
		Encoding:    "console",
		Development: true,
	}
	if err := utils.InitGlobalLogger(logOptions); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	logger := utils.GetGlobalLogger()

	// Setup metrics with unique port based on node ID
	metricsPort := config.GetString("metrics.address")
	metricsOptions := &utils.MetricsOptions{
		Address: metricsPort,
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
		"node", nodeID,
		"environment", config.GetString("node.environment"),
	)

	// Initialize peer registry
	// Use a common shared registry path from configuration.
	// Add a new key "peerRegistry.path" in your config files (e.g., "data/peer_registry.json").
	registryPath := config.GetString("peerRegistry.path")
	if registryPath == "" {
		registryPath = "data/peer_registry.json" // default shared location in project root
	}
	if err := os.MkdirAll(filepath.Dir(registryPath), 0755); err != nil {
		logger.Error("Failed to create peer registry directory", "error", err)
		os.Exit(1)
	}
	peerRegistry, err := utils.NewPeerRegistry(registryPath, logger)
	if err != nil {
		logger.Error("Failed to initialize peer registry", "error", err)
		os.Exit(1)
	}

	// Clear registry if requested
	if *clearRegistry {
		logger.Info("Clearing peer registry as requested")
		if err := peerRegistry.Clear(); err != nil {
			logger.Error("Failed to clear peer registry", "error", err)
		}
	}

	// Initialize P2P Network Manager
	networkManager, err := p2p.NewNetworkManager(config, logger)
	if err != nil {
		logger.Error("Failed to initialize network manager", "error", err)
		os.Exit(1)
	}

	// Get node's multiaddresses with peer ID
	nodeAddrs := networkManager.GetMultiaddrs()
	
	// Register this node's addresses in the peer registry
	if err := peerRegistry.RegisterPeer(nodeID, nodeAddrs); err != nil {
		logger.Error("Failed to register peer addresses", "error", err)
	}

	// Display peer ID for this node
	logger.Info("Node initialized with Peer ID",
		"peer_id", networkManager.GetHost().ID().String(),
		"addresses", nodeAddrs,
	)

	// Override bootstrap nodes from registry
	bootstrapAddrs := peerRegistry.GetBootstrapAddrs(nodeID)
	if len(bootstrapAddrs) > 0 {
		logger.Info("Using bootstrap nodes from registry", "addresses", bootstrapAddrs)
	} else {
		logger.Info("No peers in registry yet, using bootstrap nodes from config")
		bootstrapAddrs = config.GetStringSlice("p2p.bootstrapNodes")
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

	// Connect to bootstrap nodes (preferring registry over config)
	connectToBootstrapNodes(networkManager, bootstrapAddrs, logger)

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
func connectToBootstrapNodes(nm *p2p.NetworkManager, bootstrapNodes []string, logger *utils.Logger) {
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
		}
	}
}*/

//###############################################################################################################################################
//TEST Consensus Module
/*package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/IWALINK/cosine/internal/consensus"
	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

// SimulatedValidator represents a validator for testing
type SimulatedValidator struct {
	ID          peer.ID
	PrivKey     crypto.PrivKey
	Stake       uint64
	Performance float64
	Proposals   map[string]float64 // Wallet ID -> proposed value
	IsHonest    bool
}

// GenerateValidators creates a set of simulated validators
func GenerateValidators(count int) ([]*SimulatedValidator, error) {
	result := make([]*SimulatedValidator, count)
	
	for i := 0; i < count; i++ {
		// Generate validator key
		privKey, _, err := crypto.GenerateEd25519Key(rand.New(rand.NewSource(time.Now().UnixNano() + int64(i))))
		if err != nil {
			return nil, fmt.Errorf("failed to generate key: %w", err)
		}
		
		// Get peer ID
		id, err := peer.IDFromPrivateKey(privKey)
		if err != nil {
			return nil, fmt.Errorf("failed to get peer ID: %w", err)
		}
		
		// Set stake between 100k and 500k
		stake := uint64(100000 + rand.Intn(400000))
		
		// Set initial performance between 0.5 and 1.0
		performance := 0.5 + 0.5*rand.Float64()
		
		// Usually honest, but some validators are malicious
		isHonest := rand.Float64() <= 0.9 // 90% are honest
		
		result[i] = &SimulatedValidator{
			ID:          id,
			PrivKey:     privKey,
			Stake:       stake,
			Performance: performance,
			Proposals:   make(map[string]float64),
			IsHonest:    isHonest,
		}
	}
	
	return result, nil
}

func main() {
	fmt.Println("COSINE Consensus Module Test Application")
	fmt.Println("---------------------------------------")
	
	// Ensure data directory exists
	dataDir := "./data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Printf("Error creating data directory: %v\n", err)
		os.Exit(1)
	}
	
	// Initialize configuration
	configFile := "test_config.yaml"
	config, err := utils.NewConfigManager(configFile, "development")
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		fmt.Printf("Make sure %s exists in the configs directory\n", configFile)
		os.Exit(1)
	}
	
	// Initialize logger
	logger, err := utils.SetupLogging(config)
	if err != nil {
		fmt.Printf("Error setting up logger: %v\n", err)
		os.Exit(1)
	}
	
	// Create simulated validators
	// Store this as a variable at the function level since it's referenced multiple times
	validatorCount := 20
	validators, err := GenerateValidators(validatorCount)
	if err != nil {
		logger.Fatal("Failed to generate validators", "error", err)
	}
	
	logger.Info("Generated validators", "count", validatorCount)
	
	// Initialize test components
	err = runConsensusTesting(validators, validatorCount, config, logger)
	if err != nil {
		logger.Fatal("Test failed", "error", err)
	}
	
	logger.Info("All tests completed successfully")
}

func runConsensusTesting(validators []*SimulatedValidator, validatorCount int, config *utils.ConfigManager, logger *utils.Logger) error {
	fmt.Println("\n1. Testing VRF Functionality")
	fmt.Println("---------------------------")
	
	// Create VRF manager for our node
	ownValidator := validators[0]
	vrfManager, err := consensus.NewVRFManager(logger, ownValidator.PrivKey)
	if err != nil {
		return fmt.Errorf("failed to create VRF manager: %w", err)
	}
	
	// Generate a seed
	seed := []byte("test-seed-" + time.Now().String())
	vrfOutput, err := vrfManager.ComputeVRF(seed)
	if err != nil {
		return fmt.Errorf("failed to compute VRF: %w", err)
	}
	
	fmt.Printf("Generated VRF output with value: %f\n", vrfOutput.Value)
	
	// Verify the VRF output
	pubKey, err := ownValidator.ID.ExtractPublicKey()
	if err != nil {
		return fmt.Errorf("failed to extract public key: %w", err)
	}
	
	valid, err := consensus.VerifyVRF(pubKey, seed, vrfOutput)
	if err != nil {
		return fmt.Errorf("VRF verification error: %w", err)
	}
	
	fmt.Printf("VRF verification result: %v\n", valid)
	
	fmt.Println("\n2. Testing Validator Set Management")
	fmt.Println("----------------------------------")
	
	// Create storage components (only using gso, removing unused wst)
	gso := storage.NewGlobalStateObject(logger)
	
	// Option 1: Pass nil for the messaging service
	validatorSet := consensus.NewValidatorSet(nil, config, logger)
	
	// Add our validators to the set
	for i, validator := range validators {
		multiaddr := fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/p2p/%s", 9000+i, validator.ID.String())
		err := validatorSet.AddValidator(validator.ID, validator.Stake, validator.Performance, multiaddr)
		if err != nil {
			return fmt.Errorf("failed to add validator %d: %w", i, err)
		}
	}
	
	fmt.Printf("Added %d validators to the validator set\n", validatorSet.GetTotalCount())
	fmt.Printf("Total stake: %d tokens\n", validatorSet.GetTotalStake())
	
	// Simulate some validators going offline
	offlineCount := 3
	for i := 0; i < offlineCount; i++ {
		// We won't update the ping time for these validators, so they'll be considered offline
		idx := validatorCount - i - 1
		fmt.Printf("Validator %s is offline (not updating ping)\n", validators[idx].ID.String()[:10])
	}
	
	// Update ping times for online validators
	for i := 0; i < validatorCount-offlineCount; i++ {
		err := validatorSet.UpdateLastPing(validators[i].ID)
		if err != nil {
			return fmt.Errorf("failed to update validator ping: %w", err)
		}
	}
	
	onlineValidators := validatorSet.GetOnlineValidators()
	fmt.Printf("Online validators: %d out of %d\n", len(onlineValidators), validatorCount)
	
	fmt.Println("\n3. Testing Validator Subset Selection")
	fmt.Println("------------------------------------")
	
	// Create subset selector
	selector := consensus.NewSubsetSelector(validatorSet, vrfManager, config, logger)
	
	// Generate block seed
	blockNumber := uint64(1)
	blockTime := time.Now()
	prevBlockHash := sha256.Sum256([]byte(fmt.Sprintf("block-%d", blockNumber-1)))
	blockSeed := consensus.GenerateBlockSeed(prevBlockHash[:], blockTime)
	
	// Select validator subset
	result, err := selector.SelectValidatorSubset(blockSeed)
	if err != nil {
		return fmt.Errorf("failed to select validator subset: %w", err)
	}
	
	fmt.Printf("Selected %d validators out of %d online validators\n", 
		len(result.SelectedValidators), len(onlineValidators))
	fmt.Printf("Selection round ID: %s\n", result.RoundID)
	
	// Print top 5 selected validators
	fmt.Println("\nTop selected validators:")
	for i, id := range result.SelectedValidators {
		if i >= 5 {
			break
		}
		score := 0.0
		stake := uint64(0)
		perf := 0.0
		
		// Find score in result data
		for _, scoreData := range result.ScoreData {
			if scoreData.ID == id {
				score = scoreData.Score
				stake = scoreData.Validator.Stake
				perf = scoreData.Validator.Performance
				break
			}
		}
		
		fmt.Printf("%d. Validator %s - Score: %.6f, Stake: %d, Performance: %.2f\n", 
			i+1, id.String()[:10], score, stake, perf)
	}
	
	fmt.Println("\n4. Testing Performance and Credit Score Aggregation")
	fmt.Println("--------------------------------------------------")
	
	// Create performance manager
	performanceManager := consensus.NewPerformanceManager(validatorSet, selector, gso, config, logger)
	
	// Simulate credit score proposals for a test wallet
	testWallet := "0xTestWallet123456789"
	baseValue := 100.0
	
	// Generate proposals from validators
	proposals := make([]consensus.ProposalData, 0, len(result.SelectedValidators))
	
	fmt.Println("\nValidator proposals:")
	for i, validatorID := range result.SelectedValidators {
		// Find the validator
		var validator *SimulatedValidator
		for _, v := range validators {
			if v.ID == validatorID {
				validator = v
				break
			}
		}
		
		if validator == nil {
			continue
		}
		
		// Generate proposal (honest or malicious)
		proposedValue := baseValue
		
		if !validator.IsHonest {
			// Malicious validators propose significantly different values
			if rand.Float64() < 0.5 {
				proposedValue = baseValue * (1.5 + rand.Float64()*0.5) // 150-200% of base value
			} else {
				proposedValue = baseValue * (0.5 - rand.Float64()*0.3) // 20-50% of base value
			}
		} else {
			// Honest validators add small noise
			noise := (rand.Float64() - 0.5) * 5.0 // ±2.5% noise
			proposedValue = baseValue + noise
		}
		
		proposal := consensus.ProposalData{
			ValidatorID:   validatorID,
			ProposedValue: proposedValue,
			PreviousValue: baseValue - 10,
			Delta:         proposedValue - (baseValue - 10),
			Timestamp:     time.Now(),
			Status:        consensus.StatusInlier, // Will be determined by aggregation
		}
		
		proposals = append(proposals, proposal)
		
		fmt.Printf("%d. %s: %.2f (%s)\n", i+1, validatorID.String()[:10], 
			proposedValue, map[bool]string{true: "honest", false: "malicious"}[validator.IsHonest])
	}
	
	// Aggregate proposals
	aggregationResult, err := performanceManager.AggregateProposals(testWallet, proposals)
	if err != nil {
		return fmt.Errorf("failed to aggregate proposals: %w", err)
	}
	
	fmt.Println("\nAggregation results:")
	fmt.Printf("Mean: %.2f, StdDev: %.2f\n", aggregationResult.Mean, aggregationResult.StdDev)
	fmt.Printf("Median: %.2f, MAD: %.2f\n", aggregationResult.Median, aggregationResult.MAD)
	fmt.Printf("Tau threshold: %.2f, K threshold: %.2f\n", aggregationResult.Tau, aggregationResult.K)
	fmt.Printf("Inlier count: %d, Outlier count: %d\n", 
		aggregationResult.InlierCount, aggregationResult.OutlierCount)
	fmt.Printf("Final aggregated value: %.2f\n", aggregationResult.AggregatedValue)
	
	fmt.Println("\nOutlier filtering results:")
	// Let's check which validators were classified as outliers
	outlierCount := 0
	for i, prop := range aggregationResult.Proposals {
		outlierStatus := "inlier"
		if prop.Status == consensus.StatusOutlier {
			outlierStatus = "OUTLIER"
			outlierCount++
		}
		
		// Find if the validator is honest
		isHonest := true
		for _, v := range validators {
			if v.ID == prop.ValidatorID {
				isHonest = v.IsHonest
				break
			}
		}
		
		fmt.Printf("%d. %s: %.2f - %s (%s)\n", 
			i+1, prop.ValidatorID.String()[:10], prop.ProposedValue, 
			outlierStatus, map[bool]string{true: "honest", false: "malicious"}[isHonest])
	}
	
	// Simple validation of results
	totalMalicious := 0
	for _, v := range validators {
		if !v.IsHonest {
			totalMalicious++
		}
	}
	
	fmt.Printf("\nSummary: Detected %d outliers out of approximately %d malicious validators\n", 
		outlierCount, totalMalicious)
	
	fmt.Println("\nTest completed successfully!")
	return nil
}*/



//###############################################################################################################################################
//TEST CREDIT SCORE ENGINE
/*package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/IWALINK/cosine/internal/creditscore"
	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
)

// Configuration constants
const (
	configFile = "cosine.yaml"
	logLevel   = "debug"
)

// Test wallets
var (
	testWallets = []string{
		"0x1234567890123456789012345678901234567890",
		"0x2345678901234567890123456789012345678901",
		"0x3456789012345678901234567890123456789012",
		"0x4567890123456789012345678901234567890123",
		"0x5678901234567890123456789012345678901234",
	}
	
	testValidators = []string{
		"validator1",
		"validator2",
		"validator3",
		"validator4",
		"validator5",
		"validator6",
		"validator7",
	}
)

// simulateValidatorProposals generates test proposals for a wallet
func simulateValidatorProposals(baseScore float64, count int, outlierChance float64) []creditscore.ProposalEntry {
	proposals := make([]creditscore.ProposalEntry, count)
	now := time.Now()
	
	for i := 0; i < count; i++ {
		// Randomly select a validator
		validatorIdx := rand.Intn(len(testValidators))
		validator := testValidators[validatorIdx]
		
		// Generate a score (with small variation around the base score)
		score := baseScore + rand.Float64()*0.5 - 0.25
		
		// Potentially create an outlier
		if rand.Float64() < outlierChance {
			// Create obvious outlier (far from base score)
			score = baseScore + (rand.Float64()*10.0 - 5.0)
		}
		
		proposals[i] = creditscore.ProposalEntry{
			ValidatorID: validator,
			Score:       score,
			Timestamp:   now.Add(-time.Duration(rand.Intn(600)) * time.Second),
		}
	}
	
	return proposals
}

// runSimulations runs a series of credit score update simulations
func runSimulations(
	cum *creditscore.CreditUpdateManager,
	csm *creditscore.CreditScoreManager,
	wallets []string,
	logger *utils.Logger) {
	
	logger.Info("Starting credit score simulations")
	
	// Constants for simulations
	const (
		gamma       = 1.0 // Default exponent
		iterations  = 10   // Number of update iterations
		sleepTime   = 500  // Milliseconds between iterations
	)
	
	// 1. Simulate positive vote updates
	logger.Info("Simulating positive vote updates")
	for i := 0; i < iterations; i++ {
		for _, wallet := range wallets {
			proposals := simulateValidatorProposals(0.8, 7, 0.2)
			result := cum.ProcessVoteUpdates(wallet, proposals, creditscore.PositiveVoteUpdate, gamma)
			logger.Info("Positive vote update result",
				"wallet", wallet,
				"delta", result.Delta,
				"new_score", result.NewScore,
				"inliers", result.FilterResult.InlierCount,
				"outliers", result.FilterResult.OutlierCount)
		}
		
		// Allow time for scaling factors to adjust
		time.Sleep(time.Millisecond * sleepTime)
	}
	
	// 2. Simulate negative vote updates for some wallets
	logger.Info("Simulating negative vote updates")
	for i := 0; i < iterations; i++ {
		for _, wallet := range wallets[:2] { // Only apply to first two wallets
			proposals := simulateValidatorProposals(-0.8, 7, 0.2)
			result := cum.ProcessVoteUpdates(wallet, proposals, creditscore.NegativeVoteUpdate, gamma)
			logger.Info("Negative vote update result",
				"wallet", wallet,
				"delta", result.Delta,
				"new_score", result.NewScore,
				"inliers", result.FilterResult.InlierCount,
				"outliers", result.FilterResult.OutlierCount)
		}
		
		time.Sleep(time.Millisecond * sleepTime)
	}
	
	// 3. Simulate association risk updates for one wallet
	logger.Info("Simulating association risk updates")
	for i := 0; i < iterations; i++ {
		wallet := wallets[0]
		riskScore := 5.0 + float64(i)*0.5 // Increasing risk score
		
		result := cum.ProcessAssociationUpdate(wallet, riskScore, gamma)
		logger.Info("Association update result",
			"wallet", wallet,
			"risk_score", riskScore,
			"delta", result.Delta,
			"new_score", result.NewScore)
			
		time.Sleep(time.Millisecond * sleepTime)
	}
	
	// 4. Simulate rehabilitation for the penalized wallet
	logger.Info("Simulating rehabilitation updates")
	for i := 0; i < iterations; i++ {
		wallet := wallets[0]
		rehabVote := 0.9 + float64(i)*0.05 // Strong rehabilitation vote
		
		result := cum.ProcessRehabilitationUpdate(wallet, rehabVote, gamma)
		logger.Info("Rehabilitation update result",
			"wallet", wallet,
			"rehab_vote", rehabVote,
			"delta", result.Delta,
			"new_score", result.NewScore)
			
		time.Sleep(time.Millisecond * sleepTime)
	}
	
	// 5. Calculate and display metrics
	for _, wallet := range wallets {
		score := csm.GetCreditScore(wallet)
		logger.Info("Final credit score", "wallet", wallet, "score", score)
	}
	
	// 6. Calculate cosine similarities
	creditMetrics := csm.GetCreditMetrics()
	logger.Info("Credit metrics", 
		"mean", creditMetrics.Mean, 
		"std_dev", creditMetrics.StdDev,
		"wallet_count", creditMetrics.Count)
	
	for i, wallet1 := range wallets {
		for j, wallet2 := range wallets {
			if i < j {
				similarity := cum.GetWalletCosineSimilarity(wallet1, wallet2)
				logger.Info("Wallet similarity", 
					"wallet1", wallet1, 
					"wallet2", wallet2, 
					"similarity", similarity)
			}
		}
		
		// Compare to a threshold
		threshold := 50.0
		similarity := cum.GetCosineSimilarity(wallet1, threshold)
		isAbove := cum.IsAboveThreshold(wallet1, threshold, 0.95)
		
		logger.Info("Threshold comparison", 
			"wallet", wallet1, 
			"threshold", threshold, 
			"similarity", similarity, 
			"is_above", isAbove)
	}
	
	// 7. Test verification fees
	for _, wallet := range wallets {
		fee := cum.GetVerificationFee(wallet)
		logger.Info("Verification fee", "wallet", wallet, "fee", fee)
		
		// Process verification (pay fee and reset cost)
		newFee := cum.ProcessVerification(wallet)
		logger.Info("After verification", "wallet", wallet, "new_fee", newFee)
	}
}

func main() {
	// Parse command line flags
	var (
		configPath string
		logPath    string
		environment string
	)
	
	flag.StringVar(&configPath, "config", "", "Path to configuration file")
	flag.StringVar(&logPath, "log", "", "Path to log file")
	flag.StringVar(&environment, "env", "development", "Environment (development, testing, production)")
	flag.Parse()
	
	// Set default config path if not provided
	if configPath == "" {
		configPath = filepath.Join("configs", configFile)
	}
	
	// Initialize configuration
	config, err := utils.NewConfigManager(configFile, environment)
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}
	
	// Initialize logger
	logOptions := utils.DefaultLoggerOptions()
	logOptions.Level = utils.LogLevel(logLevel)
	
	if logPath != "" {
		logOptions.OutputPaths = append(logOptions.OutputPaths, logPath)
	}
	
	logger, err := utils.NewLogger(logOptions)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()
	
	logger.Info("Starting COSINE credit scoring test application",
		"environment", environment,
		"config", configPath)
	
	// Initialize metrics
	metrics, err := utils.SetupMetrics(config)
	if err != nil {
		logger.Warn("Failed to setup metrics", "error", err)
	}
	
	// Create storage components
	walletTrie := storage.NewWalletStateTrie(logger)
	globalState := storage.NewGlobalStateObject(logger)
	
	// Initialize credit scoring components
	kalmanManager := creditscore.NewKalmanManager(logger, config)
	dynamicScaling := creditscore.NewDynamicScaling(kalmanManager, globalState, config, logger, metrics)
	scoreManager := creditscore.NewCreditScoreManager(walletTrie, dynamicScaling, config, logger, metrics)
	outlierFilter := creditscore.NewOutlierFilter(kalmanManager, config, logger, metrics)
	creditUpdateManager := creditscore.NewCreditUpdateManager(scoreManager, outlierFilter, dynamicScaling, kalmanManager, config, logger, metrics)
	
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	// Set up context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	
	// Handle shutdown signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		sig := <-sigCh
		logger.Info("Received signal to stop", "signal", sig)
		cancel()
	}()
	
	// Run the simulations if context not cancelled
	select {
	case <-ctx.Done():
		logger.Info("Application stopping...")
	default:
		runSimulations(creditUpdateManager, scoreManager, testWallets, logger)
		logger.Info("Simulations completed successfully")
	}
	
	// Stop metrics server if running
	if metrics != nil {
		if err := metrics.StopServer(ctx); err != nil {
			logger.Error("Failed to stop metrics server", "error", err)
		}
	}
	
	logger.Info("Application stopped")
}*/

//###############################################################################################################################################
//TEST FRAUD DETECTION
/*package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IWALINK/cosine/internal/creditscore"
	"github.com/IWALINK/cosine/internal/fraud"
	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
)

const (
	// Default configuration file
	defaultConfigFile = "configs/test_fraud.yaml"
)

// Command-line flags
var (
	configFile  string
	environment string
	verbose     bool
)

func init() {
	flag.StringVar(&configFile, "config", defaultConfigFile, "Path to configuration file")
	flag.StringVar(&environment, "env", "development", "Environment (development, test, production)")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
}

func main() {
	flag.Parse()

	// Determine the subcommand to run
	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	subcommand := args[0]

	// Initialize configuration
	if err := utils.InitGlobalConfig(configFile, environment); err != nil {
		fmt.Printf("Error initializing configuration: %v\n", err)
		os.Exit(1)
	}
	config := utils.GetGlobalConfig()

	// Initialize logging
	logConfig := &utils.LoggerOptions{
		Level:       utils.LogLevel(config.GetString("logging.level")),
		OutputPaths: config.GetStringSlice("logging.outputs"),
		Encoding:    config.GetString("logging.encoding"),
		Development: environment != "production",
	}
	if err := utils.InitGlobalLogger(logConfig); err != nil {
		fmt.Printf("Error initializing logger: %v\n", err)
		os.Exit(1)
	}
	logger := utils.GetGlobalLogger()

	// Initialize metrics
	metricsConfig := &utils.MetricsOptions{
		Address:       config.GetString("metrics.address"),
		Path:          config.GetString("metrics.path"),
		DefaultLabels: map[string]string{"app": "cosine-cli", "env": environment},
	}
	if err := utils.InitGlobalMetrics(metricsConfig); err != nil {
		logger.Warn("Error initializing metrics", "error", err)
	}
	metrics := utils.GetGlobalMetrics()

	// Create components
	walletTrie := storage.NewWalletStateTrie(logger)
	randomWindows := fraud.NewRandomWindows(config, logger, metrics)

	// Set up remaining components that depend on both RandomWindows and WalletStateTrie
	associationAnalyzer := fraud.NewAssociationRiskAnalyzer(
		randomWindows,
		walletTrie,
		config,
		logger,
		metrics,
	)

	// Mock credit score components (simplified for the test)
	creditManager := createMockCreditScoreManager(walletTrie, config, logger, metrics)
	blacklistManager := fraud.NewBlacklistManager(
		associationAnalyzer,
		creditManager,
		config,
		logger,
		metrics,
	)

	// Handle different subcommands
	switch subcommand {
	case "help":
		printUsage()

	case "random-windows":
		testRandomWindows(randomWindows, args[1:])

	case "transactions":
		testTransactions(associationAnalyzer, args[1:])

	case "malicious":
		testMaliciousWallets(associationAnalyzer, args[1:])

	case "risk":
		testRiskAnalysis(associationAnalyzer, args[1:])

	case "blacklist":
		testBlacklisting(blacklistManager, args[1:])

	default:
		fmt.Printf("Unknown subcommand: %s\n", subcommand)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("COSINE Fraud Detection Test Application")
	fmt.Println("\nUsage:")
	fmt.Println("  cosine-cli [flags] <subcommand> [args]")
	fmt.Println("\nFlags:")
	fmt.Println("  -config string    Path to configuration file (default \"configs/test_fraud.yaml\")")
	fmt.Println("  -env string       Environment (development, test, production) (default \"development\")")
	fmt.Println("  -verbose          Enable verbose output")
	fmt.Println("\nSubcommands:")
	fmt.Println("  help              Show this help message")
	fmt.Println("  random-windows    Test random window generation")
	fmt.Println("    generate                  Generate random windows")
	fmt.Println("    deterministic <wallet>    Generate deterministic windows for a wallet")
	fmt.Println("  transactions      Manage test transactions")
	fmt.Println("    add <from> <to> <amount>  Add a transaction")
	fmt.Println("    list [address]            List transactions")
	fmt.Println("  malicious         Manage malicious wallets")
	fmt.Println("    add <address> <reason>    Add a malicious wallet")
	fmt.Println("    list                      List malicious wallets")
	fmt.Println("    check <address>           Check if a wallet is malicious")
	fmt.Println("  risk              Test association risk analysis")
	fmt.Println("    analyze <address>         Analyze association risk for a wallet")
	fmt.Println("    stats                     Get association risk statistics")
	fmt.Println("  blacklist         Test blacklisting functionality")
	fmt.Println("    penalize <address>        Apply association penalty to a wallet")
	fmt.Println("    check <address>           Check if a wallet is blacklisted")
	fmt.Println("    restrictions <address>    Get vote restrictions for a wallet")
	fmt.Println("    update <address> <status> Update blacklist status (partial, full, rehab)")
}

// Create a mock credit score update manager for testing
func createMockCreditScoreManager(
	walletTrie *storage.WalletStateTrie,
	config *utils.ConfigManager,
	logger *utils.Logger,
	metrics *utils.MetricsManager,
) *creditscore.CreditUpdateManager {
	// This is a simplified mock that won't actually update credit scores
	// but will provide the necessary interface for the blacklist manager
	return &creditscore.CreditUpdateManager{}
}

// Test random window generation
func testRandomWindows(rw *fraud.RandomWindows, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: missing random-windows subcommand")
		return
	}

	switch args[0] {
	case "generate":
		// Generate random windows
		params := rw.GenerateRandomWindows()
		fmt.Printf("Generated random windows:\n")
		fmt.Printf("  Time Window: %.2f days\n", params.TimeLimit.Hours()/24)
		fmt.Printf("  Hop Limit: %d\n", params.HopLimit)
		fmt.Printf("  Seed: %d\n", params.Seed)

	case "deterministic":
		if len(args) < 2 {
			fmt.Println("Error: missing wallet address")
			return
		}
		walletAddr := args[1]
		params := rw.GenerateForWallet(walletAddr, time.Now())
		fmt.Printf("Generated deterministic windows for wallet %s:\n", walletAddr)
		fmt.Printf("  Time Window: %.2f days\n", params.TimeLimit.Hours()/24)
		fmt.Printf("  Hop Limit: %d\n", params.HopLimit)
		fmt.Printf("  Seed: %d\n", params.Seed)

		// Generate for same wallet again to show determinism
		params2 := rw.GenerateForWallet(walletAddr, time.Now())
		fmt.Printf("\nGenerating again for same wallet and same trigger time should be different due to timestamp:\n")
		fmt.Printf("  Time Window: %.2f days\n", params2.TimeLimit.Hours()/24)
		fmt.Printf("  Hop Limit: %d\n", params2.HopLimit)
		fmt.Printf("  Seed: %d\n", params2.Seed)

		// Fix the timestamp for true determinism
		fixedTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		params3 := rw.GenerateForWallet(walletAddr, fixedTime)
		params4 := rw.GenerateForWallet(walletAddr, fixedTime)
		
		fmt.Printf("\nGenerating with fixed timestamp should be deterministic:\n")
		fmt.Printf("  First run:  Time Window: %.2f days, Hop Limit: %d\n", 
			params3.TimeLimit.Hours()/24, params3.HopLimit)
		fmt.Printf("  Second run: Time Window: %.2f days, Hop Limit: %d\n", 
			params4.TimeLimit.Hours()/24, params4.HopLimit)
		
	default:
		fmt.Printf("Unknown random-windows subcommand: %s\n", args[0])
	}
}

// Test transaction management
func testTransactions(ara *fraud.AssociationRiskAnalyzer, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: missing transactions subcommand")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 4 {
			fmt.Println("Error: transaction requires from, to, and amount arguments")
			return
		}
		
		fromAddr := args[1]
		toAddr := args[2]
		amount, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			fmt.Printf("Error parsing amount: %v\n", err)
			return
		}

		// Create and add transaction
		tx := fraud.Transaction{
			FromAddress:     fromAddr,
			ToAddress:       toAddr,
			Amount:          amount,
			Timestamp:       time.Now(),
			ClusteringCount: 1, // Will be updated by the analyzer
			TxHash:          fmt.Sprintf("tx-%d", time.Now().UnixNano()),
			ConfirmationTime: time.Now(),
		}
		
		ara.AddTransaction(tx)
		fmt.Printf("Added transaction from %s to %s for %.2f units\n", 
			fromAddr, toAddr, amount)

	case "list":
		// Not implemented for this simple test as we don't have direct access to transactions
		fmt.Println("Transaction listing not implemented in this simple test")
		fmt.Println("Transactions are stored internally in the AssociationRiskAnalyzer")

	default:
		fmt.Printf("Unknown transactions subcommand: %s\n", args[0])
	}
}

// Test malicious wallet management
func testMaliciousWallets(ara *fraud.AssociationRiskAnalyzer, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: missing malicious subcommand")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 3 {
			fmt.Println("Error: add requires address and reason arguments")
			return
		}
		
		address := args[1]
		reason := args[2]
		confidence := 1.0 // Default high confidence
		
		if len(args) > 3 {
			confidenceVal, err := strconv.ParseFloat(args[3], 64)
			if err == nil {
				confidence = confidenceVal
			}
		}

		ara.AddMaliciousWallet(address, reason, confidence)
		fmt.Printf("Added malicious wallet %s (Reason: %s, Confidence: %.2f)\n", 
			address, reason, confidence)

	case "check":
		if len(args) < 2 {
			fmt.Println("Error: check requires address argument")
			return
		}
		
		address := args[1]
		isMalicious := ara.IsMaliciousWallet(address)
		
		if isMalicious {
			fmt.Printf("Wallet %s is flagged as malicious\n", address)
		} else {
			fmt.Printf("Wallet %s is not flagged as malicious\n", address)
		}

	case "list":
		// Not implemented for this simple test as we don't have direct access to malicious wallets list
		fmt.Println("Malicious wallet listing not implemented in this simple test")
		fmt.Println("Malicious wallets are stored internally in the AssociationRiskAnalyzer")

	default:
		fmt.Printf("Unknown malicious subcommand: %s\n", args[0])
	}
}

// Test association risk analysis
func testRiskAnalysis(ara *fraud.AssociationRiskAnalyzer, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: missing risk subcommand")
		return
	}

	switch args[0] {
	case "analyze":
		if len(args) < 2 {
			fmt.Println("Error: analyze requires wallet address argument")
			return
		}
		
		walletAddr := args[1]
		riskScore, paths := ara.ProcessAssociationAnalysis(walletAddr, time.Now())
		
		fmt.Printf("Association risk analysis for wallet %s:\n", walletAddr)
		fmt.Printf("Risk Score: %.2f\n", riskScore)
		
		if len(paths) > 0 {
			fmt.Printf("Found %d paths to malicious wallets:\n", len(paths))
			for i, path := range paths {
				fmt.Printf("Path %d: %d hops, weight %.2f\n", i+1, path.HopCount, path.TotalWeight)
				
				if len(path.Transactions) > 0 {
					fmt.Printf("  Starting from: %s\n", path.Transactions[0].FromAddress)
					if len(path.Transactions) > 1 {
						fmt.Printf("  Ending at: %s\n", path.Transactions[len(path.Transactions)-1].ToAddress)
					}
				}
			}
		} else {
			fmt.Println("No suspicious paths found.")
		}

	case "stats":
		mean, stdDev := ara.GetAssociationRiskStats()
		fmt.Println("Association risk statistics across the network:")
		fmt.Printf("Mean risk score: %.2f\n", mean)
		fmt.Printf("Standard deviation: %.2f\n", stdDev)

	default:
		fmt.Printf("Unknown risk subcommand: %s\n", args[0])
	}
}

// Test blacklisting functionality
func testBlacklisting(bm *fraud.BlacklistManager, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: missing blacklist subcommand")
		return
	}

	switch args[0] {
	case "penalize":
		if len(args) < 2 {
			fmt.Println("Error: penalize requires wallet address argument")
			return
		}
		
		walletAddr := args[1]
		gamma := 1.0 // Default gamma value
		
		if len(args) > 2 {
			gammaVal, err := strconv.ParseFloat(args[2], 64)
			if err == nil {
				gamma = gammaVal
			}
		}
		
		penalty, entry := bm.ProcessAssociationPenalty(walletAddr, time.Now(), gamma)
		
		fmt.Printf("Applied association penalty to wallet %s:\n", walletAddr)
		fmt.Printf("Penalty magnitude: %.2f\n", penalty)
		
		if entry != nil {
			fmt.Printf("Blacklist entry created:\n")
			fmt.Printf("  Status: %s\n", entry.Status)
			fmt.Printf("  Association Score: %.2f\n", entry.AssociationScore)
			fmt.Printf("  Associated with %d malicious wallets\n", len(entry.AssociatedWith))
		} else {
			fmt.Println("No blacklist entry was created (possibly no penalty was needed)")
		}

	case "check":
		if len(args) < 2 {
			fmt.Println("Error: check requires wallet address argument")
			return
		}
		
		walletAddr := args[1]
		isBlacklisted := bm.IsPartiallyBlacklisted(walletAddr)
		entry := bm.GetBlacklistEntry(walletAddr)
		
		if isBlacklisted {
			fmt.Printf("Wallet %s is partially blacklisted\n", walletAddr)
			if entry != nil {
				fmt.Printf("Blacklist details:\n")
				fmt.Printf("  Status: %s\n", entry.Status)
				fmt.Printf("  Reason: %s\n", entry.Reason)
				fmt.Printf("  Credit Penalty: %.2f\n", entry.CreditPenalty)
				fmt.Printf("  Blacklisted since: %s\n", entry.BlacklistTimestamp.Format(time.RFC3339))
			}
		} else if entry != nil {
			fmt.Printf("Wallet %s has blacklist entry with status: %s\n", walletAddr, entry.Status)
		} else {
			fmt.Printf("Wallet %s is not blacklisted\n", walletAddr)
		}
		
		factor := bm.CalculateEffectiveBlacklistFactor(walletAddr)
		fmt.Printf("Effective blacklist factor: %.2f\n", factor)

	case "restrictions":
		if len(args) < 2 {
			fmt.Println("Error: restrictions requires wallet address argument")
			return
		}
		
		walletAddr := args[1]
		restrictions := bm.GetVoteRestrictions(walletAddr)
		
		if len(restrictions) > 0 {
			fmt.Printf("Wallet %s has voting restrictions:\n", walletAddr)
			fmt.Printf("Cannot vote on: %s\n", strings.Join(restrictions, ", "))
		} else {
			fmt.Printf("Wallet %s has no voting restrictions\n", walletAddr)
		}

	case "update":
		if len(args) < 3 {
			fmt.Println("Error: update requires wallet address and status arguments")
			fmt.Println("Valid statuses: partial, full, rehab")
			return
		}
		
		walletAddr := args[1]
		statusStr := args[2]
		
		var status fraud.BlacklistStatus
		switch statusStr {
		case "partial":
			status = fraud.PartialBlacklist
		case "full":
			status = fraud.FullBlacklist
		case "rehab":
			status = fraud.Rehabilitated
		default:
			fmt.Printf("Unknown status: %s (use partial, full, or rehab)\n", statusStr)
			return
		}
		
		success := bm.UpdateBlacklistStatus(walletAddr, status)
		
		if success {
			fmt.Printf("Updated blacklist status for wallet %s to '%s'\n", walletAddr, status)
		} else {
			fmt.Printf("Failed to update blacklist status (wallet %s may not be blacklisted)\n", walletAddr)
		}

	default:
		fmt.Printf("Unknown blacklist subcommand: %s\n", args[0])
	}
}*/

//###############################################################################################################################################
//TEST SMART CONTRACTS

package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/IWALINK/cosine/pkg/smartcontracts/bindings"
	"github.com/spf13/cobra"
)

// Global variables
var (
	client                 *ethclient.Client
	privateKey             *ecdsa.PrivateKey
	auth                   *bind.TransactOpts
	cosineTokenAddress     common.Address
	walletLinkingAddress   common.Address
	creditVerifyAddress    common.Address
	bridgeAddress          common.Address
	rpcURL                 string
	privateKeyHex          string
	deployMode             bool
	testMode               bool
	bridgeChainID          int64
	addressFile            string
)

// Contract addresses structure
type ContractAddresses struct {
	CosineToken        string `json:"CosineToken"`
	WalletLinking      string `json:"WalletLinking"`
	CreditVerification string `json:"CreditVerification"`
	Bridge             string `json:"Bridge"`
}

func initEthClient() error {
	var err error
	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	privateKey, err = crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("invalid private key: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transaction signer: %v", err)
	}

	// Set higher gas limit for complex transactions
	auth.GasLimit = 8000000

	return nil
}

func loadContractAddresses() error {
	// If explicit addresses were provided, use those
	if cosineTokenAddress != (common.Address{}) &&
       walletLinkingAddress != (common.Address{}) &&
       creditVerifyAddress != (common.Address{}) &&
       bridgeAddress != (common.Address{}) {
         return nil
    }


	// Otherwise, load from file
	data, err := ioutil.ReadFile(addressFile)
	if err != nil {
		return fmt.Errorf("failed to read contract addresses file: %v", err)
	}

	var addresses ContractAddresses
	if err := json.Unmarshal(data, &addresses); err != nil {
		return fmt.Errorf("failed to parse contract addresses: %v", err)
	}

	cosineTokenAddress = common.HexToAddress(addresses.CosineToken)
	walletLinkingAddress = common.HexToAddress(addresses.WalletLinking)
	creditVerifyAddress = common.HexToAddress(addresses.CreditVerification)
	bridgeAddress = common.HexToAddress(addresses.Bridge)

	fmt.Printf("Loaded contract addresses:\n")
	fmt.Printf("CosineToken: %s\n", cosineTokenAddress.Hex())
	fmt.Printf("WalletLinking: %s\n", walletLinkingAddress.Hex())
	fmt.Printf("CreditVerification: %s\n", creditVerifyAddress.Hex())
	fmt.Printf("Bridge: %s\n", bridgeAddress.Hex())

	return nil
}

func runTests() error {
	fmt.Println("Running COSINE contract tests...")

	// Load contracts from addresses
	cosineToken, err := bindings.NewCosineToken(cosineTokenAddress, client)
	if err != nil {
		return fmt.Errorf("failed to load CosineToken: %v", err)
	}

	walletLinking, err := bindings.NewWalletLinking(walletLinkingAddress, client)
	if err != nil {
		return fmt.Errorf("failed to load WalletLinking: %v", err)
	}

	creditVerification, err := bindings.NewCreditVerification(creditVerifyAddress, client)
	if err != nil {
		return fmt.Errorf("failed to load CreditVerification: %v", err)
	}

	bridge, err := bindings.NewBridge(bridgeAddress, client)
	if err != nil {
		return fmt.Errorf("failed to load Bridge: %v", err)
	}

	// Test 1: Check token name and symbol
	fmt.Println("\n--- Test 1: Basic Token Information ---")
	name, err := cosineToken.Name(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get token name: %v", err)
	}
	symbol, err := cosineToken.Symbol(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get token symbol: %v", err)
	}
	fmt.Printf("CosineToken: Name = %s, Symbol = %s\n", name, symbol)

	// Test 2: Check total supply
	totalSupply, err := cosineToken.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get total supply: %v", err)
	}
	fmt.Printf("CosineToken: Total Supply = %s\n", totalSupply.String())

	// Test 3: Setup distribution addresses
	fmt.Println("\n--- Test 3: Token Distribution ---")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Generate some test addresses
	testAddresses := make([]common.Address, 6)
	for i := 0; i < 6; i++ {
		testKey, _ := crypto.GenerateKey()
		testAddresses[i] = crypto.PubkeyToAddress(*testKey.Public().(*ecdsa.PublicKey))
	}

	// Check if distributions are already initialized to avoid errors
	distributor, err := cosineToken.DISTRIBUTORROLE(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get DISTRIBUTOR_ROLE: %v", err)
	}
	
	hasRole, err := cosineToken.HasRole(&bind.CallOpts{}, distributor, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to check role: %v", err)
	}
	
	if !hasRole {
		// Grant DISTRIBUTOR_ROLE to our address
		admin, err := cosineToken.DEFAULTADMINROLE(&bind.CallOpts{})
		if err != nil {
			return fmt.Errorf("failed to get DEFAULT_ADMIN_ROLE: %v", err)
		}
		
		hasAdminRole, err := cosineToken.HasRole(&bind.CallOpts{}, admin, fromAddress)
		if err != nil {
			return fmt.Errorf("failed to check admin role: %v", err)
		}
		
		if hasAdminRole {
			tx, err := cosineToken.GrantRole(auth, distributor, fromAddress)
			if err != nil {
				return fmt.Errorf("failed to grant DISTRIBUTOR_ROLE: %v", err)
			}
			receipt, err := bind.WaitMined(context.Background(), client, tx)
			if err != nil {
				return fmt.Errorf("failed to wait for role grant: %v", err)
			}
			fmt.Printf("Granted DISTRIBUTOR_ROLE to self. Gas used: %d\n", receipt.GasUsed)
		} else {
			return fmt.Errorf("account does not have admin role to grant DISTRIBUTOR_ROLE")
		}
	}

	// Initialize distributions
	fmt.Println("Initializing token distributions...")
	
	// Check if distributions are already initialized
	initialized, err := cosineToken.DistributionsInitialized(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to check if distributions are initialized: %v", err)
	}
	
	if !initialized {
		tx, err := cosineToken.InitializeDistributions(auth, 
			testAddresses[0], // foundation
			testAddresses[1], // advisors
			testAddresses[2], // developers
			testAddresses[3], // privateSale
			testAddresses[4], // publicSale
			testAddresses[5], // networkRewards
		)
		if err != nil {
			return fmt.Errorf("failed to initialize distributions: %v", err)
		}
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return fmt.Errorf("failed to wait for distribution initialization: %v", err)
		}
		fmt.Printf("Distributions initialized! Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("Distributions already initialized, skipping...")
	}

	// Test 4: Get public sale balance (should be immediately available)
	fmt.Println("\n--- Test 4: Public Sale Balance Check ---")
	publicSaleBalance, err := cosineToken.BalanceOf(&bind.CallOpts{}, testAddresses[4])
	if err != nil {
		return fmt.Errorf("failed to get public sale balance: %v", err)
	}
	
	// Convert to readable format
	totalSupplyFloat := new(big.Float).SetInt(totalSupply)
	publicSaleBalanceFloat := new(big.Float).SetInt(publicSaleBalance)
	percentage := new(big.Float).Quo(publicSaleBalanceFloat, totalSupplyFloat)
	percentage.Mul(percentage, big.NewFloat(100))
	percentageFloat, _ := percentage.Float64()
	
	fmt.Printf("Public Sale Address Balance: %s (%.2f%% of total supply)\n", publicSaleBalance.String(), percentageFloat)
	
	// Should be approximately 20% of total supply
	if percentageFloat < 19.0 || percentageFloat > 21.0 {
		fmt.Printf("Warning: Public sale balance percentage (%.2f%%) is not close to the expected 20%%\n", percentageFloat)
	}

	// Test 5: Test wallet linking
	fmt.Println("\n--- Test 5: Wallet Linking ---")
	
	// Generate a challenge
	tx, err := walletLinking.GenerateChallenge(auth, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to generate challenge: %v", err)
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for challenge generation: %v", err)
	}
	
	// Get the challenge hash from logs (simplified approach)
	var challengeHash [32]byte
	fmt.Printf("Challenge generated. Transaction: %s\n", tx.Hash().Hex())
	
	// Query the challenge directly from the contract
	challenge, err := walletLinking.Challenges(&bind.CallOpts{}, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to get challenge: %v", err)
	}
	challengeHash = challenge.ChallengeHash
	fmt.Printf("Challenge hash: %x\n", challengeHash)
	
	// Sign the challenge
	// Create message hash that follows Ethereum specific signing format
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n32%s", string(challengeHash[:]))
	msgHash := crypto.Keccak256Hash([]byte(msg))
	signature, err := crypto.Sign(msgHash.Bytes(), privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign challenge: %v", err)
	}
	fmt.Printf("Signature created: %x\n", signature)
	
	// Create a mock L2 wallet address (hash of our address as an example)
	l2WalletBytes := crypto.Keccak256Hash(fromAddress.Bytes())
	var l2Wallet [32]byte
	copy(l2Wallet[:], l2WalletBytes[:])
	fmt.Printf("L2 wallet identifier: %x\n", l2Wallet)
	
	// Link wallet
	chainID := big.NewInt(31337) // Local chainID
	tx, err = walletLinking.LinkWallet(auth, l2Wallet, chainID, signature)
	if err != nil {
		return fmt.Errorf("failed to link wallet: %v", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for wallet linking: %v", err)
	}
	fmt.Printf("Wallet linked! Gas used: %d\n", receipt.GasUsed)
	
	// Verify linking
	isLinked, err := walletLinking.VerifyLinking(&bind.CallOpts{}, fromAddress, l2Wallet, chainID)
	if err != nil {
		return fmt.Errorf("failed to verify linking: %v", err)
	}
	fmt.Printf("Wallet linking verified: %v\n", isLinked)

	// Test 6: Test credit verification
	fmt.Println("\n--- Test 6: Credit Verification ---")
	
	// Grant BRIDGE_ROLE to bridge address
	bridgeRole, err := cosineToken.BRIDGEROLE(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get BRIDGE_ROLE: %v", err)
	}
	
	// Check if role already granted
	hasRole, err = cosineToken.HasRole(&bind.CallOpts{}, bridgeRole, bridgeAddress)
	if err != nil {
		return fmt.Errorf("failed to check role: %v", err)
	}
	
	if !hasRole {
		tx, err = cosineToken.GrantRole(auth, bridgeRole, bridgeAddress)
		if err != nil {
			return fmt.Errorf("failed to grant BRIDGE_ROLE: %v", err)
		}
		receipt, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return fmt.Errorf("failed to wait for role granting: %v", err)
		}
		fmt.Printf("BRIDGE_ROLE granted to bridge address. Gas used: %d\n", receipt.GasUsed)
	} else {
		fmt.Println("BRIDGE_ROLE already granted to bridge address")
	}
	
	// Create a test threshold vector
	thresholdVector := bindings.CreditVerificationVector2D{
		X: big.NewInt(500000000000000000), // 0.5 * 10^18
		Y: big.NewInt(1000000000000000000), // 1.0 * 10^18
	}
	
	// Approve tokens for verification fee
	tx, err = cosineToken.Approve(auth, creditVerifyAddress, big.NewInt(1000000000000000000)) // 1 COSINE
	if err != nil {
		return fmt.Errorf("failed to approve tokens: %v", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for token approval: %v", err)
	}
	fmt.Printf("Tokens approved for credit verification. Gas used: %d\n", receipt.GasUsed)
	
	// Verify credit score
	fmt.Println("Verifying credit score...")
	tx, err = creditVerification.VerifyCreditScore(
		auth,
		l2Wallet,
		thresholdVector,
		big.NewInt(900000000000000000), // 0.9 * 10^18 threshold similarity
	)
	
	if err != nil {
		return fmt.Errorf("failed to verify credit score: %v", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for credit verification: %v", err)
	}
	fmt.Printf("Credit score verified! Gas used: %d\n", receipt.GasUsed)
	
	// Check if feedback is required
	requiresFeedback, err := creditVerification.CheckFeedbackRequirement(&bind.CallOpts{}, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to check feedback requirement: %v", err)
	}
	fmt.Printf("Feedback required: %v\n", requiresFeedback)
	
	// Provide feedback
	if requiresFeedback {
		tx, err = creditVerification.ProvideTransactionFeedback(auth, l2Wallet, true) // trusted
		if err != nil {
			return fmt.Errorf("failed to provide feedback: %v", err)
		}
		receipt, err = bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			return fmt.Errorf("failed to wait for feedback: %v", err)
		}
		fmt.Printf("Feedback provided! Gas used: %d\n", receipt.GasUsed)
	}

	// Test 7: Test bridge
	fmt.Println("\n--- Test 7: Bridge Operations ---")
	
	// Approve tokens for bridging
	bridgeAmount, ok := new(big.Int).SetString("10000000000000000000", 10)
	if !ok {
		log.Fatalf("Invalid bridge amount")
	}
	tx, err = cosineToken.Approve(auth, bridgeAddress, bridgeAmount)
	if err != nil {
		return fmt.Errorf("failed to approve tokens for bridge: %v", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for token approval: %v", err)
	}
	fmt.Printf("Tokens approved for bridge. Gas used: %d\n", receipt.GasUsed)
	
	// Initiate outgoing bridge
	destinationChainID := big.NewInt(42) // Example destination chain
	tx, err = bridge.InitiateOutgoingBridge(auth, l2Wallet, destinationChainID, bridgeAmount)
	if err != nil {
		return fmt.Errorf("failed to initiate bridge: %v", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for bridge initiation: %v", err)
	}
	fmt.Printf("Bridge operation initiated! Gas used: %d\n", receipt.GasUsed)
	
	// Get the operation ID (should be 1 for the first operation)
	operationId := big.NewInt(1)
	
	// Check bridge operation status
    status, err := bridge.CheckOperationStatus(&bind.CallOpts{}, operationId)
	if err != nil {
		return fmt.Errorf("failed to check bridge operation status: %v", err)
	}
	fmt.Printf("Bridge operation: Exists=%v, Initiator=%s, Amount=%s, Confirmations=%d, Completed=%v\n", 
		status.Exists, status.Initiator.Hex(), status.Amount.String(), status.Confirmations, status.IsCompleted)
	
	// Add validator role to confirm operations
	validatorRole, err := bridge.VALIDATORROLE(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to get VALIDATOR_ROLE: %v", err)
	}
	
	hasRole, err = bridge.HasRole(&bind.CallOpts{}, validatorRole, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to check validator role: %v", err)
	}
	
	if !hasRole {
		admin, err := bridge.DEFAULTADMINROLE(&bind.CallOpts{})
		if err != nil {
			return fmt.Errorf("failed to get DEFAULT_ADMIN_ROLE: %v", err)
		}
		
		hasAdminRole, err := bridge.HasRole(&bind.CallOpts{}, admin, fromAddress)
		if err != nil {
			return fmt.Errorf("failed to check admin role: %v", err)
		}
		
		if hasAdminRole {
			tx, err = bridge.GrantRole(auth, validatorRole, fromAddress)
			if err != nil {
				return fmt.Errorf("failed to grant VALIDATOR_ROLE: %v", err)
			}
			receipt, err = bind.WaitMined(context.Background(), client, tx)
			if err != nil {
				return fmt.Errorf("failed to wait for role grant: %v", err)
			}
			fmt.Printf("Granted VALIDATOR_ROLE to self. Gas used: %d\n", receipt.GasUsed)
		} else {
			return fmt.Errorf("account does not have admin role to grant VALIDATOR_ROLE")
		}
	}
	
	// Confirm operation (as validator)
	tx, err = bridge.ConfirmOperation(auth, operationId)
	if err != nil {
		return fmt.Errorf("failed to confirm operation: %v", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for operation confirmation: %v", err)
	}
	fmt.Printf("Operation confirmed! Gas used: %d\n", receipt.GasUsed)
	
	// Check status again after confirmation
	status, err = bridge.CheckOperationStatus(&bind.CallOpts{}, operationId)
	if err != nil {
		return fmt.Errorf("failed to check bridge operation status: %v", err)
	}
	fmt.Printf("Bridge operation: Exists=%v, Initiator=%s, Amount=%s, Confirmations=%d, Completed=%v\n",
		status.Exists, status.Initiator.Hex(), status.Amount.String(), status.Confirmations, status.IsCompleted)

	fmt.Println("\nAll tests completed successfully!")
	return nil
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cosine-cli",
		Short: "COSINE CLI for interacting with COSINE smart contracts",
		Run: func(cmd *cobra.Command, args []string) {
			if err := initEthClient(); err != nil {
				log.Fatalf("Failed to initialize Ethereum client: %v", err)
			}

			if testMode {
				if err := loadContractAddresses(); err != nil {
					log.Fatalf("Failed to load contract addresses: %v", err)
				}
				
				if err := runTests(); err != nil {
					log.Fatalf("Tests failed: %v", err)
				}
			}
		},
	}

	rootCmd.Flags().StringVar(&rpcURL, "rpc", "http://localhost:8545", "Ethereum RPC URL")
	//ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 to test for an account without permissions
	rootCmd.Flags().StringVar(&privateKeyHex, "key", "4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d", "Private key in hex format (without '0x' prefix)")
	rootCmd.Flags().BoolVar(&testMode, "test", false, "Run tests on deployed contracts")
	rootCmd.Flags().Int64Var(&bridgeChainID, "chain-id", 31337, "Chain ID for the Bridge contract")
	rootCmd.Flags().StringVar(&addressFile, "address-file", "contract_addresses.json", "JSON file containing contract addresses")
	

	var linkingAddrStr, creditAddrStr, tokenAddrStr, bridgeAddrStr string

	rootCmd.Flags().StringVar(&tokenAddrStr, "token-addr", "", "CosineToken contract address")
    rootCmd.Flags().StringVar(&linkingAddrStr, "linking-addr", "", "WalletLinking contract address")
    rootCmd.Flags().StringVar(&creditAddrStr, "credit-addr", "", "CreditVerification contract address")
    rootCmd.Flags().StringVar(&bridgeAddrStr, "bridge-addr", "", "Bridge contract address")

	if tokenAddrStr != "" {
       cosineTokenAddress = common.HexToAddress(tokenAddrStr)
    }

	if linkingAddrStr != "" {
    walletLinkingAddress = common.HexToAddress(linkingAddrStr)
	}
	if creditAddrStr != "" {
		creditVerifyAddress = common.HexToAddress(creditAddrStr)
	}
	if bridgeAddrStr != "" {
		bridgeAddress = common.HexToAddress(bridgeAddrStr)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}