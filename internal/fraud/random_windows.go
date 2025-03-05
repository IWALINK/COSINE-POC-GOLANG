package fraud

import (
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
)

// RandomWindows manages the generation of random time windows and hop limits
// for association risk analysis, making it difficult for attackers to predict
// when and how far the system will look for suspicious transactions.
type RandomWindows struct {
	minTimeWindow int   // Minimum time window in days
	maxTimeWindow int   // Maximum time window in days
	minHopLimit   int   // Minimum hop limit
	maxHopLimit   int   // Maximum hop limit
	config        *utils.ConfigManager
	logger        *utils.Logger
	metrics       *utils.MetricsManager
	mu            sync.RWMutex
}

// RandomWindowsParams contains parameters for the random window generation
type RandomWindowsParams struct {
	TimeLimit time.Duration // Random time window duration
	HopLimit  int           // Random hop depth limit
	Seed      int64         // Optional seed for reproducibility (testing)
}

// NewRandomWindows creates a new random windows generator with default or config-provided ranges
func NewRandomWindows(config *utils.ConfigManager, logger *utils.Logger, metrics *utils.MetricsManager) *RandomWindows {
	// Default values from whitepaper
	minTimeWindow := 30  // 30 days minimum time window
	maxTimeWindow := 730 // 730 days (2 years) maximum time window
	minHopLimit := 1     // Minimum hop limit
	maxHopLimit := 15    // Maximum hop limit of 15 (from whitepaper)

	// Override with configuration if available
	if config != nil {
		minTimeWindow = config.GetInt("fraud.timeWindow.min")
		maxTimeWindow = config.GetInt("fraud.timeWindow.max")
		minHopLimit = config.GetInt("fraud.hopLimit.min")
		maxHopLimit = config.GetInt("fraud.hopLimit.max")
	}

	rw := &RandomWindows{
		minTimeWindow: minTimeWindow,
		maxTimeWindow: maxTimeWindow,
		minHopLimit:   minHopLimit,
		maxHopLimit:   maxHopLimit,
		config:        config,
		logger:        logger.WithComponent("RandomWindows"),
		metrics:       metrics,
	}

	rw.logger.Info("Initialized random windows generator",
		"minTimeWindow", rw.minTimeWindow,
		"maxTimeWindow", rw.maxTimeWindow,
		"minHopLimit", rw.minHopLimit,
		"maxHopLimit", rw.maxHopLimit)

	return rw
}

// GenerateRandomWindows generates a new set of random time window and hop limit
// This is used when starting a new association risk analysis
func (rw *RandomWindows) GenerateRandomWindows() RandomWindowsParams {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	// Generate random time window in days, then convert to Duration
	timeLimitDays := utils.RandomIntInRange(rw.minTimeWindow, rw.maxTimeWindow)
	timeLimit := time.Duration(timeLimitDays) * 24 * time.Hour

	// Generate random hop limit
	hopLimit := utils.RandomIntInRange(rw.minHopLimit, rw.maxHopLimit)

	// Update metrics if available
	if rw.metrics != nil {
		rw.metrics.ObserveHistogram("random_time_window_days", float64(timeLimitDays))
		rw.metrics.ObserveHistogram("random_hop_limit", float64(hopLimit))
	}

	rw.logger.Debug("Generated random windows",
		"timeLimitDays", timeLimitDays,
		"hopLimit", hopLimit)

	return RandomWindowsParams{
		TimeLimit: timeLimit,
		HopLimit:  hopLimit,
		Seed:      time.Now().UnixNano(), // Using current time as seed
	}
}

// GenerateForWallet generates a deterministic set of random windows for a specific wallet
// This ensures the same wallet always gets the same random parameters for a given trigger event
func (rw *RandomWindows) GenerateForWallet(walletAddress string, triggerTimestamp time.Time) RandomWindowsParams {
	// Create a deterministic seed based on wallet address and trigger time
	seed := utils.DeterministicSeed(walletAddress, triggerTimestamp.Unix())

	rw.mu.RLock()
	minTime := rw.minTimeWindow
	maxTime := rw.maxTimeWindow
	minHop := rw.minHopLimit
	maxHop := rw.maxHopLimit
	rw.mu.RUnlock()

	// Use the deterministic seed for random generation
	timeLimitDays := utils.DeterministicRandomInt(seed, minTime, maxTime)
	timeLimit := time.Duration(timeLimitDays) * 24 * time.Hour

	// Use different seed component for hop limit to avoid correlation
	hopLimit := utils.DeterministicRandomInt(seed+42, minHop, maxHop)

	rw.logger.Debug("Generated deterministic windows for wallet",
		"wallet", walletAddress,
		"triggerTime", triggerTimestamp,
		"timeLimitDays", timeLimitDays,
		"hopLimit", hopLimit)

	return RandomWindowsParams{
		TimeLimit: timeLimit,
		HopLimit:  hopLimit,
		Seed:      seed,
	}
}

// IsWithinTimeWindow checks if a transaction timestamp falls within the random time window
// starting from the trigger timestamp
func (rw *RandomWindows) IsWithinTimeWindow(txTimestamp, triggerTimestamp time.Time, params RandomWindowsParams) bool {
	// Check if transaction is after trigger time and before (trigger time + time window)
	withinWindow := txTimestamp.After(triggerTimestamp) && 
	                txTimestamp.Before(triggerTimestamp.Add(params.TimeLimit))
	
	return withinWindow
}

// UpdateWindowParameters updates the min/max ranges for time windows and hop limits
func (rw *RandomWindows) UpdateWindowParameters(minTimeWindow, maxTimeWindow, minHopLimit, maxHopLimit int) {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	// Validate parameters
	if minTimeWindow > 0 && maxTimeWindow >= minTimeWindow {
		rw.minTimeWindow = minTimeWindow
		rw.maxTimeWindow = maxTimeWindow
	}

	if minHopLimit > 0 && maxHopLimit >= minHopLimit {
		rw.minHopLimit = minHopLimit
		rw.maxHopLimit = maxHopLimit
	}

	rw.logger.Info("Updated random window parameters",
		"minTimeWindow", rw.minTimeWindow,
		"maxTimeWindow", rw.maxTimeWindow,
		"minHopLimit", rw.minHopLimit,
		"maxHopLimit", rw.maxHopLimit)
}

// GetWindowParameters returns the current min/max parameters
func (rw *RandomWindows) GetWindowParameters() (minTimeWindow, maxTimeWindow, minHopLimit, maxHopLimit int) {
	rw.mu.RLock()
	defer rw.mu.RUnlock()

	return rw.minTimeWindow, rw.maxTimeWindow, rw.minHopLimit, rw.maxHopLimit
}