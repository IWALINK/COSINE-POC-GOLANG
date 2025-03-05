package creditscore

import (
	"math"
	"sync"
	"time"

	"github.com/IWALINK/cosine/internal/utils"
)

// ProposalEntry represents a validator proposal for a credit score
type ProposalEntry struct {
	ValidatorID string
	Score       float64
	Timestamp   time.Time
}

// OutlierFilterResult contains the aggregated score and outlier information
type OutlierFilterResult struct {
	AggregatedScore float64
	InlierCount     int
	OutlierCount    int
	Inliers         []ProposalEntry
	Outliers        []ProposalEntry
	Mean            float64
	Median          float64
	StdDev          float64
	MAD             float64
	TauThreshold    float64
	KThreshold      float64
}

// OutlierFilter implements hybrid outlier filtering with dynamic thresholds
type OutlierFilter struct {
	kalmanManager *KalmanManager
	config        *utils.ConfigManager
	logger        *utils.Logger
	metrics       *utils.MetricsManager
	mu            sync.RWMutex
}

// NewOutlierFilter creates a new outlier filter
func NewOutlierFilter(kalmanManager *KalmanManager, config *utils.ConfigManager, logger *utils.Logger, metrics *utils.MetricsManager) *OutlierFilter {
	return &OutlierFilter{
		kalmanManager: kalmanManager,
		config:        config,
		logger:        logger.WithComponent("OutlierFilter"),
		metrics:       metrics,
	}
}

// FilterOutliers applies hybrid outlier filtering to a set of validator proposals
func (of *OutlierFilter) FilterOutliers(proposals []ProposalEntry) OutlierFilterResult {
	if len(proposals) == 0 {
		return OutlierFilterResult{
			AggregatedScore: 0,
			InlierCount:     0,
			OutlierCount:    0,
			Inliers:         []ProposalEntry{},
			Outliers:        []ProposalEntry{},
		}
	}
	
	// Extract score values for statistical calculations
	values := make([]float64, len(proposals))
	for i, p := range proposals {
		values[i] = p.Score
	}
	
	// Step 1: Calculate mean and standard deviation
	mean := utils.Mean(values)
	stdDev := utils.StandardDeviation(values)
	
	// Step 2: Calculate median and MAD
	median := utils.Median(values)
	madNormalized := utils.NormalizedMAD(values)
	
	// Step 3: Get dynamic thresholds from Kalman filters
	tau, k := of.kalmanManager.ThresholdParams()
	
	// Step 4: Classify each proposal as inlier or outlier using the hybrid approach
	inliers := []ProposalEntry{}
	outliers := []ProposalEntry{}
	
	for _, proposal := range proposals {
		// Mean-based deviation
		meanDev := math.Abs(proposal.Score - mean)
		meanThreshold := tau * stdDev
		
		// Median-based deviation
		medianDev := math.Abs(proposal.Score - median)
		medianThreshold := k * madNormalized
		
		// A proposal is an inlier only if it passes both checks
		if meanDev <= meanThreshold && medianDev <= medianThreshold {
			inliers = append(inliers, proposal)
		} else {
			outliers = append(outliers, proposal)
		}
	}
	
	// Step 5: Compute aggregated score from inliers only
	var aggregatedScore float64
	if len(inliers) > 0 {
		// Simple mean of inlier scores
		inlierSum := 0.0
		for _, inlier := range inliers {
			inlierSum += inlier.Score
		}
		aggregatedScore = inlierSum / float64(len(inliers))
	} else {
		// If no inliers, use the median as a fallback
		aggregatedScore = median
		of.logger.Warn("No inliers found, using median as fallback", 
			"proposal_count", len(proposals),
			"median", median)
	}
	
	// Step 6: Update Kalman filters with observed max deviations
	// Find max normalized deviations for threshold updates
	maxMeanDev := 0.0
	maxMedianDev := 0.0
	
	for _, proposal := range proposals {
		meanDev := math.Abs(proposal.Score - mean) / stdDev
		medianDev := math.Abs(proposal.Score - median) / madNormalized
		
		if meanDev > maxMeanDev {
			maxMeanDev = meanDev
		}
		if medianDev > maxMedianDev {
			maxMedianDev = medianDev
		}
	}
	
	// Only update if we have enough proposals and non-zero standard deviation
	if len(proposals) >= 3 && stdDev > 0 && madNormalized > 0 {
		of.kalmanManager.UpdateThresholdTau(maxMeanDev)
		of.kalmanManager.UpdateThresholdK(maxMedianDev)
	}
	
	// Update metrics if available
	if of.metrics != nil {
		of.metrics.SetGauge("kalman_filter_tau", tau)
		of.metrics.SetGauge("kalman_filter_k", k)
		// Fix to increment counters correctly:
		of.metrics.IncCounter("inlier_proposals_total")
		of.metrics.IncCounter("outlier_proposals_total")
	}
	
	of.logger.Debug("Applied hybrid outlier filtering",
		"total_proposals", len(proposals),
		"inliers", len(inliers),
		"outliers", len(outliers),
		"mean", mean,
		"median", median,
		"stddev", stdDev,
		"mad", madNormalized,
		"tau", tau,
		"k", k,
		"aggregated_score", aggregatedScore)
	
	return OutlierFilterResult{
		AggregatedScore: aggregatedScore,
		InlierCount:     len(inliers),
		OutlierCount:    len(outliers),
		Inliers:         inliers,
		Outliers:        outliers,
		Mean:            mean,
		Median:          median,
		StdDev:          stdDev,
		MAD:             madNormalized,
		TauThreshold:    tau,
		KThreshold:      k,
	}
}

// GetIsInlier determines if a proposal would be classified as an inlier
func (of *OutlierFilter) GetIsInlier(proposal float64, mean, stdDev, median, mad float64) bool {
	tau, k := of.kalmanManager.ThresholdParams()
	
	// Mean-based check
	meanDev := math.Abs(proposal - mean)
	meanThreshold := tau * stdDev
	
	// Median-based check
	medianDev := math.Abs(proposal - median)
	medianThreshold := k * mad
	
	// Pass both checks
	return meanDev <= meanThreshold && medianDev <= medianThreshold
}

// GetThresholds returns the current threshold values
func (of *OutlierFilter) GetThresholds() (tau, k float64) {
	return of.kalmanManager.ThresholdParams()
}