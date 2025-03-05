package utils

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LogLevel represents the severity level of a log message
type LogLevel string

// Log levels
const (
	DebugLevel   LogLevel = "debug"
	InfoLevel    LogLevel = "info"
	WarningLevel LogLevel = "warning"
	ErrorLevel   LogLevel = "error"
	FatalLevel   LogLevel = "fatal"
)

// Logger wraps zap.Logger to provide structured logging capabilities
type Logger struct {
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	level       zap.AtomicLevel
	options     *LoggerOptions
	fields      map[string]interface{}
	mu          sync.RWMutex
}

// LoggerOptions contains configuration for the logger
type LoggerOptions struct {
	Level           LogLevel
	OutputPaths     []string
	ErrorOutputPaths []string
	Encoding        string
	Development     bool
	CallerSkip      int
	MaxSize         int  // in megabytes
	MaxBackups      int
	MaxAge          int  // in days
	Compress        bool
	ContextKeys     []string
}

// DefaultLoggerOptions returns default logger options
func DefaultLoggerOptions() *LoggerOptions {
	return &LoggerOptions{
		Level:        InfoLevel,
		OutputPaths:  []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:     "json",
		Development:  false,
		CallerSkip:   1,
		MaxSize:      100, // 100 MB
		MaxBackups:   5,
		MaxAge:       30,  // 30 days
		Compress:     true,
		ContextKeys:  []string{"requestID", "walletAddress", "validatorID"},
	}
}

// convertZapLevel converts LogLevel to zapcore.Level
func convertZapLevel(level LogLevel) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarningLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// NewLogger creates a new logger with the given options
func NewLogger(options *LoggerOptions) (*Logger, error) {
	if options == nil {
		options = DefaultLoggerOptions()
	}

	// Configure level
	level := zap.NewAtomicLevelAt(convertZapLevel(options.Level))

	// Create file rotation hooks for file outputs
	var cores []zapcore.Core
	encoder := getEncoder(options)

	// Process output paths to set up cores
	for _, path := range options.OutputPaths {
		if path == "stdout" {
			core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
			cores = append(cores, core)
		} else if path == "stderr" {
			core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), level)
			cores = append(cores, core)
		} else {
			// Ensure directory exists
			dir := filepath.Dir(path)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return nil, fmt.Errorf("failed to create log directory: %w", err)
				}
			}

			// Set up log rotation
			writer := zapcore.AddSync(&lumberjack.Logger{
				Filename:   path,
				MaxSize:    options.MaxSize,
				MaxBackups: options.MaxBackups,
				MaxAge:     options.MaxAge,
				Compress:   options.Compress,
			})
			core := zapcore.NewCore(encoder, writer, level)
			cores = append(cores, core)
		}
	}

	// Combine cores
	combinedCore := zapcore.NewTee(cores...)

	// Build logger options
	zapOptions := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(options.CallerSkip),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	// Add development mode if enabled
	if options.Development {
		zapOptions = append(zapOptions, zap.Development())
	}

	// Build logger
	zapLogger := zap.New(combinedCore, zapOptions...)

	return &Logger{
		logger:      zapLogger,
		sugarLogger: zapLogger.Sugar(),
		level:       level,
		options:     options,
		fields:      make(map[string]interface{}),
	}, nil
}

// getEncoder creates the appropriate encoder based on configuration
func getEncoder(options *LoggerOptions) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	if options.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if options.Encoding == "console" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// With returns a logger with the specified fields attached
func (l *Logger) With(fields map[string]interface{}) *Logger {
	l.mu.RLock()
	defer l.mu.RUnlock()

	// Create a new fields map combining existing and new fields
	newFields := make(map[string]interface{}, len(l.fields)+len(fields))
	for k, v := range l.fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}

	// Create a new sugared logger with the fields
	// Convert map to alternating key-value pairs for Sugar().With
	args := make([]interface{}, 0, len(newFields)*2)
	for k, v := range newFields {
		args = append(args, k, v)
	}

	return &Logger{
		logger:      l.logger,
		sugarLogger: l.sugarLogger.With(args...),
		level:       l.level,
		options:     l.options,
		fields:      newFields,
	}
}

// WithComponent returns a logger for a specific component
func (l *Logger) WithComponent(component string) *Logger {
	return l.With(map[string]interface{}{"component": component})
}

// WithContext extracts log fields from context and returns a logger with those fields
func (l *Logger) WithContext(ctx context.Context) *Logger {
	fields := make(map[string]interface{})
	
	// Extract fields from context based on option keys
	for _, key := range l.options.ContextKeys {
		if value := ctx.Value(key); value != nil {
			fields[key] = value
		}
	}
	
	if len(fields) > 0 {
		return l.With(fields)
	}
	return l
}

// SetLevel dynamically changes the logging level
func (l *Logger) SetLevel(level LogLevel) {
	l.level.SetLevel(convertZapLevel(level))
}

// GetLevel returns the current logging level
func (l *Logger) GetLevel() LogLevel {
	// Map zapcore.Level back to our LogLevel
	switch l.level.Level() {
	case zapcore.DebugLevel:
		return DebugLevel
	case zapcore.InfoLevel:
		return InfoLevel
	case zapcore.WarnLevel:
		return WarningLevel
	case zapcore.ErrorLevel:
		return ErrorLevel
	case zapcore.FatalLevel:
		return FatalLevel
	default:
		return InfoLevel
	}
}

// Debug logs a message at debug level
func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.sugarLogger.Debugw(msg, fields...)
}

// Info logs a message at info level
func (l *Logger) Info(msg string, fields ...interface{}) {
	l.sugarLogger.Infow(msg, fields...)
}

// Warn logs a message at warning level
func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.sugarLogger.Warnw(msg, fields...)
}

// Error logs a message at error level
func (l *Logger) Error(msg string, fields ...interface{}) {
	l.sugarLogger.Errorw(msg, fields...)
}

// Fatal logs a message at fatal level and then calls os.Exit(1)
func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.sugarLogger.Fatalw(msg, fields...)
}

// AddWriter adds a custom writer to the logger
func (l *Logger) AddWriter(w io.Writer, level LogLevel) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Create a new encoder
	encoder := getEncoder(l.options)
	
	// Create a new core with the provided writer
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(w),
		zap.NewAtomicLevelAt(convertZapLevel(level)),
	)
	
	// Create a new logger with the additional core
	oldLogger := l.logger
	
	// Build logger options
	zapOptions := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(l.options.CallerSkip),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	// Add development mode if enabled
	if l.options.Development {
		zapOptions = append(zapOptions, zap.Development())
	}
	
	l.logger = zap.New(zapcore.NewTee(oldLogger.Core(), core), zapOptions...)
	l.sugarLogger = l.logger.Sugar()
	
	return nil
}

// Close flushes any buffered log entries
func (l *Logger) Close() error {
	return l.logger.Sync()
}

// Global logger instance
var (
	globalLogger *Logger
	loggerOnce   sync.Once
)

// GetGlobalLogger returns the global logger instance
func GetGlobalLogger() *Logger {
	if globalLogger == nil {
		// Initialize with defaults if not explicitly initialized
		_ = InitGlobalLogger(nil)
	}
	return globalLogger
}

// InitGlobalLogger initializes the global logger with options
func InitGlobalLogger(options *LoggerOptions) error {
	var err error
	loggerOnce.Do(func() {
		globalLogger, err = NewLogger(options)
	})
	return err
}

// SetupLogging configures logging based on configuration
func SetupLogging(config *ConfigManager) (*Logger, error) {
	// Load logging configuration from ConfigManager
	options := &LoggerOptions{
		Level:        LogLevel(config.GetString("logging.level")),
		OutputPaths:  config.GetStringSlice("logging.outputs"),
		Encoding:     config.GetString("logging.encoding"),
		Development:  config.GetString("node.environment") != "production",
		CallerSkip:   1,
		MaxSize:      config.GetInt("logging.rotate.maxSize"),
		MaxBackups:   config.GetInt("logging.rotate.maxBackups"),
		MaxAge:       config.GetInt("logging.rotate.maxAge"),
		Compress:     config.GetBool("logging.rotate.compress"),
	}

	// Set defaults if not specified
	if len(options.OutputPaths) == 0 {
		options.OutputPaths = []string{"stdout", fmt.Sprintf("./logs/cosine-%s.log", time.Now().Format("2006-01-02"))}
	}
	
	// Create and initialize the global logger
	logger, err := NewLogger(options)
	if err != nil {
		return nil, err
	}
	
	// Register logger to receive configuration updates
	config.OnConfigChange("logging.level", func(value interface{}) {
		if strLevel, ok := value.(string); ok {
			logger.SetLevel(LogLevel(strLevel))
			logger.Info("Log level changed", "level", strLevel)
		}
	})
	
	globalLogger = logger
	return logger, nil
}