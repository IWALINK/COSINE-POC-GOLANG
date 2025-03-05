package utils

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricType defines the type of metric
type MetricType string

// MetricLabel defines a label for a metric
type MetricLabel struct {
	Name  string
	Value string
}

// Supported metric types
const (
	CounterMetric   MetricType = "counter"
	GaugeMetric     MetricType = "gauge"
	HistogramMetric MetricType = "histogram"
	SummaryMetric   MetricType = "summary"
)

// MetricsManager manages the collection and exposure of metrics
type MetricsManager struct {
	registry      *prometheus.Registry
	collectors    map[string]interface{}
	server        *http.Server
	mu            sync.RWMutex
	defaultLabels map[string]string
}

// MetricsOptions contains configuration for the metrics manager
type MetricsOptions struct {
	Address           string
	Path              string
	DefaultLabels     map[string]string
	HistogramBuckets  map[string][]float64
	SummaryObjectives map[string]map[float64]float64
}

// MetricsCollector is an alias for MetricsManager to maintain backward compatibility
// with modules that expect this name
type MetricsCollector = MetricsManager

// DefaultMetricsOptions returns default options for the metrics manager
func DefaultMetricsOptions() *MetricsOptions {
	return &MetricsOptions{
		Address:       ":9090",
		Path:          "/metrics",
		DefaultLabels: map[string]string{"app": "cosine"},
		HistogramBuckets: map[string][]float64{
			"default":          prometheus.DefBuckets,
			"latency_seconds":  []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 5, 10},
			"transaction_size": []float64{1, 10, 100, 1000, 10000, 100000, 1000000},
		},
		SummaryObjectives: map[string]map[float64]float64{
			"default": {0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
	}
}

// NewMetricsManager creates a new metrics manager
func NewMetricsManager(options *MetricsOptions) *MetricsManager {
	if options == nil {
		options = DefaultMetricsOptions()
	}

	registry := prometheus.NewRegistry()
	
	// Register default collectors
	registry.MustRegister(prometheus.NewGoCollector())
	registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))

	return &MetricsManager{
		registry:      registry,
		collectors:    make(map[string]interface{}),
		defaultLabels: options.DefaultLabels,
	}
}

// StartServer starts the metrics HTTP server
func (mm *MetricsManager) StartServer(options *MetricsOptions) error {
	if options == nil {
		options = DefaultMetricsOptions()
	}

	mux := http.NewServeMux()
	mux.Handle(options.Path, promhttp.HandlerFor(mm.registry, promhttp.HandlerOpts{}))

	mm.server = &http.Server{
		Addr:    options.Address,
		Handler: mux,
	}

	go func() {
		if err := mm.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			GetGlobalLogger().Error("metrics server failed", "error", err)
		}
	}()

	GetGlobalLogger().Info("metrics server started", "address", options.Address, "path", options.Path)
	return nil
}

// StopServer stops the metrics HTTP server
func (mm *MetricsManager) StopServer(ctx context.Context) error {
	if mm.server != nil {
		return mm.server.Shutdown(ctx)
	}
	return nil
}

// CreateMetric creates a new metric
func (mm *MetricsManager) CreateMetric(name, help string, metricType MetricType, labelNames []string) error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	if _, exists := mm.collectors[name]; exists {
		return fmt.Errorf("metric '%s' already exists", name)
	}

	var collector interface{}

	// Create metric with factory functions
	switch metricType {
	case CounterMetric:
		collector = promauto.With(mm.registry).NewCounterVec(
			prometheus.CounterOpts{
				Name: name,
				Help: help,
			},
			labelNames,
		)
	case GaugeMetric:
		collector = promauto.With(mm.registry).NewGaugeVec(
			prometheus.GaugeOpts{
				Name: name,
				Help: help,
			},
			labelNames,
		)
	case HistogramMetric:
		collector = promauto.With(mm.registry).NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    name,
				Help:    help,
				Buckets: prometheus.DefBuckets,
			},
			labelNames,
		)
	case SummaryMetric:
		collector = promauto.With(mm.registry).NewSummaryVec(
			prometheus.SummaryOpts{
				Name:       name,
				Help:       help,
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			},
			labelNames,
		)
	default:
		return fmt.Errorf("unsupported metric type: %s", metricType)
	}

	mm.collectors[name] = collector
	return nil
}

// CreateCustomHistogram creates a histogram with custom buckets
func (mm *MetricsManager) CreateCustomHistogram(name, help string, buckets []float64, labelNames []string) error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	if _, exists := mm.collectors[name]; exists {
		return fmt.Errorf("metric '%s' already exists", name)
	}

	collector := promauto.With(mm.registry).NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    name,
			Help:    help,
			Buckets: buckets,
		},
		labelNames,
	)

	mm.collectors[name] = collector
	return nil
}

// CreateCustomSummary creates a summary with custom objectives
func (mm *MetricsManager) CreateCustomSummary(name, help string, objectives map[float64]float64, labelNames []string) error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	if _, exists := mm.collectors[name]; exists {
		return fmt.Errorf("metric '%s' already exists", name)
	}

	collector := promauto.With(mm.registry).NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       name,
			Help:       help,
			Objectives: objectives,
		},
		labelNames,
	)

	mm.collectors[name] = collector
	return nil
}

// IncCounter increments a counter metric
func (mm *MetricsManager) IncCounter(name string, labelValues ...string) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if counter, ok := mm.collectors[name].(*prometheus.CounterVec); ok {
		counter.WithLabelValues(labelValues...).Inc()
	}
}

// AddCounter adds a value to a counter metric
func (mm *MetricsManager) AddCounter(name string, value float64, labelValues ...string) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if counter, ok := mm.collectors[name].(*prometheus.CounterVec); ok {
		counter.WithLabelValues(labelValues...).Add(value)
	}
}

// SetGauge sets a gauge metric value
func (mm *MetricsManager) SetGauge(name string, value float64, labelValues ...string) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if gauge, ok := mm.collectors[name].(*prometheus.GaugeVec); ok {
		gauge.WithLabelValues(labelValues...).Set(value)
	}
}

// IncGauge increments a gauge metric
func (mm *MetricsManager) IncGauge(name string, labelValues ...string) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if gauge, ok := mm.collectors[name].(*prometheus.GaugeVec); ok {
		gauge.WithLabelValues(labelValues...).Inc()
	}
}

// DecGauge decrements a gauge metric
func (mm *MetricsManager) DecGauge(name string, labelValues ...string) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if gauge, ok := mm.collectors[name].(*prometheus.GaugeVec); ok {
		gauge.WithLabelValues(labelValues...).Dec()
	}
}

// AddGauge adds a value to a gauge metric
func (mm *MetricsManager) AddGauge(name string, value float64, labelValues ...string) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if gauge, ok := mm.collectors[name].(*prometheus.GaugeVec); ok {
		gauge.WithLabelValues(labelValues...).Add(value)
	}
}

// ObserveHistogram observes a value for a histogram metric
func (mm *MetricsManager) ObserveHistogram(name string, value float64, labelValues ...string) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if histogram, ok := mm.collectors[name].(*prometheus.HistogramVec); ok {
		histogram.WithLabelValues(labelValues...).Observe(value)
	}
}

// ObserveSummary observes a value for a summary metric
func (mm *MetricsManager) ObserveSummary(name string, value float64, labelValues ...string) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	if summary, ok := mm.collectors[name].(*prometheus.SummaryVec); ok {
		summary.WithLabelValues(labelValues...).Observe(value)
	}
}

// TimeFunctionDuration times a function's execution and records it as a histogram observation
func (mm *MetricsManager) TimeFunctionDuration(name string, labelValues ...string) func() {
	start := time.Now()
	return func() {
		mm.ObserveHistogram(name, time.Since(start).Seconds(), labelValues...)
	}
}

// CollectDefaultValidatorMetrics creates and collects default metrics for validators
func (mm *MetricsManager) CollectDefaultValidatorMetrics() error {
	// Validator VRF selection metrics
	if err := mm.CreateMetric("validator_selections_total", "Total number of times a validator is selected in VRF subset", CounterMetric, []string{"validator_id"}); err != nil {
		return err
	}

	// Performance score metrics
	if err := mm.CreateMetric("validator_performance_score", "Current performance score of validator", GaugeMetric, []string{"validator_id"}); err != nil {
		return err
	}

	// Proposal metrics
	if err := mm.CreateMetric("validator_proposals_total", "Total number of proposals made by validator", CounterMetric, []string{"validator_id", "outcome"}); err != nil {
		return err
	}

	// Rewards metrics
	if err := mm.CreateMetric("validator_rewards_total", "Total rewards earned by validator", CounterMetric, []string{"validator_id"}); err != nil {
		return err
	}

	// Ping latency metrics
	buckets := []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 5, 10}
	if err := mm.CreateCustomHistogram("validator_ping_latency_seconds", "Validator ping latency in seconds", buckets, []string{"validator_id"}); err != nil {
		return err
	}

	return nil
}

// CollectDefaultCreditScoreMetrics creates and collects default metrics for credit scoring
func (mm *MetricsManager) CollectDefaultCreditScoreMetrics() error {
	// Vote metrics
	if err := mm.CreateMetric("votes_total", "Total number of votes", CounterMetric, []string{"vote_type"}); err != nil {
		return err
	}

	// Kalman filter parameter metrics
	if err := mm.CreateMetric("kalman_filter_tau", "Current tau parameter value", GaugeMetric, []string{}); err != nil {
		return err
	}
	if err := mm.CreateMetric("kalman_filter_k", "Current k parameter value", GaugeMetric, []string{}); err != nil {
		return err
	}

	// Dynamic scaling factor metrics
	if err := mm.CreateMetric("scaling_factor", "Current value of dynamic scaling factor", GaugeMetric, []string{"factor_type"}); err != nil {
		return err
	}

	// Outlier detection metrics
	if err := mm.CreateMetric("outlier_proposals_total", "Total number of outlier proposals", CounterMetric, []string{}); err != nil {
		return err
	}
	if err := mm.CreateMetric("inlier_proposals_total", "Total number of inlier proposals", CounterMetric, []string{}); err != nil {
		return err
	}

	// Credit score update metrics
	if err := mm.CreateMetric("credit_score_updates_total", "Total number of credit score updates", CounterMetric, []string{"update_type"}); err != nil {
		return err
	}

	buckets := []float64{0.001, 0.01, 0.1, 1.0, 10.0, 100.0, 1000.0}
	if err := mm.CreateCustomHistogram("credit_score_shift_magnitude", "Magnitude of credit score shifts", buckets, []string{"update_type"}); err != nil {
		return err
	}

	return nil
}

// CollectDefaultFraudDetectionMetrics creates and collects default metrics for fraud detection
func (mm *MetricsManager) CollectDefaultFraudDetectionMetrics() error {
	// Association risk metrics
	if err := mm.CreateMetric("association_penalties_total", "Total number of association penalties applied", CounterMetric, []string{}); err != nil {
		return err
	}

	buckets := []float64{0.1, 1.0, 5.0, 10.0, 50.0, 100.0, 500.0, 1000.0}
	if err := mm.CreateCustomHistogram("association_risk_score", "Distribution of association risk scores", buckets, []string{}); err != nil {
		return err
	}

	// Random window metrics
	if err := mm.CreateMetric("random_time_window_days", "Distribution of random time windows in days", HistogramMetric, []string{}); err != nil {
		return err
	}
	if err := mm.CreateMetric("random_hop_limit", "Distribution of random hop limits", HistogramMetric, []string{}); err != nil {
		return err
	}

	return nil
}

// CollectDefaultStorageMetrics creates and collects default metrics for storage
func (mm *MetricsManager) CollectDefaultStorageMetrics() error {
	// Wallet state metrics
	if err := mm.CreateMetric("wallet_count", "Total number of wallets in the system", GaugeMetric, []string{}); err != nil {
		return err
	}

	// Transaction metrics
	if err := mm.CreateMetric("transactions_total", "Total number of transactions", CounterMetric, []string{"transaction_type"}); err != nil {
		return err
	}

	// Storage metrics
	if err := mm.CreateMetric("ipfs_operations_total", "Total number of IPFS operations", CounterMetric, []string{"operation"}); err != nil {
		return err
	}

	buckets := []float64{0.001, 0.01, 0.1, 1.0, 5.0, 10.0, 30.0}
	if err := mm.CreateCustomHistogram("storage_operation_duration_seconds", "Duration of storage operations in seconds", buckets, []string{"operation"}); err != nil {
		return err
	}

	return nil
}

// CollectDefaultAPIMetrics creates and collects default metrics for API
func (mm *MetricsManager) CollectDefaultAPIMetrics() error {
	// API request metrics
	if err := mm.CreateMetric("api_requests_total", "Total number of API requests", CounterMetric, []string{"endpoint", "method", "status"}); err != nil {
		return err
	}

	buckets := []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1.0, 5.0, 10.0}
	if err := mm.CreateCustomHistogram("api_request_duration_seconds", "API request duration in seconds", buckets, []string{"endpoint", "method"}); err != nil {
		return err
	}

	return nil
}

// CollectDefaultNetworkMetrics creates and collects default metrics for network
func (mm *MetricsManager) CollectDefaultNetworkMetrics() error {
	// Peer metrics
	if err := mm.CreateMetric("peers_count", "Current number of connected peers", GaugeMetric, []string{}); err != nil {
		return err
	}

	// Message metrics
	if err := mm.CreateMetric("messages_total", "Total number of P2P messages", CounterMetric, []string{"message_type", "direction"}); err != nil {
		return err
	}

	buckets := []float64{10, 100, 1000, 10000, 100000, 1000000}
	if err := mm.CreateCustomHistogram("message_size_bytes", "Size of P2P messages in bytes", buckets, []string{"message_type"}); err != nil {
		return err
	}

	return nil
}

// CollectAllDefaultMetrics creates and collects all default metrics
func (mm *MetricsManager) CollectAllDefaultMetrics() error {
	if err := mm.CollectDefaultValidatorMetrics(); err != nil {
		return err
	}
	if err := mm.CollectDefaultCreditScoreMetrics(); err != nil {
		return err
	}
	if err := mm.CollectDefaultFraudDetectionMetrics(); err != nil {
		return err
	}
	if err := mm.CollectDefaultStorageMetrics(); err != nil {
		return err
	}
	if err := mm.CollectDefaultAPIMetrics(); err != nil {
		return err
	}
	if err := mm.CollectDefaultNetworkMetrics(); err != nil {
		return err
	}
	return nil
}

// Global metrics manager instance
var (
	globalMetrics *MetricsManager
	metricsOnce   sync.Once
)

// GetGlobalMetrics returns the global metrics manager instance
func GetGlobalMetrics() *MetricsManager {
	if globalMetrics == nil {
		// Initialize with defaults if not explicitly initialized
		_ = InitGlobalMetrics(nil)
	}
	return globalMetrics
}

// InitGlobalMetrics initializes the global metrics manager with options
func InitGlobalMetrics(options *MetricsOptions) error {
	var err error
	metricsOnce.Do(func() {
		globalMetrics = NewMetricsManager(options)
		if err = globalMetrics.CollectAllDefaultMetrics(); err != nil {
			return
		}
		if options != nil {
			err = globalMetrics.StartServer(options)
		}
	})
	return err
}

// SetupMetrics configures metrics based on configuration
func SetupMetrics(config *ConfigManager) (*MetricsManager, error) {
	// Load metrics configuration from ConfigManager
	options := &MetricsOptions{
		Address:       config.GetString("metrics.address"),
		Path:          config.GetString("metrics.path"),
		DefaultLabels: map[string]string{"app": "cosine", "env": config.GetString("node.environment")},
	}

	// Set defaults if not specified
	if options.Address == "" {
		options.Address = ":9090"
	}
	if options.Path == "" {
		options.Path = "/metrics"
	}

	// Create and initialize the global metrics manager
	metrics := NewMetricsManager(options)
	if err := metrics.CollectAllDefaultMetrics(); err != nil {
		return nil, err
	}

	// Start the metrics server if enabled
	if config.GetBool("metrics.enabled") {
		if err := metrics.StartServer(options); err != nil {
			return nil, err
		}
	}

	globalMetrics = metrics
	return metrics, nil
}