package utils

import (
	"math"
	"math/rand"
	"sort"
	"time"
	"fmt"
    "crypto/sha256"
    "encoding/binary"
)

// Initialize random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Vector represents a mathematical vector
type Vector []float64

// NewVector creates a new vector with specified size
func NewVector(size int) Vector {
	return make(Vector, size)
}

// VectorFromValues creates a new vector from values
func VectorFromValues(values ...float64) Vector {
	v := make(Vector, len(values))
	copy(v, values)
	return v
}

// Norm returns the Euclidean norm (L2 norm) of the vector
func (v Vector) Norm() float64 {
	sumSquares := 0.0
	for _, val := range v {
		sumSquares += val * val
	}
	return math.Sqrt(sumSquares)
}

// Normalize returns a unit vector in the same direction
func (v Vector) Normalize() Vector {
	norm := v.Norm()
	if norm == 0 {
		return Vector(make([]float64, len(v)))
	}
	
	result := make(Vector, len(v))
	for i, val := range v {
		result[i] = val / norm
	}
	return result
}

// Dot returns the dot product of two vectors
func (v Vector) Dot(other Vector) float64 {
	if len(v) != len(other) {
		return 0
	}
	
	sum := 0.0
	for i, val := range v {
		sum += val * other[i]
	}
	return sum
}

// CosineSimilarity calculates the cosine similarity between two vectors
func CosineSimilarity(v1, v2 Vector) float64 {
	dot := v1.Dot(v2)
	normV1 := v1.Norm()
	normV2 := v2.Norm()
	
	if normV1 == 0 || normV2 == 0 {
		return 0
	}
	
	return dot / (normV1 * normV2)
}

// Mean calculates the arithmetic mean of a slice of float64 values
func Mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// WeightedMean calculates the weighted arithmetic mean
func WeightedMean(values, weights []float64) float64 {
	if len(values) == 0 || len(values) != len(weights) {
		return 0
	}
	
	sum := 0.0
	weightSum := 0.0
	for i, v := range values {
		sum += v * weights[i]
		weightSum += weights[i]
	}
	
	if weightSum == 0 {
		return 0
	}
	
	return sum / weightSum
}

// Variance calculates the variance of a slice of float64 values
func Variance(values []float64) float64 {
	if len(values) < 2 {
		return 0
	}
	
	mean := Mean(values)
	sumSquareDiffs := 0.0
	for _, v := range values {
		diff := v - mean
		sumSquareDiffs += diff * diff
	}
	return sumSquareDiffs / float64(len(values))
}

// StandardDeviation calculates the standard deviation
func StandardDeviation(values []float64) float64 {
	return math.Sqrt(Variance(values))
}

// Median calculates the median value
func Median(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	
	// Make a copy to avoid modifying the original slice
	valuesCopy := make([]float64, len(values))
	copy(valuesCopy, values)
	sort.Float64s(valuesCopy)
	
	n := len(valuesCopy)
	if n%2 == 1 {
		return valuesCopy[n/2]
	}
	return (valuesCopy[n/2-1] + valuesCopy[n/2]) / 2
}

// MedianAbsoluteDeviation calculates the median absolute deviation
func MedianAbsoluteDeviation(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	
	med := Median(values)
	
	// Calculate absolute deviations from the median
	deviations := make([]float64, len(values))
	for i, v := range values {
		deviations[i] = math.Abs(v - med)
	}
	
	// Return the median of the absolute deviations
	return Median(deviations)
}

// NormalizedMAD calculates the normalized median absolute deviation
func NormalizedMAD(values []float64) float64 {
	// 1.4826 is the scaling factor for normal distribution
	return 1.4826 * MedianAbsoluteDeviation(values)
}

// ExponentialMovingAverage updates an EMA with a new value
func ExponentialMovingAverage(currentEMA, newValue float64, alpha float64) float64 {
	return alpha*newValue + (1-alpha)*currentEMA
}

// KalmanFilter represents a Kalman filter for scalar values
type KalmanFilter struct {
	X float64 // State estimate
	P float64 // Estimation error covariance
	Q float64 // Process noise covariance
	R float64 // Measurement noise covariance
}

// NewKalmanFilter creates a new Kalman filter with initial values
func NewKalmanFilter(initialX, initialP, q, r float64) *KalmanFilter {
	return &KalmanFilter{
		X: initialX,
		P: initialP,
		Q: q,
		R: r,
	}
}

// Update updates the Kalman filter with a new measurement
func (kf *KalmanFilter) Update(measurement float64) float64 {
	// Prediction step
	x_pred := kf.X
	p_pred := kf.P + kf.Q
	
	// Update step
	k := p_pred / (p_pred + kf.R) // Kalman gain
	kf.X = x_pred + k*(measurement-x_pred)
	kf.P = (1 - k) * p_pred
	
	return kf.X
}

// DynamicScalingKalmanFilter represents a Kalman filter for dynamic scaling factors
type DynamicScalingKalmanFilter struct {
	KalmanFilter
	ExpectedImpact float64
}

// NewDynamicScalingKalmanFilter creates a new Kalman filter for dynamic scaling
func NewDynamicScalingKalmanFilter(initialK, initialP, q, r float64) *DynamicScalingKalmanFilter {
	return &DynamicScalingKalmanFilter{
		KalmanFilter: KalmanFilter{
			X: initialK,
			P: initialP,
			Q: q,
			R: r,
		},
		ExpectedImpact: 0,
	}
}

// UpdateWithImpact updates the scaling factor based on observed vs expected impact
func (dskf *DynamicScalingKalmanFilter) UpdateWithImpact(observedImpact, expectedImpact float64) float64 {
	// Calculate difference between observed and expected impact
	diff := observedImpact - expectedImpact
	
	// Use standard Kalman filter update
	return dskf.Update(dskf.X + diff)
}

// OutlierDetectionKalmanFilter represents a Kalman filter for threshold parameters
type OutlierDetectionKalmanFilter struct {
	TauFilter *KalmanFilter
	KFilter   *KalmanFilter
}

// NewOutlierDetectionKalmanFilter creates Kalman filters for tau and k parameters
func NewOutlierDetectionKalmanFilter(initTau, initK, initP, qTau, qK, rTau, rK float64) *OutlierDetectionKalmanFilter {
	return &OutlierDetectionKalmanFilter{
		TauFilter: NewKalmanFilter(initTau, initP, qTau, rTau),
		KFilter:   NewKalmanFilter(initK, initP, qK, rK),
	}
}

// UpdateThresholds updates both tau and k thresholds based on observed max deviations
func (odkf *OutlierDetectionKalmanFilter) UpdateThresholds(maxMeanDev, maxMedianDev float64) (tau, k float64) {
	tau = odkf.TauFilter.Update(maxMeanDev)
	k = odkf.KFilter.Update(maxMedianDev)
	return tau, k
}

// RandomInRange returns a random float64 in the specified range [min, max]
func RandomInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomIntInRange returns a random int in the specified range [min, max]
func RandomIntInRange(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomTimeWindow generates a random time window in the given range
func RandomTimeWindow(minDays, maxDays int) time.Duration {
	days := RandomIntInRange(minDays, maxDays)
	return time.Duration(days) * 24 * time.Hour
}

// RandomHopLimit generates a random hop limit in the given range
func RandomHopLimit(min, max int) int {
	return RandomIntInRange(min, max)
}

// Clamp constrains a value to the specified range [min, max]
func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// TransactionWeight calculates the weight of a transaction based on amount, clustering, and time
func TransactionWeight(amount float64, clusterCount int, timestamp time.Time, current time.Time, beta, delta float64) float64 {
	// Log(1 + amount) - diminishing returns for large amounts
	amountTerm := math.Log1p(amount)
	
	// 1 + beta * (clusterCount - 1) - clustering factor
	clusterTerm := 1.0 + beta*float64(clusterCount-1)
	
	// e^(-delta * timeDiff) - time decay
	timeDiff := current.Sub(timestamp).Seconds()
	timeTerm := math.Exp(-delta * timeDiff)
	
	return amountTerm * clusterTerm * timeTerm
}

// CreditScoreShift calculates the credit score shift based on deviation from mean
func CreditScoreShift(deltaV, sigmaV, kFactor, gamma float64) float64 {
	standardizedDev := math.Abs(deltaV) / sigmaV
	return kFactor * math.Pow(standardizedDev, gamma)
}

// VectorizeScore embeds a normalized score in a 2D vector for cosine similarity
func VectorizeScore(score, mean, stdDev float64) Vector {
	normalizedScore := (score - mean) / stdDev
	return VectorFromValues(normalizedScore, 1)
}

// DeterministicSeed generates a deterministic seed based on a string and timestamp
func DeterministicSeed(input string, timestamp int64) int64 {
	// Combine input and timestamp
	combined := input + ":" + fmt.Sprintf("%d", timestamp)
	
	// Hash the combined string
	hash := sha256.Sum256([]byte(combined))
	
	// Use first 8 bytes of hash as int64 seed
	return int64(binary.BigEndian.Uint64(hash[:8]))
}

// DeterministicRandomInt generates a deterministic random integer in range [min, max]
func DeterministicRandomInt(seed int64, min, max int) int {
	// Create a new source and rand with the seed
	source := rand.NewSource(seed)
	rnd := rand.New(source)
	
	// Generate random number in range
	return min + rnd.Intn(max-min+1)
}