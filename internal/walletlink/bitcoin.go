package walletlink

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
)

// Common errors
var (
	ErrInvalidBitcoinAddress = errors.New("invalid Bitcoin address")
	ErrInvalidBitcoinSignature = errors.New("invalid Bitcoin signature")
	ErrBitcoinChallengeExpired = errors.New("Bitcoin challenge has expired")
	ErrBitcoinChallengeNotFound = errors.New("no challenge found for Bitcoin address")
	ErrBitcoinRPCError = errors.New("Bitcoin RPC error")
)

// BitcoinConfig contains configuration for Bitcoin interactions
type BitcoinConfig struct {
	RPCHost            string `json:"rpc_host"`
	RPCPort            string `json:"rpc_port"`
	RPCUser            string `json:"rpc_user"`
	RPCPassword        string `json:"rpc_password"`
	Network            string `json:"network"`
	ChallengeExpiration int    `json:"challenge_expiration_minutes"`
	ChainID            uint64  `json:"chain_id"` // Chain ID for Bitcoin (e.g., 502)
}

// BitcoinWalletLinker handles linking between Bitcoin addresses and L2 wallets
type BitcoinWalletLinker struct {
	client      *rpcclient.Client
	logger      *utils.Logger
	metrics     *utils.MetricsManager
	chainParams *chaincfg.Params
	config      *BitcoinConfig
	
	// Cache for challenges (btc address => challenge)
	challengeCache map[string]string
	challengeTimes map[string]time.Time
	
	// Challenge expiration time (in seconds)
	challengeExpiration time.Duration
	
	mu sync.RWMutex
}

// NewBitcoinWalletLinker creates a new Bitcoin wallet linker
func NewBitcoinWalletLinker(
	config *BitcoinConfig,
	logger *utils.Logger,
	metrics *utils.MetricsManager,
) (*BitcoinWalletLinker, error) {
	logger = logger.WithComponent("BitcoinWalletLinker")

	// Set chain parameters based on network
	var chainParams *chaincfg.Params
	switch strings.ToLower(config.Network) {
	case "mainnet":
		chainParams = &chaincfg.MainNetParams
	case "testnet":
		chainParams = &chaincfg.TestNet3Params
	case "regtest":
		chainParams = &chaincfg.RegressionNetParams
	default:
		return nil, fmt.Errorf("unsupported Bitcoin network: %s", config.Network)
	}

	// Connect to Bitcoin node via RPC
	rpcConf := &rpcclient.ConnConfig{
		Host:         fmt.Sprintf("%s:%s", config.RPCHost, config.RPCPort),
		User:         config.RPCUser,
		Pass:         config.RPCPassword,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(rpcConf, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Bitcoin RPC client: %w", err)
	}

	// Test connection
	blockCount, err := client.GetBlockCount()
	if err != nil {
		client.Shutdown()
		return nil, fmt.Errorf("failed to connect to Bitcoin node: %w", err)
	}

	logger.Info("Connected to Bitcoin node", 
		"network", config.Network, 
		"block_count", blockCount)

	linker := &BitcoinWalletLinker{
		client:             client,
		logger:             logger,
		metrics:            metrics,
		chainParams:        chainParams,
		config:             config,
		challengeCache:     make(map[string]string),
		challengeTimes:     make(map[string]time.Time),
		challengeExpiration: time.Duration(config.ChallengeExpiration) * time.Minute,
	}

	return linker, nil
}

// ValidateBitcoinAddress validates a Bitcoin address
func (l *BitcoinWalletLinker) ValidateBitcoinAddress(address string) (bool, error) {
	_, err := btcutil.DecodeAddress(address, l.chainParams)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// GenerateChallenge creates a challenge for a Bitcoin address to sign
func (l *BitcoinWalletLinker) GenerateChallenge(ctx context.Context, btcAddress string) (string, error) {
	// Validate the Bitcoin address
	valid, err := l.ValidateBitcoinAddress(btcAddress)
	if err != nil {
		return "", fmt.Errorf("error validating Bitcoin address: %w", err)
	}
	if !valid {
		return "", ErrInvalidBitcoinAddress
	}

	startTime := time.Now()
	l.logger.Info("Generating challenge", "btc_address", btcAddress)

	// Generate a random challenge
	challengeBytes := make([]byte, 32)
	_, err = rand.Read(challengeBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random challenge: %w", err)
	}

	// Hash the challenge for consistency and convert to string
	challengeHash := sha256.Sum256(challengeBytes)
	challenge := "COSINE Bitcoin Challenge: " + hex.EncodeToString(challengeHash[:])

	// Cache the challenge
	l.mu.Lock()
	l.challengeCache[btcAddress] = challenge
	l.challengeTimes[btcAddress] = time.Now()
	l.mu.Unlock()

	// Record metrics
	if l.metrics != nil {
		l.metrics.IncCounter("challenges_generated_total", "chain", "bitcoin")
		l.metrics.ObserveHistogram("challenge_generation_duration_seconds", time.Since(startTime).Seconds())
	}

	l.logger.Info("Challenge generated successfully", 
		"btc_address", btcAddress, 
		"challenge", challenge,
		"duration", time.Since(startTime))

	return challenge, nil
}

// VerifySignature verifies that a signature was created by the owner of a Bitcoin address
func (l *BitcoinWalletLinker) VerifySignature(ctx context.Context, btcAddress, signature string) (bool, error) {
	// Validate the Bitcoin address
	valid, err := l.ValidateBitcoinAddress(btcAddress)
	if err != nil {
		return false, fmt.Errorf("error validating Bitcoin address: %w", err)
	}
	if !valid {
		return false, ErrInvalidBitcoinAddress
	}

	// Get the challenge from cache
	l.mu.RLock()
	challenge, exists := l.challengeCache[btcAddress]
	challengeTime, timeExists := l.challengeTimes[btcAddress]
	l.mu.RUnlock()

	if !exists || !timeExists {
		return false, ErrBitcoinChallengeNotFound
	}

	// Check if the challenge has expired
	if time.Since(challengeTime) > l.challengeExpiration {
		return false, ErrBitcoinChallengeExpired
	}

	startTime := time.Now()
	l.logger.Info("Verifying signature", 
		"btc_address", btcAddress, 
		"challenge", challenge)

	// Use RPC to verify the signature (if connected to a Bitcoin node)
	// This is the most reliable way to verify Bitcoin signatures
	valid, err = l.verifySignatureViaRPC(btcAddress, challenge, signature)
	if err != nil {
		// If RPC fails, fall back to local verification
		l.logger.Warn("RPC signature verification failed, falling back to local verification", 
			"error", err)
		valid, err = l.verifySignatureLocally(btcAddress, challenge, signature)
		if err != nil {
			return false, err
		}
	}

	// Record metrics
	if l.metrics != nil {
		l.metrics.IncCounter("signatures_verified_total", "chain", "bitcoin", "result", fmt.Sprintf("%t", valid))
		l.metrics.ObserveHistogram("signature_verification_duration_seconds", time.Since(startTime).Seconds())
	}

	l.logger.Info("Signature verification result", 
		"btc_address", btcAddress, 
		"is_valid", valid,
		"duration", time.Since(startTime))

	return valid, nil
}

// verifySignatureViaRPC verifies a signature using the Bitcoin RPC API
func (l *BitcoinWalletLinker) verifySignatureViaRPC(address, message, signature string) (bool, error) {
    addr, err := btcutil.DecodeAddress(address, l.chainParams)
    if err != nil {
        return false, fmt.Errorf("failed to decode Bitcoin address: %w", err)
    }
    valid, err := l.client.VerifyMessage(addr, signature, message)
    if err != nil {
        return false, fmt.Errorf("%w: %v", ErrBitcoinRPCError, err)
    }
    return valid, nil
}


// verifySignatureLocally verifies a Bitcoin signature locally (as a fallback)
func (l *BitcoinWalletLinker) verifySignatureLocally(address, message, signatureBase64 string) (bool, error) {
	// Decode the address
	addr, err := btcutil.DecodeAddress(address, l.chainParams)
	if err != nil {
		return false, fmt.Errorf("failed to decode address: %w", err)
	}

	// Decode the signature from base64
	signatureBytes, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %w", err)
	}

	// Bitcoin message signing prefixes the message with a standard header
	// and double-hashes it
	var buf bytes.Buffer
	wire.WriteVarString(&buf, 0, "Bitcoin Signed Message:\n")
	wire.WriteVarString(&buf, 0, message)
	messageHash := chainhash.DoubleHashB(buf.Bytes())

	// Remove header byte and recover the public key
	if len(signatureBytes) < 65 {
		return false, fmt.Errorf("signature too short")
	}

	signature := signatureBytes[1:]
	compressed := signatureBytes[0] >= 31
	

	// Recover the public key
	// Note: This is a simplified approach - actual recovery would handle
	// all recovery ID possibilities
	pubKey, _, err := ecdsa.RecoverCompact(signature, messageHash)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	// Calculate the recovered address
	var pubKeyHash []byte
	if compressed {
		pubKeyHash = btcutil.Hash160(pubKey.SerializeCompressed())
	} else {
		pubKeyHash = btcutil.Hash160(pubKey.SerializeUncompressed())
	}

	var recoveredAddr btcutil.Address
	switch addr.(type) {
	case *btcutil.AddressPubKeyHash:
		recoveredAddr, err = btcutil.NewAddressPubKeyHash(pubKeyHash, l.chainParams)
	case *btcutil.AddressScriptHash:
		// P2SH addresses need special handling
		return false, fmt.Errorf("P2SH address verification not supported locally")
	default:
		return false, fmt.Errorf("unsupported address type")
	}

	if err != nil {
		return false, fmt.Errorf("failed to create recovered address: %w", err)
	}

	// Compare the recovered address with the expected address
	return recoveredAddr.EncodeAddress() == address, nil
}

// LinkWallet links a Bitcoin address to an L2 wallet
func (l *BitcoinWalletLinker) LinkWallet(ctx context.Context, btcAddress string, l2WalletHex string, signatureBase64 string) error {
	// Validate the Bitcoin address
	valid, err := l.ValidateBitcoinAddress(btcAddress)
	if err != nil {
		return fmt.Errorf("error validating Bitcoin address: %w", err)
	}
	if !valid {
		return ErrInvalidBitcoinAddress
	}

	// Validate the L2 wallet (should be 32 bytes in hex)
	l2WalletBytes, err := hex.DecodeString(strings.TrimPrefix(l2WalletHex, "0x"))
	if err != nil || len(l2WalletBytes) != 32 {
		return fmt.Errorf("invalid L2 wallet format: %w", err)
	}

	startTime := time.Now()
	l.logger.Info("Linking wallet", 
		"btc_address", btcAddress, 
		"l2_wallet", l2WalletHex)

	// First verify the signature
	isValid, err := l.VerifySignature(ctx, btcAddress, signatureBase64)
	if err != nil {
		l.logger.Error("Signature verification failed", "error", err)
		return err
	}

	if !isValid {
		return ErrInvalidBitcoinSignature
	}

	// The actual linking of Bitcoin address to L2 wallet would be handled by the mapping manager
	// We'll return success here since the verification passed
	// address_mapping.go will handle the actual database entry

	// Record metrics
	if l.metrics != nil {
		l.metrics.IncCounter("wallets_linked_total", "chain", "bitcoin")
		l.metrics.ObserveHistogram("wallet_linking_duration_seconds", time.Since(startTime).Seconds())
	}

	l.logger.Info("Wallet verified and ready for linking", 
		"btc_address", btcAddress, 
		"l2_wallet", l2WalletHex,
		"duration", time.Since(startTime))

	return nil
}

// GetBitcoinTransactions gets recent transactions for a Bitcoin address
func (l *BitcoinWalletLinker) GetBitcoinTransactions(ctx context.Context, btcAddress string) ([]map[string]interface{}, error) {
	// In a real implementation, this would use Bitcoin RPC or an indexer to get transactions
	// For this example, we'll just return dummy data
	
	l.logger.Info("Getting Bitcoin transactions", "btc_address", btcAddress)
	
	// For demonstration purposes, return an empty slice
	return []map[string]interface{}{}, nil
}

// Close closes the Bitcoin wallet linker and its RPC connection
func (l *BitcoinWalletLinker) Close() {
	if l.client != nil {
		l.client.Shutdown()
		l.logger.Info("Bitcoin wallet linker closed")
	}
}