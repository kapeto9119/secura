package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config holds all the configuration for the application
type Config struct {
	// Server settings
	Port        string
	Environment string
	LogLevel    string

	// Database settings
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Service URLs
	NLPServiceURL string

	// Blockchain settings
	BlockchainNodeURL         string
	BlockchainContractAddress string

	// JWT settings
	JWTSecret      string
	JWTExpiryHours int

	// OpenAI settings
	OpenAIAPIKey string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()

	config := &Config{
		// Server settings
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("APP_ENV", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),

		// Database settings
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "secura"),
		DBPassword: getEnv("DB_PASSWORD", "securapassword"),
		DBName:     getEnv("DB_NAME", "secura"),

		// Service URLs
		NLPServiceURL: getEnv("NLP_SERVICE_URL", "http://localhost:8000"),

		// Blockchain settings
		BlockchainNodeURL:         getEnv("BLOCKCHAIN_NODE_URL", "http://localhost:8545"),
		BlockchainContractAddress: getEnv("BLOCKCHAIN_CONTRACT_ADDRESS", "0x0000000000000000000000000000000000000000"),

		// JWT settings
		JWTSecret:      getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiryHours: 24, // Default 24 hours

		// OpenAI settings
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", ""),
	}

	return config, nil
}

// SetupLogger configures the global logger
func SetupLogger(logLevel string) *zap.Logger {
	level := zap.InfoLevel

	switch strings.ToLower(logLevel) {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := config.Build()
	return logger
}

// Helper function to get environment variables with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
