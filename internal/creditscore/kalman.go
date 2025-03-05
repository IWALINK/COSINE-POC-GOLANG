package creditscore

import (
	"sync"
	"github.com/IWALINK/cosine/internal/utils"
)

// KalmanParams contains parameters for Kalman filter initialization
type KalmanParams struct {
	InitialValue     float64
	InitialP         float64
	ProcessNoise     float64 // Q
	MeasurementNoise float64 // R
}

// DefaultKalmanParams returns default parameters for a Kalman filter
func DefaultKalmanParams() KalmanParams {
	return KalmanParams{
		InitialValue:     1.0,
		InitialP:         1.0,
		ProcessNoise:     0.01,
		MeasurementNoise: 1.0,
	}
}

// ThresholdKalmanParams returns default parameters for threshold Kalman filters
func ThresholdKalmanParams() (tauParams, kParams KalmanParams) {
	tauParams = KalmanParams{
		InitialValue:     2.5, // Initial tau threshold (mean-based)
		InitialP:         1.0,
		ProcessNoise:     0.01,
		MeasurementNoise: 1.0,
	}
	
	kParams = KalmanParams{
		InitialValue:     3.0, // Initial k threshold (median-based)
		InitialP:         1.0,
		ProcessNoise:     0.01,
		MeasurementNoise: 1.0,
	}
	
	return
}

// KalmanManager manages multiple Kalman filters for different parameters
type KalmanManager struct {
	mu             sync.RWMutex
	thresholdTau   *utils.KalmanFilter         // For mean-based threshold
	thresholdK     *utils.KalmanFilter         // For median-based threshold
	scalingFactors map[string]*utils.DynamicScalingKalmanFilter
	logger         *utils.Logger
}

// NewKalmanManager creates a new Kalman filter manager
func NewKalmanManager(logger *utils.Logger, config *utils.ConfigManager) *KalmanManager {
	tauParams, kParams := ThresholdKalmanParams()
	
	// Override from config if available
	if config != nil {
		tauParams.InitialValue = config.GetFloat64("creditscore.kalman.tau_initial")
		kParams.InitialValue = config.GetFloat64("creditscore.kalman.k_initial")
		tauParams.InitialP = config.GetFloat64("creditscore.kalman.p_initial")
		kParams.InitialP = config.GetFloat64("creditscore.kalman.p_initial")
		tauParams.ProcessNoise = config.GetFloat64("creditscore.kalman.q_tau")
		kParams.ProcessNoise = config.GetFloat64("creditscore.kalman.q_k")
		tauParams.MeasurementNoise = config.GetFloat64("creditscore.kalman.r_tau")
		kParams.MeasurementNoise = config.GetFloat64("creditscore.kalman.r_k")
	}
	
	km := &KalmanManager{
		thresholdTau:   utils.NewKalmanFilter(tauParams.InitialValue, tauParams.InitialP, tauParams.ProcessNoise, tauParams.MeasurementNoise),
		thresholdK:     utils.NewKalmanFilter(kParams.InitialValue, kParams.InitialP, kParams.ProcessNoise, kParams.MeasurementNoise),
		scalingFactors: make(map[string]*utils.DynamicScalingKalmanFilter),
		logger:         logger.WithComponent("KalmanManager"),
	}
	
	// Initialize scaling factor Kalman filters
	if config != nil {
		k_neg, k_pos, k_assoc, k_rehab, q, r := config.GetScalingFactorParams()
		km.scalingFactors["K_neg"] = utils.NewDynamicScalingKalmanFilter(k_neg, tauParams.InitialP, q, r)
		km.scalingFactors["K_pos"] = utils.NewDynamicScalingKalmanFilter(k_pos, tauParams.InitialP, q, r)
		km.scalingFactors["K_assoc"] = utils.NewDynamicScalingKalmanFilter(k_assoc, tauParams.InitialP, q, r)
		km.scalingFactors["K_rehab"] = utils.NewDynamicScalingKalmanFilter(k_rehab, tauParams.InitialP, q, r)
	} else {
		// Use defaults
		km.scalingFactors["K_neg"] = utils.NewDynamicScalingKalmanFilter(1.0, 1.0, 0.01, 1.0)
		km.scalingFactors["K_pos"] = utils.NewDynamicScalingKalmanFilter(1.0, 1.0, 0.01, 1.0)
		km.scalingFactors["K_assoc"] = utils.NewDynamicScalingKalmanFilter(1.0, 1.0, 0.01, 1.0)
		km.scalingFactors["K_rehab"] = utils.NewDynamicScalingKalmanFilter(1.0, 1.0, 0.01, 1.0)
	}
	
	km.logger.Info("Initialized Kalman filter manager", 
		"tau", tauParams.InitialValue, 
		"k", kParams.InitialValue,
		"scaling_factors", len(km.scalingFactors))
	
	return km
}

// UpdateThresholdTau updates the tau threshold with a new measurement
func (km *KalmanManager) UpdateThresholdTau(measurement float64) float64 {
	km.mu.Lock()
	defer km.mu.Unlock()
	
	updated := km.thresholdTau.Update(measurement)
	km.logger.Debug("Updated threshold tau", "old", km.thresholdTau.X, "new", updated, "measurement", measurement)
	return updated
}

// UpdateThresholdK updates the k threshold with a new measurement
func (km *KalmanManager) UpdateThresholdK(measurement float64) float64 {
	km.mu.Lock()
	defer km.mu.Unlock()
	
	updated := km.thresholdK.Update(measurement)
	km.logger.Debug("Updated threshold k", "old", km.thresholdK.X, "new", updated, "measurement", measurement)
	return updated
}

// UpdateScalingFactor updates a scaling factor based on observed vs expected impact
func (km *KalmanManager) UpdateScalingFactor(name string, observedImpact, expectedImpact float64) float64 {
	km.mu.Lock()
	defer km.mu.Unlock()
	
	kf, exists := km.scalingFactors[name]
	if !exists {
		km.logger.Warn("Attempted to update non-existent scaling factor", "name", name)
		// Create it with defaults
		kf = utils.NewDynamicScalingKalmanFilter(1.0, 1.0, 0.01, 1.0)
		km.scalingFactors[name] = kf
	}
	
	oldValue := kf.X
	updated := kf.UpdateWithImpact(observedImpact, expectedImpact)
	
	km.logger.Debug("Updated scaling factor", 
		"name", name, 
		"old", oldValue, 
		"new", updated, 
		"observed", observedImpact, 
		"expected", expectedImpact)
	
	return updated
}

// GetThresholdTau returns the current tau threshold value
func (km *KalmanManager) GetThresholdTau() float64 {
	km.mu.RLock()
	defer km.mu.RUnlock()
	return km.thresholdTau.X
}

// GetThresholdK returns the current k threshold value
func (km *KalmanManager) GetThresholdK() float64 {
	km.mu.RLock()
	defer km.mu.RUnlock()
	return km.thresholdK.X
}

// GetScalingFactor returns a scaling factor value
func (km *KalmanManager) GetScalingFactor(name string) float64 {
	km.mu.RLock()
	defer km.mu.RUnlock()
	
	kf, exists := km.scalingFactors[name]
	if !exists {
		return 1.0 // Default
	}
	return kf.X
}

// GetAllScalingFactors returns all scaling factors
func (km *KalmanManager) GetAllScalingFactors() map[string]float64 {
	km.mu.RLock()
	defer km.mu.RUnlock()
	
	result := make(map[string]float64)
	for name, kf := range km.scalingFactors {
		result[name] = kf.X
	}
	return result
}

// ThresholdParams returns the current threshold parameters for outlier filtering
func (km *KalmanManager) ThresholdParams() (tau, k float64) {
	km.mu.RLock()
	defer km.mu.RUnlock()
	return km.thresholdTau.X, km.thresholdK.X
}