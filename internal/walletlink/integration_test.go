package walletlink_test

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
	"github.com/IWALINK/cosine/internal/walletlink"
	"github.com/IWALINK/cosine/pkg/smartcontracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test configuration
const (
	// Ethereum configuration (for local testing)
	ethRPCEndpoint   = "http://localhost:8545"
	ethChainID       = 31337 // Hardhat local network
	
	// Bitcoin configuration (for local testing)
	btcRPCHost       = "localhost"
	btcRPCPort       = "18443"
	btcRPCUser       = "bitcoinrpc"
	btcRPCPassword   = "password"
	btcNetwork       = "regtest"
)

// Mock wallet state trie for testing
type mockWalletStateTrie struct {
	wallets map[string]*storage.WalletState
}

func newMockWalletStateTrie() *mockWalletStateTrie {
	return &mockWalletStateTrie{
		wallets: make(map[string]*storage.WalletState),
	}
}

func (m *mockWalletStateTrie) GetWallet(address string) *storage.WalletState {
	wallet, exists := m.wallets[address]
	if !exists {
		wallet = storage.NewWalletState(address)
		m.wallets[address] = wallet
	}
	return wallet
}

func (m *mockWalletStateTrie) UpdateWallet(wallet *storage.WalletState) {
	m.wallets[wallet.Address] = wallet
}

func (m *mockWalletStateTrie) GetAllWallets() []*storage.WalletState {
	wallets := make([]*storage.WalletState, 0, len(m.wallets))
	for _, wallet := range m.wallets {
		wallets = append(wallets, wallet)
	}
	return wallets
}

// setupTestEnvironment sets up the test environment for chain integration tests
func setupTestEnvironment(t *testing.T) (
	*walletlink.EthereumWalletLinker, 
	*walletlink.BitcoinWalletLinker, 
	*walletlink.ProofVerifier, 
	*walletlink.AddressMappingManager,
	*utils.Logger,
	*utils.MetricsManager,
) {
	// Initialize logger
	loggerOptions := utils.DefaultLoggerOptions()
	loggerOptions.Level = utils.DebugLevel
	logger, err := utils.NewLogger(loggerOptions)
	require.NoError(t, err, "Failed to initialize logger")

	// Initialize metrics
	metricsOptions := utils.DefaultMetricsOptions()
	metrics := utils.NewMetricsManager(metricsOptions)

	// Setup Ethereum client
	ethConfig := &smartcontracts.ContractConfig{
		RPCEndpoint:          ethRPCEndpoint,
		ChainID:              ethChainID,
		WalletLinkingAddress: "", // Will be set by test if needed
		RetryAttempts:        3,
		RetryDelay:           100,
	}

	// Create a key for testing
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err, "Failed to generate private key")
	ethConfig.OperatorPrivateKey = hex.EncodeToString(crypto.FromECDSA(privateKey))

	ethClient, err := smartcontracts.New(ethConfig, logger, metrics)
	if err != nil {
		t.Skipf("Skipping Ethereum tests: %v", err)
		return nil, nil, nil, nil, logger, metrics
	}

	// Create Ethereum wallet linker
	ethLinker := walletlink.NewEthereumWalletLinker(ethClient, logger, metrics)

	// Setup Bitcoin wallet linker
	btcConfig := &walletlink.BitcoinConfig{
		RPCHost:             btcRPCHost,
		RPCPort:             btcRPCPort,
		RPCUser:             btcRPCUser,
		RPCPassword:         btcRPCPassword,
		Network:             btcNetwork,
		ChallengeExpiration: 60, // 60 minutes
		ChainID:             502, // Bitcoin chain ID
	}

	var btcLinker *walletlink.BitcoinWalletLinker
	btcLinker, err = walletlink.NewBitcoinWalletLinker(btcConfig, logger, metrics)
	if err != nil {
		t.Logf("Bitcoin tests will be skipped: %v", err)
		btcLinker = nil
	}

	// Create proof verifier
	proofVerifier := walletlink.NewProofVerifier(ethLinker, btcLinker, logger, metrics)

	// Setup mock wallet state trie
	walletTrie := storage.NewWalletStateTrie(logger)

	// Create L1 mapping trie
	l1Trie := storage.NewLinkedL1Trie(logger, walletTrie)

	// Create address mapping manager
	mappingManager := walletlink.NewAddressMappingManager(l1Trie, proofVerifier, logger, metrics)

	return ethLinker, btcLinker, proofVerifier, mappingManager, logger, metrics
}

// Test Ethereum wallet linking
func TestEthereumWalletLinking(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping integration tests")
	}

	ethLinker, _, _, mappingManager, logger, _ := setupTestEnvironment(t)

	if ethLinker == nil {
		t.Skip("Ethereum linker not available")
	}

	// Create a testing context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Generate a test Ethereum key pair
	privateKey, ethAddress, err := walletlink.GenerateEthereumKeyPair()
	require.NoError(t, err, "Failed to generate Ethereum key pair")

	// Create a test L2 wallet
	l2Wallet := make([]byte, 32)
	l2Wallet[0] = 0x01
	l2WalletHex := "0x" + hex.EncodeToString(l2Wallet)

	// Step 1: Generate a challenge
	challenge, err := ethLinker.GenerateChallenge(ctx, ethAddress.Hex())
	require.NoError(t, err, "Failed to generate challenge")
	assert.NotEmpty(t, challenge, "Challenge should not be empty")

	logger.Info("Generated challenge", "challenge", challenge)

	// Step 2: Sign the challenge
	// Convert hex challenge to bytes for signing
	challengeBytes, err := hexutil.Decode(challenge)
	require.NoError(t, err, "Failed to decode challenge")

	// Follow Ethereum personal sign format (prefixed with "\x19Ethereum Signed Message:\n" + length of message)
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n32%s", string(challengeBytes))
	msgHash := crypto.Keccak256([]byte(msg))
	sig, err := crypto.Sign(msgHash, privateKey)
	require.NoError(t, err, "Failed to sign message")

	// Convert signature to hex
	sigHex := hexutil.Encode(sig)

	// Step 3: Verify the signature
	valid, err := ethLinker.VerifySignature(ctx, ethAddress.Hex(), sigHex)
	require.NoError(t, err, "Failed to verify signature")
	assert.True(t, valid, "Signature should be valid")

	logger.Info("Verified signature", "valid", valid)

	// Step 4: Link the wallet
	err = mappingManager.LinkAddressWithVerification(
		ctx,
		walletlink.ChainEthereum,
		uint64(ethChainID),
		ethAddress.Hex(),
		l2WalletHex,
		sigHex,
	)
	require.NoError(t, err, "Failed to link wallet")

	// Step 5: Check if the address is linked
	isLinked, err := mappingManager.IsAddressLinked(ctx, uint64(ethChainID), ethAddress.Hex())
	require.NoError(t, err, "Failed to check if address is linked")
	assert.True(t, isLinked, "Address should be linked")

	// Step 6: Get the linked L2 wallet
	linkedL2Wallet, err := mappingManager.GetLinkedL2Wallet(ctx, uint64(ethChainID), ethAddress.Hex())
	require.NoError(t, err, "Failed to get linked L2 wallet")
	assert.Equal(t, l2WalletHex, linkedL2Wallet, "Linked L2 wallet should match")

	// Step 7: Get the linked L1 address
	linkedEthAddress, err := mappingManager.GetLinkedL1Address(ctx, uint64(ethChainID), l2WalletHex)
	require.NoError(t, err, "Failed to get linked L1 address")
	assert.Equal(t, ethAddress.Hex(), linkedEthAddress, "Linked Ethereum address should match")

	// Step 8: Get all linked addresses
	linkedAddresses, err := mappingManager.GetAllLinkedAddresses(ctx, l2WalletHex)
	require.NoError(t, err, "Failed to get all linked addresses")
	assert.Len(t, linkedAddresses, 1, "Should have one linked address")
	assert.Equal(t, uint64(ethChainID), linkedAddresses[0].ChainID, "Chain ID should match")
	assert.Equal(t, ethAddress.Hex(), linkedAddresses[0].L1Address, "L1 address should match")

	// Step 9: Get L2 wallet info
	walletInfo, err := mappingManager.GetL2WalletInfo(ctx, l2WalletHex)
	require.NoError(t, err, "Failed to get L2 wallet info")
	assert.Equal(t, l2WalletHex, walletInfo.L2Wallet, "L2 wallet should match")
	assert.Len(t, walletInfo.LinkedChains, 1, "Should have one linked chain")
}

// Test the helper utility for Ethereum key generation and signing
func TestEthereumKeyUtilities(t *testing.T) {
	// Test key pair generation
	privateKey, address, err := walletlink.GenerateEthereumKeyPair()
	require.NoError(t, err, "Failed to generate Ethereum key pair")
	assert.NotNil(t, privateKey, "Private key should not be nil")
	assert.NotEqual(t, common.Address{}, address, "Address should not be empty")

	// Test signing and verification
	message := []byte("test message")
	msgHash := crypto.Keccak256Hash(message)
	
	// Sign the message
	signature, err := crypto.Sign(msgHash.Bytes(), privateKey)
	require.NoError(t, err, "Failed to sign message")
	
	// Verify the signature
	pubKey, err := crypto.Ecrecover(msgHash.Bytes(), signature)
	require.NoError(t, err, "Failed to recover public key")
	
	// Convert to ECDSA public key
	ecdsaPubKey, err := crypto.UnmarshalPubkey(pubKey)
	require.NoError(t, err, "Failed to unmarshal public key")
	
	// Derive address from public key
	recoveredAddress := crypto.PubkeyToAddress(*ecdsaPubKey)
	
	// Compare with the original address
	assert.Equal(t, address, recoveredAddress, "Recovered address should match original")
}

// Test utility for sign and verify Ethereum message
func signAndVerifyEthereumMessage(t *testing.T, privateKey *ecdsa.PrivateKey, message []byte) {
	// Hash the message
	messageHash := crypto.Keccak256Hash(message)
	
	// Sign the message
	signature, err := crypto.Sign(messageHash.Bytes(), privateKey)
	require.NoError(t, err, "Failed to sign message")
	
	// Get public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok, "Failed to cast public key to ECDSA")
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	
	// Verify the signature
	sigPublicKey, err := crypto.Ecrecover(messageHash.Bytes(), signature)
	require.NoError(t, err, "Failed to recover public key")
	
	// Convert to ECDSA public key
	ecdsaPubKey, err := crypto.UnmarshalPubkey(sigPublicKey)
	require.NoError(t, err, "Failed to unmarshal public key")
	
	// Derive address from public key
	recoveredAddress := crypto.PubkeyToAddress(*ecdsaPubKey)
	
	// Compare with the original address
	assert.Equal(t, address, recoveredAddress, "Recovered address should match original")
}

// If Bitcoin wallet linker is available, this test would verify Bitcoin wallet linking
func TestBitcoinWalletLinking(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TESTS") == "true" {
		t.Skip("Skipping integration tests")
	}

	_, btcLinker, _, _, _, _ := setupTestEnvironment(t)
	if btcLinker == nil {
		t.Skip("Bitcoin linker not available")
	}

	// Test would be similar to Ethereum but using Bitcoin addresses and signatures
	// Skipping detailed implementation as it requires a running Bitcoin node
	t.Log("Bitcoin wallet linking test would go here")
}

// Test chain type detection utility
func TestDetectChainType(t *testing.T) {
	testCases := []struct {
		name          string
		address       string
		expectedChain walletlink.ChainType
		expectedError bool
	}{
		{
			name:          "Valid Ethereum address",
			address:       "0x8ba1f109551bD432803012645Ac136ddd64DBA72",
			expectedChain: walletlink.ChainEthereum,
			expectedError: false,
		},
		{
			name:          "Valid Bitcoin mainnet address",
			address:       "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
			expectedChain: walletlink.ChainBitcoin,
			expectedError: false,
		},
		{
			name:          "Valid Bitcoin testnet address",
			address:       "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			expectedChain: walletlink.ChainBitcoin,
			expectedError: false,
		},
		{
			name:          "Invalid address",
			address:       "not-an-address",
			expectedChain: "",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chainType, err := walletlink.DetectChainType(tc.address)
			if tc.expectedError {
				assert.Error(t, err, "Expected an error for invalid address")
			} else {
				assert.NoError(t, err, "Expected no error for valid address")
				assert.Equal(t, tc.expectedChain, chainType, "Chain type should match")
			}
		})
	}
}