package walletlink

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
	"github.com/IWALINK/cosine/pkg/smartcontracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Common errors
var (
	ErrInvalidEthereumAddress      = errors.New("invalid Ethereum address")
	ErrInvalidSignature    = errors.New("invalid signature")
	ErrChallengeExpired    = errors.New("challenge has expired")
	ErrChallengeNotFound   = errors.New("no challenge found for address")
	ErrAddressAlreadyLinked = errors.New("address is already linked to an L2 wallet")
	ErrL2WalletAlreadyLinked = errors.New("L2 wallet is already linked to an address on this chain")
)

// EthereumWalletLinker handles linking between Ethereum addresses and L2 wallets
type EthereumWalletLinker struct {
	ethClient    *smartcontracts.EthereumClient
	logger       *utils.Logger
	metrics      *utils.MetricsManager
	
	// Cache for challenges (address => challenge)
	challengeCache map[common.Address][32]byte
	challengeTimes map[common.Address]time.Time
	
	// Challenge expiration time (in seconds)
	challengeExpiration time.Duration
	
	mu             sync.RWMutex
}

// NewEthereumWalletLinker creates a new Ethereum wallet linker
func NewEthereumWalletLinker(
	ethClient *smartcontracts.EthereumClient,
	logger *utils.Logger,
	metrics *utils.MetricsManager,
) *EthereumWalletLinker {
	return &EthereumWalletLinker{
		ethClient:          ethClient,
		logger:             logger.WithComponent("EthereumWalletLinker"),
		metrics:            metrics,
		challengeCache:     make(map[common.Address][32]byte),
		challengeTimes:     make(map[common.Address]time.Time),
		challengeExpiration: 1 * time.Hour, // Same as contract (1 hour)
	}
}

// GenerateChallenge creates a challenge for an Ethereum address to sign
func (l *EthereumWalletLinker) GenerateChallenge(ctx context.Context, ethAddressHex string) (string, error) {
	// Validate the Ethereum address
	if !common.IsHexAddress(ethAddressHex) {
		return "", ErrInvalidEthereumAddress
	}
	ethAddress := common.HexToAddress(ethAddressHex)

	startTime := time.Now()
	l.logger.Info("Generating challenge", "eth_address", ethAddress.Hex())

	// Generate challenge via smart contract
	challengeHash, err := l.ethClient.GenerateChallenge(ctx, ethAddress)
	if err != nil {
		l.logger.Error("Failed to generate challenge", "error", err, "eth_address", ethAddress.Hex())
		return "", fmt.Errorf("failed to generate challenge: %w", err)
	}

	// Cache the challenge
	l.mu.Lock()
	l.challengeCache[ethAddress] = challengeHash
	l.challengeTimes[ethAddress] = time.Now()
	l.mu.Unlock()

	// Convert to hex string
	challengeHex := hexutil.Encode(challengeHash[:])

	// Record metrics
	if l.metrics != nil {
		l.metrics.IncCounter("challenges_generated_total", "chain", "ethereum")
		l.metrics.ObserveHistogram("challenge_generation_duration_seconds", time.Since(startTime).Seconds())
	}

	l.logger.Info("Challenge generated successfully", 
		"eth_address", ethAddress.Hex(), 
		"challenge", challengeHex,
		"duration", time.Since(startTime))

	return challengeHex, nil
}

// VerifySignature verifies that a signature was created by the owner of an Ethereum address
func (l *EthereumWalletLinker) VerifySignature(ctx context.Context, ethAddressHex, signatureHex string) (bool, error) {
	// Validate the Ethereum address
	if !common.IsHexAddress(ethAddressHex) {
		return false, ErrInvalidEthereumAddress
	}
	ethAddress := common.HexToAddress(ethAddressHex)

	// Validate the signature
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %w", err)
	}

	// Get the challenge from cache
	l.mu.RLock()
	challengeHash, exists := l.challengeCache[ethAddress]
	challengeTime, timeExists := l.challengeTimes[ethAddress]
	l.mu.RUnlock()

	if !exists || !timeExists {
		return false, ErrChallengeNotFound
	}

	// Check if the challenge has expired
	if time.Since(challengeTime) > l.challengeExpiration {
		return false, ErrChallengeExpired
	}

	// Verify the signature using Ethereum's personal sign format
	// Following EIP-191 standard for Ethereum signed data
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n32%s", string(challengeHash[:]))
	msgHash := crypto.Keccak256Hash([]byte(msg))
	
	// The signature needs to be in the proper format
	// If it's 65 bytes (common Ethereum signature), the last byte is the recovery ID
	if len(signature) != 65 {
		return false, fmt.Errorf("invalid signature length: %d", len(signature))
	}
	
	// Adjust V value for compatibility with Ethereum's ecrecover
	signature[64] = signature[64] - 27

	// Get the public key associated with the signature
	pubKey, err := crypto.Ecrecover(msgHash.Bytes(), signature)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}
	
	// Convert to ECDSA public key
	ecdsaPubKey, err := crypto.UnmarshalPubkey(pubKey)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal public key: %w", err)
	}
	
	// Derive address from public key
	recoveredAddress := crypto.PubkeyToAddress(*ecdsaPubKey)
	
	// Compare with the expected address
	isValid := recoveredAddress == ethAddress
	
	l.logger.Info("Signature verification result", 
		"eth_address", ethAddress.Hex(), 
		"recovered_address", recoveredAddress.Hex(),
		"is_valid", isValid)
	
	// Record metrics
	if l.metrics != nil {
		l.metrics.IncCounter("signatures_verified_total", "chain", "ethereum", "result", fmt.Sprintf("%t", isValid))
	}
	
	return isValid, nil
}

// LinkWallet links an Ethereum address to an L2 wallet
func (l *EthereumWalletLinker) LinkWallet(ctx context.Context, ethAddressHex string, l2WalletHex string, signatureHex string) error {
	// Validate the Ethereum address
	if !common.IsHexAddress(ethAddressHex) {
		return ErrInvalidEthereumAddress
	}
	ethAddress := common.HexToAddress(ethAddressHex)
	
	// Validate the L2 wallet (should be 32 bytes in hex)
	l2WalletBytes, err := hexutil.Decode(l2WalletHex)
	if err != nil || len(l2WalletBytes) != 32 {
		return fmt.Errorf("invalid L2 wallet format: %w", err)
	}
	
	// Convert to [32]byte
	var l2Wallet [32]byte
	copy(l2Wallet[:], l2WalletBytes)
	
	// Validate the signature
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return fmt.Errorf("invalid signature format: %w", err)
	}
	
	startTime := time.Now()
	l.logger.Info("Linking wallet", 
		"eth_address", ethAddress.Hex(), 
		"l2_wallet", l2WalletHex)
	
	// First verify the signature
	isValid, err := l.VerifySignature(ctx, ethAddressHex, signatureHex)
	if err != nil {
		l.logger.Error("Signature verification failed", "error", err)
		return err
	}
	
	if !isValid {
		return ErrInvalidSignature
	}
	
	// Check if either address is already linked
	isAlreadyLinked, err := l.IsAddressLinked(ctx, ethAddressHex)
	if err != nil {
		return fmt.Errorf("failed to check if address is linked: %w", err)
	}
	
	if isAlreadyLinked {
		return ErrAddressAlreadyLinked
	}
	
	isL2WalletLinked, err := l.IsL2WalletLinked(ctx, l2WalletHex)
	if err != nil {
		return fmt.Errorf("failed to check if L2 wallet is linked: %w", err)
	}
	
	if isL2WalletLinked {
		return ErrL2WalletAlreadyLinked
	}
	
	// Link the wallet via smart contract
	err = l.ethClient.LinkWallet(ctx, l2Wallet, signature)
	if err != nil {
		l.logger.Error("Failed to link wallet", "error", err)
		return fmt.Errorf("failed to link wallet: %w", err)
	}
	
	// Record metrics
	if l.metrics != nil {
		l.metrics.IncCounter("wallets_linked_total", "chain", "ethereum")
		l.metrics.ObserveHistogram("wallet_linking_duration_seconds", time.Since(startTime).Seconds())
	}
	
	l.logger.Info("Wallet linked successfully", 
		"eth_address", ethAddress.Hex(), 
		"l2_wallet", l2WalletHex,
		"duration", time.Since(startTime))
	
	return nil
}

// IsAddressLinked checks if an Ethereum address is already linked to an L2 wallet
func (l *EthereumWalletLinker) IsAddressLinked(ctx context.Context, ethAddressHex string) (bool, error) {
	// Validate the Ethereum address
	if !common.IsHexAddress(ethAddressHex) {
		return false, ErrInvalidEthereumAddress
	}
	ethAddress := common.HexToAddress(ethAddressHex)
	
	l.logger.Debug("Checking if address is linked", "eth_address", ethAddress.Hex())
	
	// Get the linked L2 wallet from the contract
	l2Wallet, err := l.ethClient.GetLinkedL2Wallet(ctx, ethAddress)
	if err != nil {
		// Special case: if the error indicates "no linked wallet", return false without error
		if err.Error() == "execution reverted" || err.Error() == "no L2 wallet found" {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if address is linked: %w", err)
	}
	
	// Check if the returned L2 wallet is empty (all zeros)
	isZero := true
	for _, b := range l2Wallet {
		if b != 0 {
			isZero = false
			break
		}
	}
	
	isLinked := !isZero
	l.logger.Debug("Address link check result", 
		"eth_address", ethAddress.Hex(), 
		"is_linked", isLinked)
	
	return isLinked, nil
}

// IsL2WalletLinked checks if an L2 wallet is already linked to an Ethereum address
func (l *EthereumWalletLinker) IsL2WalletLinked(ctx context.Context, l2WalletHex string) (bool, error) {
	// Validate the L2 wallet (should be 32 bytes in hex)
	l2WalletBytes, err := hexutil.Decode(l2WalletHex)
	if err != nil || len(l2WalletBytes) != 32 {
		return false, fmt.Errorf("invalid L2 wallet format: %w", err)
	}
	
	// Convert to [32]byte
	var l2Wallet [32]byte
	copy(l2Wallet[:], l2WalletBytes)
	
	l.logger.Debug("Checking if L2 wallet is linked", "l2_wallet", l2WalletHex)
	
	// Get the linked Ethereum address from the contract
	ethAddress, err := l.ethClient.GetLinkedL1Address(ctx, l2Wallet)
	if err != nil {
		// Special case: if the error indicates "no linked address", return false without error
		if err.Error() == "execution reverted" || err.Error() == "no L1 address found" {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if L2 wallet is linked: %w", err)
	}
	
	// Check if the returned address is non-zero
	isLinked := ethAddress != common.Address{}
	
	l.logger.Debug("L2 wallet link check result", 
		"l2_wallet", l2WalletHex, 
		"is_linked", isLinked,
		"eth_address", ethAddress.Hex())
	
	return isLinked, nil
}

// GetLinkedL2Wallet gets the L2 wallet linked to an Ethereum address
func (l *EthereumWalletLinker) GetLinkedL2Wallet(ctx context.Context, ethAddressHex string) (string, error) {
	// Validate the Ethereum address
	if !common.IsHexAddress(ethAddressHex) {
		return "", ErrInvalidEthereumAddress
	}
	ethAddress := common.HexToAddress(ethAddressHex)
	
	l.logger.Debug("Getting linked L2 wallet", "eth_address", ethAddress.Hex())
	
	// Get the linked L2 wallet from the contract
	l2Wallet, err := l.ethClient.GetLinkedL2Wallet(ctx, ethAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get linked L2 wallet: %w", err)
	}
	
	// Check if the returned L2 wallet is empty (all zeros)
	isZero := true
	for _, b := range l2Wallet {
		if b != 0 {
			isZero = false
			break
		}
	}
	
	if isZero {
		return "", fmt.Errorf("no L2 wallet linked to this address")
	}
	
	// Convert to hex string
	l2WalletHex := hexutil.Encode(l2Wallet[:])
	
	l.logger.Debug("Got linked L2 wallet", 
		"eth_address", ethAddress.Hex(), 
		"l2_wallet", l2WalletHex)
	
	return l2WalletHex, nil
}

// GetLinkedEthereumAddress gets the Ethereum address linked to an L2 wallet
func (l *EthereumWalletLinker) GetLinkedEthereumAddress(ctx context.Context, l2WalletHex string) (string, error) {
	// Validate the L2 wallet (should be 32 bytes in hex)
	l2WalletBytes, err := hexutil.Decode(l2WalletHex)
	if err != nil || len(l2WalletBytes) != 32 {
		return "", fmt.Errorf("invalid L2 wallet format: %w", err)
	}
	
	// Convert to [32]byte
	var l2Wallet [32]byte
	copy(l2Wallet[:], l2WalletBytes)
	
	l.logger.Debug("Getting linked Ethereum address", "l2_wallet", l2WalletHex)
	
	// Get the linked Ethereum address from the contract
	ethAddress, err := l.ethClient.GetLinkedL1Address(ctx, l2Wallet)
	if err != nil {
		return "", fmt.Errorf("failed to get linked Ethereum address: %w", err)
	}
	
	// Check if the returned address is non-zero
	if ethAddress == (common.Address{}) {
		return "", fmt.Errorf("no Ethereum address linked to this L2 wallet")
	}
	
	l.logger.Debug("Got linked Ethereum address", 
		"l2_wallet", l2WalletHex, 
		"eth_address", ethAddress.Hex())
	
	return ethAddress.Hex(), nil
}

// VerifyCreditScore verifies the credit score of an L2 wallet
func (l *EthereumWalletLinker) VerifyCreditScore(ctx context.Context, l2WalletHex string, threshold float64) (bool, error) {
	// Validate the L2 wallet (should be 32 bytes in hex)
	l2WalletBytes, err := hexutil.Decode(l2WalletHex)
	if err != nil || len(l2WalletBytes) != 32 {
		return false, fmt.Errorf("invalid L2 wallet format: %w", err)
	}
	
	// Convert to [32]byte
	var l2Wallet [32]byte
	copy(l2Wallet[:], l2WalletBytes)
	
	startTime := time.Now()
	l.logger.Info("Verifying credit score", "l2_wallet", l2WalletHex, "threshold", threshold)
	
	// Convert float threshold to fixed-point representation (18 decimals)
	thresholdScaled := new(big.Int).Mul(
		big.NewInt(int64(threshold * 100)), 
		big.NewInt(10_000_000_000_000_000)) // 10^16
	
	// Define threshold vector
	// In a standard 2D vector space used for cosine similarity
	thresholdX := thresholdScaled
	thresholdY := big.NewInt(1_000_000_000_000_000_000) // 1.0 with 18 decimals (10^18)
	
	// Define similarity threshold (e.g., 0.9 or 90%)
	similarityThreshold := big.NewInt(900_000_000_000_000_000) // 0.9 with 18 decimals
	
	// Verify credit score via smart contract
	_, passed, err := l.ethClient.VerifyCreditScore(ctx, l2Wallet, thresholdX, thresholdY, similarityThreshold)
	if err != nil {
		l.logger.Error("Failed to verify credit score", "error", err)
		return false, fmt.Errorf("failed to verify credit score: %w", err)
	}
	
	// Record metrics
	if l.metrics != nil {
		l.metrics.IncCounter("credit_score_verifications_total", "chain", "ethereum", "result", fmt.Sprintf("%t", passed))
		l.metrics.ObserveHistogram("credit_score_verification_duration_seconds", time.Since(startTime).Seconds())
	}
	
	l.logger.Info("Credit score verification result", 
		"l2_wallet", l2WalletHex, 
		"threshold", threshold,
		"passed", passed,
		"duration", time.Since(startTime))
	
	return passed, nil
}

// ProvideTransactionFeedback provides feedback after a transaction
func (l *EthereumWalletLinker) ProvideTransactionFeedback(ctx context.Context, l2WalletHex string, isTrusted bool) error {
	// Validate the L2 wallet (should be 32 bytes in hex)
	l2WalletBytes, err := hexutil.Decode(l2WalletHex)
	if err != nil || len(l2WalletBytes) != 32 {
		return fmt.Errorf("invalid L2 wallet format: %w", err)
	}
	
	// Convert to [32]byte
	var l2Wallet [32]byte
	copy(l2Wallet[:], l2WalletBytes)
	
	startTime := time.Now()
	l.logger.Info("Providing transaction feedback", "l2_wallet", l2WalletHex, "is_trusted", isTrusted)
	
	// Provide feedback via smart contract
	err = l.ethClient.ProvideTransactionFeedback(ctx, l2Wallet, isTrusted)
	if err != nil {
		l.logger.Error("Failed to provide transaction feedback", "error", err)
		return fmt.Errorf("failed to provide transaction feedback: %w", err)
	}
	
	// Record metrics
	if l.metrics != nil {
		l.metrics.IncCounter("transaction_feedback_total", "chain", "ethereum", "trusted", fmt.Sprintf("%t", isTrusted))
		l.metrics.ObserveHistogram("transaction_feedback_duration_seconds", time.Since(startTime).Seconds())
	}
	
	l.logger.Info("Transaction feedback provided", 
		"l2_wallet", l2WalletHex, 
		"is_trusted", isTrusted,
		"duration", time.Since(startTime))
	
	return nil
}

// CheckFeedbackRequirement checks if feedback is required
func (l *EthereumWalletLinker) CheckFeedbackRequirement(ctx context.Context, ethAddressHex string) (bool, error) {
	// Validate the Ethereum address
	if !common.IsHexAddress(ethAddressHex) {
		return false, ErrInvalidEthereumAddress
	}
	ethAddress := common.HexToAddress(ethAddressHex)
	
	l.logger.Debug("Checking feedback requirement", "eth_address", ethAddress.Hex())
	
	// Check feedback requirement via smart contract
	requiresFeedback, err := l.ethClient.CheckFeedbackRequirement(ctx, ethAddress)
	if err != nil {
		return false, fmt.Errorf("failed to check feedback requirement: %w", err)
	}
	
	l.logger.Debug("Feedback requirement check result", 
		"eth_address", ethAddress.Hex(), 
		"requires_feedback", requiresFeedback)
	
	return requiresFeedback, nil
}

// GenerateEthereumKeyPair generates a new Ethereum key pair for testing
func GenerateEthereumKeyPair() (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to generate key pair: %w", err)
	}
	
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, fmt.Errorf("error casting public key to ECDSA")
	}
	
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	
	return privateKey, address, nil
}