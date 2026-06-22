package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	App    AppConfig
	GRPC   GRPCConfig
	DB     DatabaseConfig
	Logger LoggerConfig
	Server ServerConfig
}

type AppConfig struct {
	Name        string
	Environment string
}

type GRPCConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Host              string
	Port              string
	User              string
	Password          string
	Name              string
	SSLMode           string
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	ConnectRetries    int
	ConnectRetryDelay time.Duration
}

type LoggerConfig struct {
	Level  string
	Pretty bool
}

type ServerConfig struct {
	ShutdownTimeout time.Duration
}

func LoadConfig() (*Config, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	maxConns, err := getOptionalInt32("DB_MAX_CONNS", 10)
	if err != nil {
		return nil, err
	}

	minConns, err := getOptionalInt32("DB_MIN_CONNS", 1)
	if err != nil {
		return nil, err
	}

	maxConnLifetimeMinutes, err := getOptionalInt("DB_MAX_CONN_LIFETIME_MINUTES", 30)
	if err != nil {
		return nil, err
	}

	maxConnIdleMinutes, err := getOptionalInt("DB_MAX_CONN_IDLE_TIME_MINUTES", 5)
	if err != nil {
		return nil, err
	}

	connectRetries, err := getOptionalInt("DB_CONNECT_RETRIES", 3)
	if err != nil {
		return nil, err
	}

	connectRetryDelaySeconds, err := getOptionalInt("DB_CONNECT_RETRY_DELAY_SECONDS", 2)
	if err != nil {
		return nil, err
	}

	logPretty, err := getOptionalBool("LOG_PRETTY", true)
	if err != nil {
		return nil, err
	}

	shutdownTimeoutSeconds, err := getOptionalInt("SHUTDOWN_TIMEOUT_SECONDS", 10)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		App: AppConfig{
			Name:        getOptionalEnv("APP_NAME", "db-apeiron"),
			Environment: getOptionalEnv("ENVIRONMENT", "development"),
		},

		GRPC: GRPCConfig{
			Host: getOptionalEnv("GRPC_HOST", "127.0.0.1"),
			Port: getOptionalEnv("GRPC_PORT", "50051"),
		},

		DB: DatabaseConfig{
			Host:              getOptionalEnv("DB_HOST", "127.0.0.1"),
			Port:              getOptionalEnv("DB_PORT", "5432"),
			User:              getOptionalEnv("DB_USER", "postgres"),
			Password:          getEnv("DB_PASSWORD"),
			Name:              getOptionalEnv("DB_NAME", "apeiron"),
			SSLMode:           getOptionalEnv("DB_SSL_MODE", "disable"),
			MaxConns:          maxConns,
			MinConns:          minConns,
			MaxConnLifetime:   time.Duration(maxConnLifetimeMinutes) * time.Minute,
			MaxConnIdleTime:   time.Duration(maxConnIdleMinutes) * time.Minute,
			ConnectRetries:    connectRetries,
			ConnectRetryDelay: time.Duration(connectRetryDelaySeconds) * time.Second,
		},

		Logger: LoggerConfig{
			Level:  getOptionalEnv("LOG_LEVEL", "info"),
			Pretty: logPretty,
		},

		Server: ServerConfig{
			ShutdownTimeout: time.Duration(shutdownTimeoutSeconds) * time.Second,
		},
	}

	return cfg, nil
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func getOptionalEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getInt(key string) (int, error) {
	value := os.Getenv(key)

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf(
			"invalid integer value for %s: %w",
			key,
			err,
		)
	}

	return parsed, nil
}

func getOptionalInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf(
			"invalid integer value for %s: %w",
			key,
			err,
		)
	}

	return parsed, nil
}

func getInt32(key string) (int32, error) {
	value, err := getInt(key)
	if err != nil {
		return 0, err
	}

	return int32(value), nil
}

func getOptionalInt32(key string, fallback int32) (int32, error) {
	value, err := getOptionalInt(key, int(fallback))
	if err != nil {
		return 0, err
	}

	return int32(value), nil
}

func getBool(key string) (bool, error) {
	value := os.Getenv(key)

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf(
			"invalid boolean value for %s: %w",
			key,
			err,
		)
	}

	return parsed, nil
}

func getOptionalBool(key string, fallback bool) (bool, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf(
			"invalid boolean value for %s: %w",
			key,
			err,
		)
	}

	return parsed, nil
}
