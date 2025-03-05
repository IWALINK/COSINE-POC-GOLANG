package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ConfigManager handles loading, accessing, and watching configuration values.
// It supports dynamic reloading of configuration values and type-safe access.
type ConfigManager struct {
	v           *viper.Viper
	configFile  string
	environment string
	mu          sync.RWMutex
	callbacks   map[string][]func(interface{})
}

// DefaultConfigPaths returns default paths to search for config files
func DefaultConfigPaths() []string {
	return []string{
		".",
		"./configs",
		"../configs",
		"/etc/cosine",
	}
}

// NewConfigManager creates a new configuration manager
func NewConfigManager(configFile, environment string) (*ConfigManager, error) {
	v := viper.New()
	cm := &ConfigManager{
		v:           v,
		configFile:  configFile,
		environment: environment,
		callbacks:   make(map[string][]func(interface{})),
	}

	// Set up viper configuration
	v.SetConfigType("yaml")
	v.SetEnvPrefix("COSINE")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Default configuration values
	cm.setDefaults()

	// Find and load the configuration file
	found := false
	for _, path := range DefaultConfigPaths() {
		fullPath := filepath.Join(path, configFile)
		if _, err := os.Stat(fullPath); err == nil {
			v.SetConfigFile(fullPath)
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("configuration file %s not found in any of the default paths", configFile)
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	// Enable hot reloading of configuration
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		cm.mu.RLock()
		defer cm.mu.RUnlock()

		// Notify all registered callbacks of configuration changes
		for key, funcs := range cm.callbacks {
			value := v.Get(key)
			for _, f := range funcs {
				f(value)
			}
		}
	})

	return cm, nil
}

// setDefaults sets default configuration values
func (cm *ConfigManager) setDefaults() {
	// Network and node configuration
	cm.v.SetDefault("node.id", "")
	cm.v.SetDefault("node.environment", "development")
	cm.v.SetDefault("node.datadir", "./data")

	// P2P network settings
	cm.v.SetDefault("p2p.port", 9000)
	cm.v.SetDefault("p2p.bootstrapNodes", []string{})
	cm.v.SetDefault("p2p.maxPeers", 50)
	cm.v.SetDefault("p2p.pingInterval", "5m") // 5 minutes ping interval

	// API server settings
	cm.v.SetDefault("api.enabled", true)
	cm.v.SetDefault("api.port", 8080)
	cm.v.SetDefault("api.cors.enabled", false)
	cm.v.SetDefault("api.cors.allowedOrigins", []string{"*"})

	// VRF consensus parameters
	cm.v.SetDefault("consensus.alpha", 0.5)
	cm.v.SetDefault("consensus.beta", 0.2)
	cm.v.SetDefault("consensus.gamma", 0.1)
	cm.v.SetDefault("consensus.maxValidators", 100)

	// Kalman filter parameters for hybrid outlier detection
	cm.v.SetDefault("creditscore.kalman.tau_initial", 2.5)
	cm.v.SetDefault("creditscore.kalman.k_initial", 3.0)
	cm.v.SetDefault("creditscore.kalman.p_initial", 1.0)
	cm.v.SetDefault("creditscore.kalman.q_tau", 0.01)
	cm.v.SetDefault("creditscore.kalman.q_k", 0.01)
	cm.v.SetDefault("creditscore.kalman.r_tau", 1.0)
	cm.v.SetDefault("creditscore.kalman.r_k", 1.0)

	// Dynamic scaling factors
	cm.v.SetDefault("creditscore.scaling.k_neg_initial", 1.0)
	cm.v.SetDefault("creditscore.scaling.k_pos_initial", 1.0)
	cm.v.SetDefault("creditscore.scaling.k_assoc_initial", 1.0)
	cm.v.SetDefault("creditscore.scaling.k_rehab_initial", 1.0)
	cm.v.SetDefault("creditscore.scaling.q", 0.01) // Process noise variance
	cm.v.SetDefault("creditscore.scaling.r", 1.0)  // Measurement noise variance

	// Fraud detection parameters
	cm.v.SetDefault("fraud.timeWindow.min", "30d")  // Minimum time window (30 days)
	cm.v.SetDefault("fraud.timeWindow.max", "730d") // Maximum time window (730 days)
	cm.v.SetDefault("fraud.hopLimit.min", 1)        // Minimum hop limit
	cm.v.SetDefault("fraud.hopLimit.max", 15)       // Maximum hop limit
	cm.v.SetDefault("fraud.association.beta", 0.5)  // Clustering weight
	cm.v.SetDefault("fraud.association.delta", 0.01) // Time decay factor

	// Voting and governance
	cm.v.SetDefault("voting.newWalletReputation", 0.1)
	cm.v.SetDefault("voting.alpha", 0.05) // Reputation increase rate
	cm.v.SetDefault("voting.beta", 0.1)   // Reputation decrease rate
	cm.v.SetDefault("voting.lambda", 2.5) // Std dev tolerance for vote impact

	// IPFS storage settings
	cm.v.SetDefault("storage.ipfs.enabled", true)
	cm.v.SetDefault("storage.ipfs.gateway", "https://ipfs.io/ipfs/")
	cm.v.SetDefault("storage.ipfs.api", "/ip4/127.0.0.1/tcp/5001")

	// Transaction fees
	cm.v.SetDefault("fees.verification.kappa", 1.2) // Fee multiplier
	cm.v.SetDefault("fees.token.baseFee", 0.002)   // Base fee for token transfers
}

// GetString returns a string configuration value
func (cm *ConfigManager) GetString(key string) string {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.v.GetString(key)
}

// GetInt returns an integer configuration value
func (cm *ConfigManager) GetInt(key string) int {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.v.GetInt(key)
}

// GetFloat64 returns a float64 configuration value
func (cm *ConfigManager) GetFloat64(key string) float64 {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.v.GetFloat64(key)
}

// GetBool returns a boolean configuration value
func (cm *ConfigManager) GetBool(key string) bool {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.v.GetBool(key)
}

// GetStringSlice returns a string slice configuration value
func (cm *ConfigManager) GetStringSlice(key string) []string {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.v.GetStringSlice(key)
}

// GetDuration returns a duration configuration value
func (cm *ConfigManager) GetDuration(key string) time.Duration {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.v.GetDuration(key)
}

// GetKalmanFilterParams returns parameters for a Kalman filter
func (cm *ConfigManager) GetKalmanFilterParams(prefix string) (tau, k, p, q_tau, q_k, r_tau, r_k float64) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	
	base := "creditscore.kalman."
	if prefix != "" {
		base = prefix + "."
	}
	
	tau = cm.v.GetFloat64(base + "tau_initial")
	k = cm.v.GetFloat64(base + "k_initial")
	p = cm.v.GetFloat64(base + "p_initial")
	q_tau = cm.v.GetFloat64(base + "q_tau")
	q_k = cm.v.GetFloat64(base + "q_k")
	r_tau = cm.v.GetFloat64(base + "r_tau")
	r_k = cm.v.GetFloat64(base + "r_k")
	
	return
}

// GetScalingFactorParams returns parameters for dynamic scaling factor updates
func (cm *ConfigManager) GetScalingFactorParams() (k_neg, k_pos, k_assoc, k_rehab, q, r float64) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	
	k_neg = cm.v.GetFloat64("creditscore.scaling.k_neg_initial")
	k_pos = cm.v.GetFloat64("creditscore.scaling.k_pos_initial")
	k_assoc = cm.v.GetFloat64("creditscore.scaling.k_assoc_initial")
	k_rehab = cm.v.GetFloat64("creditscore.scaling.k_rehab_initial")
	q = cm.v.GetFloat64("creditscore.scaling.q")
	r = cm.v.GetFloat64("creditscore.scaling.r")
	
	return
}

// OnConfigChange registers a callback function to be called when a specific configuration value changes
func (cm *ConfigManager) OnConfigChange(key string, callback func(interface{})) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	cm.callbacks[key] = append(cm.callbacks[key], callback)
}

// Save writes the current configuration to the file
func (cm *ConfigManager) Save() error {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	return cm.v.WriteConfig()
}

// Set sets a configuration value
func (cm *ConfigManager) Set(key string, value interface{}) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	cm.v.Set(key, value)
}

// LoadValidatorConfig loads validator-specific configuration
func (cm *ConfigManager) LoadValidatorConfig() map[string]interface{} {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	
	return cm.v.GetStringMap("validator")
}

// LoadAPIConfig loads API server configuration
func (cm *ConfigManager) LoadAPIConfig() map[string]interface{} {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	
	return cm.v.GetStringMap("api")
}

// GetConfig returns direct access to the Viper instance (use with caution)
func (cm *ConfigManager) GetConfig() *viper.Viper {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	
	return cm.v
}

// IsProduction returns true if the environment is production
func (cm *ConfigManager) IsProduction() bool {
	return cm.environment == "production"
}

// Global configuration instance
var (
	globalConfig *ConfigManager
	configOnce   sync.Once
)

// GetGlobalConfig returns the global configuration instance
func GetGlobalConfig() *ConfigManager {
	return globalConfig
}

// InitGlobalConfig initializes the global configuration instance
func InitGlobalConfig(configFile, environment string) error {
	var err error
	configOnce.Do(func() {
		globalConfig, err = NewConfigManager(configFile, environment)
	})
	return err
}