package creditscore

import (
	"sync"
	"time"
    "math"
	"github.com/IWALINK/cosine/internal/storage"
	"github.com/IWALINK/cosine/internal/utils"
)

// ScalingFactorType represents different types of scaling factors
type ScalingFactorType string

// Scaling factor types
const (
	NegativeVoteScaling  ScalingFactorType = "K_neg"
	PositiveVoteScaling  ScalingFactorType = "K_pos"
	AssociationScaling   ScalingFactorType = "K_assoc"
	RehabilitationScaling ScalingFactorType = "K_rehab"
)

// DynamicScaling manages the dynamic scaling factors for credit score updates
type DynamicScaling struct {
	kalmanManager *KalmanManager
	globalState   *storage.GlobalStateObject
	config        *utils.ConfigManager
	logger        *utils.Logger
	metrics       *utils.MetricsManager
	mu            sync.RWMutex
}

// NewDynamicScaling creates a new dynamic scaling manager
func NewDynamicScaling(kalmanManager *KalmanManager, globalState *storage.GlobalStateObject, 
	config *utils.ConfigManager, logger *utils.Logger, metrics *utils.MetricsManager) *DynamicScaling {
	
	ds := &DynamicScaling{
		kalmanManager: kalmanManager,
		globalState:   globalState,
		config:        config,
		logger:        logger.WithComponent("DynamicScaling"),
		metrics:       metrics,
	}
	
	return ds
}

// ScalingInfo contains the current state of a scaling factor
type ScalingInfo struct {
	Type          ScalingFactorType
	Value         float64
	LastUpdated   time.Time
	UpdateCount   int
	AvgObserved   float64
	AvgExpected   float64
}

// GetScalingFactor returns the current value of a scaling factor
func (ds *DynamicScaling) GetScalingFactor(sfType ScalingFactorType) float64 {
	// Try to get from global state first if available
	if ds.globalState != nil {
		value := ds.globalState.GetScalingFactor(string(sfType))
		if value > 0 {
			return value
		}
	}
	
	// Fallback to Kalman manager
	return ds.kalmanManager.GetScalingFactor(string(sfType))
}

// UpdateScalingFactor updates a scaling factor based on observed vs expected impact
func (ds *DynamicScaling) UpdateScalingFactor(sfType ScalingFactorType, observedImpact, expectedImpact float64) float64 {
	// Get process noise (Q) and measurement noise (R) from config
	q := 0.01
	r := 1.0
	
	if ds.config != nil {
		q = ds.config.GetFloat64("creditscore.scaling.q")
		r = ds.config.GetFloat64("creditscore.scaling.r")
	}
	
	var newValue float64
	
	// Update in global state if available
	if ds.globalState != nil {
		newValue = ds.globalState.UpdateScalingFactor(string(sfType), observedImpact, expectedImpact, q, r)
	} else {
		// Fallback to Kalman manager
		newValue = ds.kalmanManager.UpdateScalingFactor(string(sfType), observedImpact, expectedImpact)
	}
	
	// Update metrics if available
	if ds.metrics != nil {
		ds.metrics.SetGauge("scaling_factor", newValue, string(sfType))
	}
	
	ds.logger.Debug("Updated scaling factor", 
		"type", sfType, 
		"new_value", newValue, 
		"observed_impact", observedImpact, 
		"expected_impact", expectedImpact)
	
	return newValue
}

// GetAllScalingFactors returns all current scaling factors
func (ds *DynamicScaling) GetAllScalingFactors() map[ScalingFactorType]float64 {
	result := make(map[ScalingFactorType]float64)
	
	// Add all scaling factor types
	result[NegativeVoteScaling] = ds.GetScalingFactor(NegativeVoteScaling)
	result[PositiveVoteScaling] = ds.GetScalingFactor(PositiveVoteScaling)
	result[AssociationScaling] = ds.GetScalingFactor(AssociationScaling)
	result[RehabilitationScaling] = ds.GetScalingFactor(RehabilitationScaling)
	
	return result
}

// CalculateVoteShift uses the appropriate scaling factor to calculate a vote-based credit score shift
func (ds *DynamicScaling) CalculateVoteShift(deltaV, sigmaV float64, gamma float64) float64 {
	// Determine if this is a positive or negative vote
	var scalingType ScalingFactorType
	var sign float64
	
	if deltaV < 0 {
		scalingType = NegativeVoteScaling
		sign = -1.0
	} else {
		scalingType = PositiveVoteScaling
		sign = 1.0
	}
	
	// Get the appropriate scaling factor
	kFactor := ds.GetScalingFactor(scalingType)
	
	// Calculate standardized deviation
	if sigmaV == 0 {
		sigmaV = 1.0 // Avoid division by zero
	}
	standardizedDev := math.Abs(deltaV) / sigmaV
	
	// Calculate shift magnitude
	magnitude := kFactor * math.Pow(standardizedDev, gamma)
	
	return sign * magnitude
}

// CalculateAssociationShift calculates a credit score shift based on association risk
func (ds *DynamicScaling) CalculateAssociationShift(riskScore, meanRisk, stdDevRisk float64, gamma float64) float64 {
	kFactor := ds.GetScalingFactor(AssociationScaling)
	
	// Only apply penalty if risk exceeds mean
	if riskScore <= meanRisk {
		return 0
	}
	
	// Avoid division by zero
	if stdDevRisk == 0 {
		stdDevRisk = 1.0
	}
	
	// Calculate standardized deviation
	standardizedDev := (riskScore - meanRisk) / stdDevRisk
	
	// Calculate negative shift
	return -kFactor * math.Pow(standardizedDev, gamma)
}

// CalculateRehabShift calculates a rehabilitation credit score shift
func (ds *DynamicScaling) CalculateRehabShift(rehabVote, meanRehab, stdDevRehab float64, gamma float64) float64 {
	kFactor := ds.GetScalingFactor(RehabilitationScaling)
	
	// Only apply rehab if vote exceeds mean
	if rehabVote <= meanRehab {
		return 0
	}
	
	// Avoid division by zero
	if stdDevRehab == 0 {
		stdDevRehab = 1.0
	}
	
	// Calculate standardized deviation
	standardizedDev := (rehabVote - meanRehab) / stdDevRehab
	
	// Calculate positive shift
	return kFactor * math.Pow(standardizedDev, gamma)
}

// UpdateScalingFactorFromShift updates a scaling factor based on actual credit score shift
func (ds *DynamicScaling) UpdateScalingFactorFromShift(sfType ScalingFactorType, actualShift, expectedShift float64) float64 {
	// Use absolute values since we're concerned with magnitude, not direction
	observedImpact := math.Abs(actualShift)
	expectedImpact := math.Abs(expectedShift)
	
	return ds.UpdateScalingFactor(sfType, observedImpact, expectedImpact)
}