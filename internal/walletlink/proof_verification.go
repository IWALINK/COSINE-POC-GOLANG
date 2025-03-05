package walletlink

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
)

// Common errors
var (
	ErrUnsupportedChain     = errors.New("unsupported blockchain")
	ErrInvalidAddress       = errors.New("invalid address format")
	ErrVerificationFailed   = errors.New("proof verification failed")
	ErrAddressNotLinked     = errors.New("address is not linked to an L2 wallet")
	ErrL2WalletNotLinked    = errors.New("L2 wallet is not linked to an address")
	ErrNoLinker             = errors.New("no linker available for the specified chain")
)

// ChainType represents a blockchain type
type ChainType string

// Supported blockchain types
const (
	ChainEthereum ChainType = "ethereum"
	ChainBitcoin  ChainType = "bitcoin"
)

// ProofVerifier handles proof verification across different blockchains
type ProofVerifier struct {
	ethereumLinker *EthereumWalletLinker
	bitcoinLinker  *BitcoinWalletLinker
	logger         *utils.Logger
	metrics        *utils.MetricsManager
}

// NewProofVerifier creates a new proof verifier
func NewProofVerifier(
	ethereumLinker *EthereumWalletLinker,
	bitcoinLinker *BitcoinWalletLinker,
	logger *utils.Logger,
	metrics *utils.MetricsManager,
) *ProofVerifier {
	return &ProofVerifier{
		ethereumLinker: ethereumLinker,
		bitcoinLinker:  bitcoinLinker,
		logger:         logger.WithComponent("ProofVerifier"),
		metrics:        metrics,
	}
}

// GenerateChallenge generates a challenge for the specified chain and address
func (pv *ProofVerifier) GenerateChallenge(ctx context.Context, chain ChainType, address string) (string, error) {
	startTime := time.Now()
	pv.logger.Info("Generating challenge", "chain", chain, "address", address)

	var (
		challenge string
		err       error
	)

	switch chain {
	case ChainEthereum:
		if pv.ethereumLinker == nil {
			return "", ErrNoLinker
		}
		challenge, err = pv.ethereumLinker.GenerateChallenge(ctx, address)
	case ChainBitcoin:
		if pv.bitcoinLinker == nil {
			return "", ErrNoLinker
		}
		challenge, err = pv.bitcoinLinker.GenerateChallenge(ctx, address)
	default:
		return "", ErrUnsupportedChain
	}

	if err != nil {
		pv.logger.Error("Challenge generation failed", 
			"error", err, 
			"chain", chain, 
			"address", address)
		return "", fmt.Errorf("failed to generate challenge: %w", err)
	}

	// Record metrics
	if pv.metrics != nil {
		pv.metrics.IncCounter("challenges_generated_total", "chain", string(chain))
		pv.metrics.ObserveHistogram("challenge_generation_duration_seconds", time.Since(startTime).Seconds())
	}

	pv.logger.Info("Challenge generated successfully", 
		"chain", chain, 
		"address", address, 
		"challenge", challenge,
		"duration", time.Since(startTime))

	return challenge, nil
}

// VerifyProof verifies a proof for the specified chain, address, and L2 wallet
func (pv *ProofVerifier) VerifyProof(ctx context.Context, chain ChainType, address, l2WalletHex, signature string) (bool, error) {
	startTime := time.Now()
	pv.logger.Info("Verifying proof", 
		"chain", chain, 
		"address", address, 
		"l2_wallet", l2WalletHex)

	var (
		isValid bool
		err     error
	)

	switch chain {
	case ChainEthereum:
		if pv.ethereumLinker == nil {
			return false, ErrNoLinker
		}
		isValid, err = pv.ethereumLinker.VerifySignature(ctx, address, signature)
	case ChainBitcoin:
		if pv.bitcoinLinker == nil {
			return false, ErrNoLinker
		}
		isValid, err = pv.bitcoinLinker.VerifySignature(ctx, address, signature)
	default:
		return false, ErrUnsupportedChain
	}

	if err != nil {
		pv.logger.Error("Proof verification failed", 
			"error", err, 
			"chain", chain, 
			"address", address)
		return false, fmt.Errorf("failed to verify proof: %w", err)
	}

	// Record metrics
	if pv.metrics != nil {
		pv.metrics.IncCounter("proofs_verified_total", 
			"chain", string(chain), 
			"result", fmt.Sprintf("%t", isValid))
		pv.metrics.ObserveHistogram("proof_verification_duration_seconds", time.Since(startTime).Seconds())
	}

	pv.logger.Info("Proof verification result", 
		"chain", chain, 
		"address", address, 
		"l2_wallet", l2WalletHex, 
		"is_valid", isValid,
		"duration", time.Since(startTime))

	return isValid, nil
}

// LinkWallet links an address to an L2 wallet after verifying the proof
func (pv *ProofVerifier) LinkWallet(ctx context.Context, chain ChainType, address, l2WalletHex, signature string) error {
	startTime := time.Now()
	pv.logger.Info("Linking wallet with proof verification", 
		"chain", chain, 
		"address", address, 
		"l2_wallet", l2WalletHex)

	// First verify the proof
	isValid, err := pv.VerifyProof(ctx, chain, address, l2WalletHex, signature)
	if err != nil {
		return err
	}

	if !isValid {
		return ErrVerificationFailed
	}

	// If the proof is valid, proceed with linking
	switch chain {
	case ChainEthereum:
		if pv.ethereumLinker == nil {
			return ErrNoLinker
		}
		err = pv.ethereumLinker.LinkWallet(ctx, address, l2WalletHex, signature)
	case ChainBitcoin:
		if pv.bitcoinLinker == nil {
			return ErrNoLinker
		}
		err = pv.bitcoinLinker.LinkWallet(ctx, address, l2WalletHex, signature)
	default:
		return ErrUnsupportedChain
	}

	if err != nil {
		pv.logger.Error("Wallet linking failed", 
			"error", err, 
			"chain", chain, 
			"address", address)
		return fmt.Errorf("failed to link wallet: %w", err)
	}

	// Record metrics
	if pv.metrics != nil {
		pv.metrics.IncCounter("wallets_linked_total", "chain", string(chain))
		pv.metrics.ObserveHistogram("wallet_linking_duration_seconds", time.Since(startTime).Seconds())
	}

	pv.logger.Info("Wallet linked successfully", 
		"chain", chain, 
		"address", address, 
		"l2_wallet", l2WalletHex,
		"duration", time.Since(startTime))

	return nil
}

// IsAddressLinked checks if an address is linked to an L2 wallet
func (pv *ProofVerifier) IsAddressLinked(ctx context.Context, chain ChainType, address string) (bool, error) {
	pv.logger.Debug("Checking if address is linked", "chain", chain, "address", address)

	switch chain {
	case ChainEthereum:
		if pv.ethereumLinker == nil {
			return false, ErrNoLinker
		}
		return pv.ethereumLinker.IsAddressLinked(ctx, address)
	case ChainBitcoin:
		// For Bitcoin, we don't have a direct IsAddressLinked method
		// We'll rely on the AddressMapping to check this
		return false, nil
	default:
		return false, ErrUnsupportedChain
	}
}

// IsL2WalletLinked checks if an L2 wallet is linked to an address on the specified chain
func (pv *ProofVerifier) IsL2WalletLinked(ctx context.Context, chain ChainType, l2WalletHex string) (bool, error) {
	pv.logger.Debug("Checking if L2 wallet is linked", "chain", chain, "l2_wallet", l2WalletHex)

	switch chain {
	case ChainEthereum:
		if pv.ethereumLinker == nil {
			return false, ErrNoLinker
		}
		return pv.ethereumLinker.IsL2WalletLinked(ctx, l2WalletHex)
	case ChainBitcoin:
		// For Bitcoin, we don't have a direct IsL2WalletLinked method
		// We'll rely on the AddressMapping to check this
		return false, nil
	default:
		return false, ErrUnsupportedChain
	}
}

// GetLinkedL2Wallet gets the L2 wallet linked to an address
func (pv *ProofVerifier) GetLinkedL2Wallet(ctx context.Context, chain ChainType, address string) (string, error) {
	pv.logger.Debug("Getting linked L2 wallet", "chain", chain, "address", address)

	switch chain {
	case ChainEthereum:
		if pv.ethereumLinker == nil {
			return "", ErrNoLinker
		}
		return pv.ethereumLinker.GetLinkedL2Wallet(ctx, address)
	case ChainBitcoin:
		// For Bitcoin, we don't have a direct GetLinkedL2Wallet method
		// We'll rely on the AddressMapping to retrieve this
		return "", nil
	default:
		return "", ErrUnsupportedChain
	}
}

// GetLinkedL1Address gets the address linked to an L2 wallet
func (pv *ProofVerifier) GetLinkedL1Address(ctx context.Context, chain ChainType, l2WalletHex string) (string, error) {
	pv.logger.Debug("Getting linked L1 address", "chain", chain, "l2_wallet", l2WalletHex)

	switch chain {
	case ChainEthereum:
		if pv.ethereumLinker == nil {
			return "", ErrNoLinker
		}
		return pv.ethereumLinker.GetLinkedEthereumAddress(ctx, l2WalletHex)
	case ChainBitcoin:
		// For Bitcoin, we don't have a direct GetLinkedAddress method
		// We'll rely on the AddressMapping to retrieve this
		return "", nil
	default:
		return "", ErrUnsupportedChain
	}
}

// VerifyCreditScore verifies the credit score of an L2 wallet
func (pv *ProofVerifier) VerifyCreditScore(ctx context.Context, chain ChainType, l2WalletHex string, threshold float64) (bool, error) {
	pv.logger.Info("Verifying credit score", "chain", chain, "l2_wallet", l2WalletHex, "threshold", threshold)

	switch chain {
	case ChainEthereum:
		if pv.ethereumLinker == nil {
			return false, ErrNoLinker
		}
		return pv.ethereumLinker.VerifyCreditScore(ctx, l2WalletHex, threshold)
	case ChainBitcoin:
		// Bitcoin doesn't support credit score verification directly
		// We might want to implement this differently or return an error
		return false, errors.New("credit score verification not supported for Bitcoin")
	default:
		return false, ErrUnsupportedChain
	}
}

// ProvideTransactionFeedback provides feedback after a transaction
func (pv *ProofVerifier) ProvideTransactionFeedback(ctx context.Context, chain ChainType, l2WalletHex string, isTrusted bool) error {
	pv.logger.Info("Providing transaction feedback", 
		"chain", chain, 
		"l2_wallet", l2WalletHex, 
		"is_trusted", isTrusted)

	switch chain {
	case ChainEthereum:
		if pv.ethereumLinker == nil {
			return ErrNoLinker
		}
		return pv.ethereumLinker.ProvideTransactionFeedback(ctx, l2WalletHex, isTrusted)
	case ChainBitcoin:
		// Bitcoin doesn't support transaction feedback directly
		// We might want to implement this differently or return an error
		return errors.New("transaction feedback not supported for Bitcoin")
	default:
		return ErrUnsupportedChain
	}
}

// DetectChainType tries to detect the chain type from an address format
func DetectChainType(address string) (ChainType, error) {
	// Ethereum addresses are 0x followed by 40 hex characters
	if strings.HasPrefix(address, "0x") && len(address) == 42 {
		return ChainEthereum, nil
	}

	// Bitcoin addresses have specific prefixes based on network and type
	// This is a simplified check - in practice, you'd want to validate the checksum, etc.
	if strings.HasPrefix(address, "1") || 
	   strings.HasPrefix(address, "3") || 
	   strings.HasPrefix(address, "bc1") {
		return ChainBitcoin, nil
	}

	// Testnet Bitcoin addresses
	if strings.HasPrefix(address, "m") || 
	   strings.HasPrefix(address, "n") || 
	   strings.HasPrefix(address, "2") || 
	   strings.HasPrefix(address, "tb1") {
		return ChainBitcoin, nil
	}

	return "", ErrInvalidAddress
}