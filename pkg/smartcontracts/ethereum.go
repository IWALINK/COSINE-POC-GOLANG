package smartcontracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
	"github.com/IWALINK/cosine/pkg/smartcontracts/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ContractConfig represents the configuration for Ethereum contracts
type ContractConfig struct {
	RPCEndpoint          string `json:"rpc_endpoint"`
	ChainID              int64  `json:"chain_id"`
	WalletLinkingAddress string `json:"wallet_linking_address"`
	CosineTokenAddress   string `json:"cosine_token_address"`
	CreditVerifyAddress  string `json:"credit_verify_address"`
	BridgeAddress        string `json:"bridge_address"`
	OperatorPrivateKey   string `json:"operator_private_key,omitempty"`
	RetryAttempts        int    `json:"retry_attempts"`
	RetryDelay           int    `json:"retry_delay_ms"`
}

// EthereumClient represents a client for interacting with Ethereum contracts
type EthereumClient struct {
	config          *ContractConfig
	client          *ethclient.Client
	walletLinking   *bindings.WalletLinking
	cosineToken     *bindings.CosineToken
	creditVerify    *bindings.CreditVerification
	bridge          *bindings.Bridge
	auth            *bind.TransactOpts
	privateKey      *ecdsa.PrivateKey
	logger          *utils.Logger
	metrics         *utils.MetricsManager
	operatorAddress common.Address
	mu              sync.RWMutex
}

// New creates a new Ethereum client
func New(config *ContractConfig, logger *utils.Logger, metrics *utils.MetricsManager) (*EthereumClient, error) {
	logger = logger.WithComponent("EthereumClient")
	logger.Info("Initializing Ethereum client", "rpc_endpoint", config.RPCEndpoint)

	client, err := ethclient.Dial(config.RPCEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	// Create a new Ethereum client
	ethClient := &EthereumClient{
		config:  config,
		client:  client,
		logger:  logger,
		metrics: metrics,
	}

	// Load the private key if provided
	if config.OperatorPrivateKey != "" {
		if err := ethClient.loadPrivateKey(); err != nil {
			return nil, err
		}
	}

	// Load the contracts
	if err := ethClient.loadContracts(); err != nil {
		return nil, err
	}

	logger.Info("Ethereum client initialized", 
		"chain_id", config.ChainID,
		"wallet_linking_address", config.WalletLinkingAddress)

	return ethClient, nil
}

// loadPrivateKey loads the private key for transaction signing
func (ec *EthereumClient) loadPrivateKey() error {
	var err error
	// Remove 0x prefix if present
	privKeyStr := strings.TrimPrefix(ec.config.OperatorPrivateKey, "0x")
	
	ec.privateKey, err = crypto.HexToECDSA(privKeyStr)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	// Derive the public key and address
	publicKey := ec.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}
	ec.operatorAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	// Create auth for transactions
	chainID := big.NewInt(ec.config.ChainID)
	ec.auth, err = bind.NewKeyedTransactorWithChainID(ec.privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create authorized transactor: %w", err)
	}

	ec.logger.Info("Loaded operator key", "address", ec.operatorAddress.Hex())
	return nil
}

// loadContracts loads the contract bindings
func (ec *EthereumClient) loadContracts() error {
	var err error

	// Load WalletLinking contract
	if ec.config.WalletLinkingAddress != "" {
		walletLinkingAddress := common.HexToAddress(ec.config.WalletLinkingAddress)
		ec.walletLinking, err = bindings.NewWalletLinking(walletLinkingAddress, ec.client)
		if err != nil {
			return fmt.Errorf("failed to load WalletLinking contract: %w", err)
		}
	}

	// Load CosineToken contract
	if ec.config.CosineTokenAddress != "" {
		cosineTokenAddress := common.HexToAddress(ec.config.CosineTokenAddress)
		ec.cosineToken, err = bindings.NewCosineToken(cosineTokenAddress, ec.client)
		if err != nil {
			return fmt.Errorf("failed to load CosineToken contract: %w", err)
		}
	}

	// Load CreditVerification contract
	if ec.config.CreditVerifyAddress != "" {
		creditVerifyAddress := common.HexToAddress(ec.config.CreditVerifyAddress)
		ec.creditVerify, err = bindings.NewCreditVerification(creditVerifyAddress, ec.client)
		if err != nil {
			return fmt.Errorf("failed to load CreditVerification contract: %w", err)
		}
	}

	// Load Bridge contract
	if ec.config.BridgeAddress != "" {
		bridgeAddress := common.HexToAddress(ec.config.BridgeAddress)
		ec.bridge, err = bindings.NewBridge(bridgeAddress, ec.client)
		if err != nil {
			return fmt.Errorf("failed to load Bridge contract: %w", err)
		}
	}

	return nil
}

// GenerateChallenge generates a challenge for wallet linking
func (ec *EthereumClient) GenerateChallenge(ctx context.Context, ethAddress common.Address) ([32]byte, error) {
	if ec.walletLinking == nil {
		return [32]byte{}, fmt.Errorf("wallet linking contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Generating challenge", "eth_address", ethAddress.Hex())

	// Set gas price and limit
	ec.mu.Lock()
	gasPrice, err := ec.client.SuggestGasPrice(ctx)
	if err != nil {
		ec.mu.Unlock()
		return [32]byte{}, fmt.Errorf("failed to suggest gas price: %w", err)
	}
	
	ec.auth.GasPrice = gasPrice
	ec.auth.GasLimit = 300000 // Adjust as needed
	ec.mu.Unlock()

	// Generate the challenge with retry mechanism
	var (
		tx         *types.Transaction
		challengeHash [32]byte
		attempts   = 0
		retryDelay = time.Duration(ec.config.RetryDelay) * time.Millisecond
	)

	for attempts < ec.config.RetryAttempts {
		tx, err = ec.walletLinking.GenerateChallenge(ec.auth, ethAddress)
		if err == nil {
			break
		}
		
		ec.logger.Warn("Failed to generate challenge, retrying", 
			"error", err, 
			"attempt", attempts+1, 
			"max_attempts", ec.config.RetryAttempts)
		
		attempts++
		if attempts >= ec.config.RetryAttempts {
			return [32]byte{}, fmt.Errorf("failed to generate challenge after %d attempts: %w", ec.config.RetryAttempts, err)
		}
		
		time.Sleep(retryDelay)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, ec.client, tx)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to wait for challenge generation: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return [32]byte{}, fmt.Errorf("challenge generation transaction failed")
	}

	// Get the challenge from the contract
	challenge, err := ec.walletLinking.Challenges(&bind.CallOpts{}, ethAddress)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to get challenge: %w", err)
	}

	challengeHash = challenge.ChallengeHash
	
	ec.logger.Info("Challenge generated", 
		"eth_address", ethAddress.Hex(), 
		"challenge_hash", fmt.Sprintf("%x", challengeHash), 
		"tx_hash", tx.Hash().Hex(),
		"duration", time.Since(startTime))

	// Record metrics
	if ec.metrics != nil {
		ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "generate_challenge")
	}

	return challengeHash, nil
}

// LinkWallet links an Ethereum address to an L2 wallet
func (ec *EthereumClient) LinkWallet(ctx context.Context, l2Wallet [32]byte, signature []byte) error {
	if ec.walletLinking == nil {
		return fmt.Errorf("wallet linking contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Linking wallet", "l2_wallet", fmt.Sprintf("%x", l2Wallet))

	// Set gas price and limit
	ec.mu.Lock()
	gasPrice, err := ec.client.SuggestGasPrice(ctx)
	if err != nil {
		ec.mu.Unlock()
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}
	
	ec.auth.GasPrice = gasPrice
	ec.auth.GasLimit = 300000 // Adjust as needed
	ec.mu.Unlock()

	// Link the wallet with retry mechanism
	var (
		tx         *types.Transaction
		attempts   = 0
		retryDelay = time.Duration(ec.config.RetryDelay) * time.Millisecond
		chainID    = big.NewInt(ec.config.ChainID)
	)

	for attempts < ec.config.RetryAttempts {
		tx, err = ec.walletLinking.LinkWallet(ec.auth, l2Wallet, chainID, signature)
		if err == nil {
			break
		}
		
		ec.logger.Warn("Failed to link wallet, retrying", 
			"error", err, 
			"attempt", attempts+1, 
			"max_attempts", ec.config.RetryAttempts)
		
		attempts++
		if attempts >= ec.config.RetryAttempts {
			return fmt.Errorf("failed to link wallet after %d attempts: %w", ec.config.RetryAttempts, err)
		}
		
		time.Sleep(retryDelay)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, ec.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for link wallet transaction: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("link wallet transaction failed")
	}

	ec.logger.Info("Wallet linked", 
		"l2_wallet", fmt.Sprintf("%x", l2Wallet), 
		"tx_hash", tx.Hash().Hex(),
		"duration", time.Since(startTime))

	// Record metrics
	if ec.metrics != nil {
		ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "link_wallet")
	}

	return nil
}

// VerifyLinking checks if an Ethereum address is linked to an L2 wallet
func (ec *EthereumClient) VerifyLinking(ctx context.Context, ethAddress common.Address, l2Wallet [32]byte) (bool, error) {
	if ec.walletLinking == nil {
		return false, fmt.Errorf("wallet linking contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Verifying linking", 
		"eth_address", ethAddress.Hex(), 
		"l2_wallet", fmt.Sprintf("%x", l2Wallet))

	chainID := big.NewInt(ec.config.ChainID)
	isLinked, err := ec.walletLinking.VerifyLinking(&bind.CallOpts{Context: ctx}, ethAddress, l2Wallet, chainID)
	if err != nil {
		return false, fmt.Errorf("failed to verify linking: %w", err)
	}

	ec.logger.Debug("Verified linking", 
		"eth_address", ethAddress.Hex(), 
		"l2_wallet", fmt.Sprintf("%x", l2Wallet), 
		"is_linked", isLinked,
		"duration", time.Since(startTime))

	// Record metrics
	if ec.metrics != nil {
		ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "verify_linking")
	}

	return isLinked, nil
}

// GetLinkedL2Wallet gets the L2 wallet linked to an Ethereum address
func (ec *EthereumClient) GetLinkedL2Wallet(ctx context.Context, ethAddress common.Address) ([32]byte, error) {
	if ec.walletLinking == nil {
		return [32]byte{}, fmt.Errorf("wallet linking contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Getting linked L2 wallet", "eth_address", ethAddress.Hex())

	l2Wallet, err := ec.walletLinking.GetLinkedL2Wallet(&bind.CallOpts{Context: ctx}, ethAddress)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to get linked L2 wallet: %w", err)
	}

	ec.logger.Debug("Got linked L2 wallet", 
		"eth_address", ethAddress.Hex(), 
		"l2_wallet", fmt.Sprintf("%x", l2Wallet),
		"duration", time.Since(startTime))

	// Record metrics
	if ec.metrics != nil {
		ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "get_linked_l2_wallet")
	}

	return l2Wallet, nil
}

// GetLinkedL1Address gets the Ethereum address linked to an L2 wallet
func (ec *EthereumClient) GetLinkedL1Address(ctx context.Context, l2Wallet [32]byte) (common.Address, error) {
	if ec.walletLinking == nil {
		return common.Address{}, fmt.Errorf("wallet linking contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Getting linked L1 address", "l2_wallet", fmt.Sprintf("%x", l2Wallet))

	chainID := big.NewInt(ec.config.ChainID)
	ethAddress, err := ec.walletLinking.GetLinkedL1Address(&bind.CallOpts{Context: ctx}, l2Wallet, chainID)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get linked L1 address: %w", err)
	}

	ec.logger.Debug("Got linked L1 address", 
		"l2_wallet", fmt.Sprintf("%x", l2Wallet), 
		"eth_address", ethAddress.Hex(),
		"duration", time.Since(startTime))

	// Record metrics
	if ec.metrics != nil {
		ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "get_linked_l1_address")
	}

	return ethAddress, nil
}

// VerifyCreditScore verifies the credit score of an L2 wallet
func (ec *EthereumClient) VerifyCreditScore(ctx context.Context, l2Wallet [32]byte, thresholdX, thresholdY *big.Int, thresholdSimilarity *big.Int) (int64, bool, error) {
	if ec.creditVerify == nil {
		return 0, false, fmt.Errorf("credit verification contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Verifying credit score", "l2_wallet", fmt.Sprintf("%x", l2Wallet))

	// Set gas price and limit
	ec.mu.Lock()
	gasPrice, err := ec.client.SuggestGasPrice(ctx)
	if err != nil {
		ec.mu.Unlock()
		return 0, false, fmt.Errorf("failed to suggest gas price: %w", err)
	}
	
	ec.auth.GasPrice = gasPrice
	ec.auth.GasLimit = 300000 // Adjust as needed
	ec.mu.Unlock()

	// Create threshold vector
	thresholdVector := bindings.CreditVerificationVector2D{
		X: thresholdX,
		Y: thresholdY,
	}

	// First, approve tokens for verification fee
	fee, err := ec.creditVerify.CalculateVerificationFee(&bind.CallOpts{Context: ctx}, l2Wallet)
	if err != nil {
		return 0, false, fmt.Errorf("failed to calculate verification fee: %w", err)
	}

	ec.logger.Debug("Approving tokens for verification", "fee", fee.String())
	
	approveTx, err := ec.cosineToken.Approve(ec.auth, common.HexToAddress(ec.config.CreditVerifyAddress), fee)
	if err != nil {
		return 0, false, fmt.Errorf("failed to approve tokens: %w", err)
	}

	_, err = bind.WaitMined(ctx, ec.client, approveTx)
	if err != nil {
		return 0, false, fmt.Errorf("failed to wait for approval transaction: %w", err)
	}

	// Verify credit score with retry mechanism
	var (
		tx         *types.Transaction
		attempts   = 0
		retryDelay = time.Duration(ec.config.RetryDelay) * time.Millisecond
	)

	for attempts < ec.config.RetryAttempts {
		tx, err = ec.creditVerify.VerifyCreditScore(ec.auth, l2Wallet, thresholdVector, thresholdSimilarity)
		if err == nil {
			break
		}
		
		ec.logger.Warn("Failed to verify credit score, retrying", 
			"error", err, 
			"attempt", attempts+1, 
			"max_attempts", ec.config.RetryAttempts)
		
		attempts++
		if attempts >= ec.config.RetryAttempts {
			return 0, false, fmt.Errorf("failed to verify credit score after %d attempts: %w", ec.config.RetryAttempts, err)
		}
		
		time.Sleep(retryDelay)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, ec.client, tx)
	if err != nil {
		return 0, false, fmt.Errorf("failed to wait for credit verification: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return 0, false, fmt.Errorf("credit verification transaction failed")
	}

	// Parse the transaction logs to get the result
	creditVerifyAbi, err := abi.JSON(strings.NewReader(bindings.CreditVerificationMetaData.ABI))
	if err != nil {
		return 0, false, fmt.Errorf("failed to parse credit verification ABI: %w", err)
	}

	for _, log := range receipt.Logs {
		event, err := creditVerifyAbi.EventByID(log.Topics[0])
		if err != nil {
			continue
		}
		
		if event.Name == "CreditVerified" {
			var result struct {
				Verifier       common.Address
				L2Wallet       [32]byte
				Fee            *big.Int
				PassedThreshold bool
			}
			
			err = creditVerifyAbi.UnpackIntoInterface(&result, "CreditVerified", log.Data)
			if err != nil {
				return 0, false, fmt.Errorf("failed to unpack CreditVerified event: %w", err)
			}
			
			// We don't get the actual similarity from the event, so we'll return -1 as a placeholder
			// In a real implementation, we would either modify the contract to return this value
			// or estimate it based on other available data
			similarity := int64(-1)
			
			ec.logger.Info("Credit score verified", 
				"l2_wallet", fmt.Sprintf("%x", l2Wallet), 
				"passed_threshold", result.PassedThreshold,
				"tx_hash", tx.Hash().Hex(),
				"duration", time.Since(startTime))
			
			// Record metrics
			if ec.metrics != nil {
				ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "verify_credit_score")
			}
			
			return similarity, result.PassedThreshold, nil
		}
	}

	return 0, false, fmt.Errorf("failed to find CreditVerified event in logs")
}

// ProvideTransactionFeedback provides feedback after a transaction
func (ec *EthereumClient) ProvideTransactionFeedback(ctx context.Context, l2Wallet [32]byte, isTrusted bool) error {
	if ec.creditVerify == nil {
		return fmt.Errorf("credit verification contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Providing transaction feedback", 
		"l2_wallet", fmt.Sprintf("%x", l2Wallet), 
		"is_trusted", isTrusted)

	// Set gas price and limit
	ec.mu.Lock()
	gasPrice, err := ec.client.SuggestGasPrice(ctx)
	if err != nil {
		ec.mu.Unlock()
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}
	
	ec.auth.GasPrice = gasPrice
	ec.auth.GasLimit = 300000 // Adjust as needed
	ec.mu.Unlock()

	// Provide feedback with retry mechanism
	var (
		tx         *types.Transaction
		attempts   = 0
		retryDelay = time.Duration(ec.config.RetryDelay) * time.Millisecond
	)

	for attempts < ec.config.RetryAttempts {
		tx, err = ec.creditVerify.ProvideTransactionFeedback(ec.auth, l2Wallet, isTrusted)
		if err == nil {
			break
		}
		
		ec.logger.Warn("Failed to provide feedback, retrying", 
			"error", err, 
			"attempt", attempts+1, 
			"max_attempts", ec.config.RetryAttempts)
		
		attempts++
		if attempts >= ec.config.RetryAttempts {
			return fmt.Errorf("failed to provide feedback after %d attempts: %w", ec.config.RetryAttempts, err)
		}
		
		time.Sleep(retryDelay)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, ec.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for feedback transaction: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("feedback transaction failed")
	}

	ec.logger.Info("Transaction feedback provided", 
		"l2_wallet", fmt.Sprintf("%x", l2Wallet), 
		"is_trusted", isTrusted,
		"tx_hash", tx.Hash().Hex(),
		"duration", time.Since(startTime))

	// Record metrics
	if ec.metrics != nil {
		ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "provide_feedback")
	}

	return nil
}

// CheckFeedbackRequirement checks if feedback is required
func (ec *EthereumClient) CheckFeedbackRequirement(ctx context.Context, address common.Address) (bool, error) {
	if ec.creditVerify == nil {
		return false, fmt.Errorf("credit verification contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Checking feedback requirement", "address", address.Hex())

	requiresFeedback, err := ec.creditVerify.CheckFeedbackRequirement(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return false, fmt.Errorf("failed to check feedback requirement: %w", err)
	}

	ec.logger.Debug("Checked feedback requirement", 
		"address", address.Hex(), 
		"requires_feedback", requiresFeedback,
		"duration", time.Since(startTime))

	// Record metrics
	if ec.metrics != nil {
		ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "check_feedback_requirement")
	}

	return requiresFeedback, nil
}

// InitiateOutgoingBridge initiates a bridge operation from Ethereum to COSINE L2
func (ec *EthereumClient) InitiateOutgoingBridge(ctx context.Context, l2Wallet [32]byte, destinationChainID *big.Int, amount *big.Int) (uint64, error) {
	if ec.bridge == nil {
		return 0, fmt.Errorf("bridge contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Initiating outgoing bridge operation", 
		"l2_wallet", fmt.Sprintf("%x", l2Wallet), 
		"destination_chain_id", destinationChainID.String(),
		"amount", amount.String())

	// Set gas price and limit
	ec.mu.Lock()
	gasPrice, err := ec.client.SuggestGasPrice(ctx)
	if err != nil {
		ec.mu.Unlock()
		return 0, fmt.Errorf("failed to suggest gas price: %w", err)
	}
	
	ec.auth.GasPrice = gasPrice
	ec.auth.GasLimit = 300000 // Adjust as needed
	ec.mu.Unlock()

	// First, approve tokens for bridge
	approveTx, err := ec.cosineToken.Approve(ec.auth, common.HexToAddress(ec.config.BridgeAddress), amount)
	if err != nil {
		return 0, fmt.Errorf("failed to approve tokens: %w", err)
	}

	_, err = bind.WaitMined(ctx, ec.client, approveTx)
	if err != nil {
		return 0, fmt.Errorf("failed to wait for approval transaction: %w", err)
	}

	// Initiate bridge operation with retry mechanism
	var (
		tx         *types.Transaction
		attempts   = 0
		retryDelay = time.Duration(ec.config.RetryDelay) * time.Millisecond
	)

	for attempts < ec.config.RetryAttempts {
		tx, err = ec.bridge.InitiateOutgoingBridge(ec.auth, l2Wallet, destinationChainID, amount)
		if err == nil {
			break
		}
		
		ec.logger.Warn("Failed to initiate bridge operation, retrying", 
			"error", err, 
			"attempt", attempts+1, 
			"max_attempts", ec.config.RetryAttempts)
		
		attempts++
		if attempts >= ec.config.RetryAttempts {
			return 0, fmt.Errorf("failed to initiate bridge operation after %d attempts: %w", ec.config.RetryAttempts, err)
		}
		
		time.Sleep(retryDelay)
	}

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, ec.client, tx)
	if err != nil {
		return 0, fmt.Errorf("failed to wait for bridge initiation: %w", err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return 0, fmt.Errorf("bridge initiation transaction failed")
	}

	// Parse the logs to get the operation ID
	bridgeAbi, err := abi.JSON(strings.NewReader(bindings.BridgeMetaData.ABI))
	if err != nil {
		return 0, fmt.Errorf("failed to parse bridge ABI: %w", err)
	}

	var operationID uint64
	for _, log := range receipt.Logs {
		event, err := bridgeAbi.EventByID(log.Topics[0])
		if err != nil {
			continue
		}
		
		if event.Name == "BridgeInitiated" {
			var result struct {
				OperationId       *big.Int
				Initiator         common.Address
				L2Wallet          [32]byte
				SourceChainId     *big.Int
				DestinationChainId *big.Int
				Amount            *big.Int
				Fee               *big.Int
				IsLock            bool
			}
			
			err = bridgeAbi.UnpackIntoInterface(&result, "BridgeInitiated", log.Data)
			if err != nil {
				return 0, fmt.Errorf("failed to unpack BridgeInitiated event: %w", err)
			}
			
			operationID = result.OperationId.Uint64()
			
			ec.logger.Info("Outgoing bridge operation initiated", 
				"operation_id", operationID,
				"l2_wallet", fmt.Sprintf("%x", l2Wallet), 
				"destination_chain_id", destinationChainID.String(),
				"amount", amount.String(),
				"tx_hash", tx.Hash().Hex(),
				"duration", time.Since(startTime))
			
			// Record metrics
			if ec.metrics != nil {
				ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "initiate_outgoing_bridge")
			}
			
			return operationID, nil
		}
	}

	return 0, fmt.Errorf("failed to find BridgeInitiated event in logs")
}

// CheckOperationStatus checks the status of a bridge operation
func (ec *EthereumClient) CheckOperationStatus(ctx context.Context, operationID *big.Int) (bool, common.Address, *big.Int, *big.Int, bool, error) {
	if ec.bridge == nil {
		return false, common.Address{}, nil, big.NewInt(0), false, fmt.Errorf("bridge contract not initialized")
	}

	startTime := time.Now()
	ec.logger.Debug("Checking bridge operation status", "operation_id", operationID.String())

	status, err := ec.bridge.CheckOperationStatus(&bind.CallOpts{Context: ctx}, operationID)
	if err != nil {
		return false, common.Address{}, nil, big.NewInt(0), false, fmt.Errorf("failed to check operation status: %w", err)
	}

	ec.logger.Debug("Checked bridge operation status", 
		"operation_id", operationID.String(), 
		"exists", status.Exists,
		"initiator", status.Initiator.Hex(),
		"amount", status.Amount.String(),
		"confirmations", status.Confirmations,
		"is_completed", status.IsCompleted,
		"duration", time.Since(startTime))

	// Record metrics
	if ec.metrics != nil {
		ec.metrics.ObserveHistogram("eth_operation_duration_seconds", time.Since(startTime).Seconds(), "operation", "check_operation_status")
	}

	return status.Exists, status.Initiator, status.Amount, status.Confirmations, status.IsCompleted, nil
}

// Close closes the Ethereum client
func (ec *EthereumClient) Close() {
	if ec.client != nil {
		ec.client.Close()
	}
	ec.logger.Info("Ethereum client closed")
}