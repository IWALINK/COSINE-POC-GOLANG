// Package consensus provides consensus and validator selection mechanisms for COSINE.
package consensus

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
	"math"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

// VRFOutput represents the output of a VRF computation along with a proof
type VRFOutput struct {
	// The raw output bytes from the VRF
	Output []byte
	
	// The proof that can be used to verify the VRF output
	Proof []byte
	
	// Value is a normalized float64 in [0,1) derived from the output
	Value float64
	
	// Timestamp when the VRF was computed
	Timestamp time.Time
}

// VRFManager handles Verifiable Random Function operations
type VRFManager struct {
	logger     *utils.Logger
	privateKey crypto.PrivKey
	peerId     peer.ID
}

// NewVRFManager creates a new VRF manager
func NewVRFManager(logger *utils.Logger, privateKey crypto.PrivKey) (*VRFManager, error) {
	if privateKey == nil {
		return nil, fmt.Errorf("private key is required")
	}
	
	// Get peer ID from private key
	peerId, err := peer.IDFromPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get peer ID from private key: %w", err)
	}
	
	return &VRFManager{
		logger:     logger.WithComponent("VRFManager"),
		privateKey: privateKey,
		peerId:     peerId,
	}, nil
}

// ComputeVRF computes a VRF output for a given seed
func (vm *VRFManager) ComputeVRF(seed []byte) (*VRFOutput, error) {
	// In a real VRF implementation, we would use a specialized VRF library
	// For COSINE MVP, we'll use a simpler approach based on digital signatures
	
	// Sign the seed with our private key
	signature, err := vm.privateKey.Sign(seed)
	if err != nil {
		return nil, fmt.Errorf("failed to sign seed: %w", err)
	}
	
	// Hash the signature to get VRF output
	hash := sha256.Sum256(signature)
	
	// Convert first 8 bytes to a uint64 for normalization
	value := binary.BigEndian.Uint64(hash[:8])
	
	// Normalize to [0,1)
	normalized := float64(value) / float64(1<<64)
	
	vm.logger.Debug("Computed VRF",
		"peer_id", vm.peerId.String(),
		"seed_hash", fmt.Sprintf("%x", sha256.Sum256(seed)),
		"value", normalized)
	
	return &VRFOutput{
		Output:    hash[:],
		Proof:     signature,
		Value:     normalized,
		Timestamp: time.Now(),
	}, nil
}

// VerifyVRF verifies a VRF output against a public key
func VerifyVRF(pubKey crypto.PubKey, seed []byte, vrfOutput *VRFOutput) (bool, error) {
	// Verify the signature (proof)
	valid, err := pubKey.Verify(seed, vrfOutput.Proof)
	if err != nil || !valid {
		return false, fmt.Errorf("invalid VRF proof: %v", err)
	}
	
	// Hash the signature to get VRF output
	hash := sha256.Sum256(vrfOutput.Proof)
	
	// Check if the output matches
	for i := range hash {
		if i >= len(vrfOutput.Output) || hash[i] != vrfOutput.Output[i] {
			return false, fmt.Errorf("VRF output does not match proof")
		}
	}
	
	// Convert first 8 bytes to a uint64 for normalization
	value := binary.BigEndian.Uint64(hash[:8])
	
	// Normalize to [0,1) and check value
	normalized := float64(value) / float64(1<<64)
	if normalized != vrfOutput.Value {
		return false, fmt.Errorf("VRF normalized value does not match: %f vs %f", normalized, vrfOutput.Value)
	}
	
	return true, nil
}

// ComputeScore computes a selection score based on VRF output and validator weight
func ComputeScore(vrfValue float64, stake float64, performance float64, alpha float64, beta float64) float64 {
	// Calculate the weight according to formula: W_v = S_v^α × (1 + β*P_v)
	weight := math.Pow(stake, alpha) * (1 + beta*performance)
	
	// Safety check to avoid division by zero
	if weight <= 0 {
		return float64(^uint64(0)) // Maximum float64 value
	}
	
	// Calculate score according to formula: score_v = y_v / W_v
	return vrfValue / weight
}

// GenerateBlockSeed generates a seed for VRF based on previous block hash and timestamp
func GenerateBlockSeed(prevBlockHash []byte, timestamp time.Time) []byte {
	hasher := sha256.New()
	hasher.Write(prevBlockHash)
	
	// Add timestamp to make the seed unique even with same block hash
	timeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBytes, uint64(timestamp.UnixNano()))
	hasher.Write(timeBytes)
	
	return hasher.Sum(nil)
}

// ScoreToBigInt converts a score to a big.Int for precise comparison
func ScoreToBigInt(score float64) *big.Int {
	// Multiply by a large factor to preserve precision
	scaledScore := score * 1e18
	
	// Convert to big.Int
	bigScore := new(big.Int)
	bigScore.SetString(fmt.Sprintf("%.0f", scaledScore), 10)
	
	return bigScore
}