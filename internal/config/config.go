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

	maxConns, err := getInt32("DB_MAX_CONNS")
	if err != nil {
		return nil, err
	}

	minConns, err := getInt32("DB_MIN_CONNS")
	if err != nil {
		return nil, err
	}

	maxConnLifetimeMinutes, err := getInt("DB_MAX_CONN_LIFETIME_MINUTES")
	if err != nil {
		return nil, err
	}

	maxConnIdleMinutes, err := getInt("DB_MAX_CONN_IDLE_TIME_MINUTES")
	if err != nil {
		return nil, err
	}

	connectRetries, err := getInt("DB_CONNECT_RETRIES")
	if err != nil {
		return nil, err
	}

	connectRetryDelaySeconds, err := getInt("DB_CONNECT_RETRY_DELAY_SECONDS")
	if err != nil {
		return nil, err
	}

	logPretty, err := getBool("LOG_PRETTY")
	if err != nil {
		return nil, err
	}

	shutdownTimeoutSeconds, err := getInt("SHUTDOWN_TIMEOUT_SECONDS")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		App: AppConfig{
			Name:        getEnv("APP_NAME"),
			Environment: getEnv("ENVIRONMENT"),
		},

		GRPC: GRPCConfig{
			Host: getEnv("GRPC_HOST"),
			Port: getEnv("GRPC_PORT"),
		},

		DB: DatabaseConfig{
			Host:              getEnv("DB_HOST"),
			Port:              getEnv("DB_PORT"),
			User:              getEnv("DB_USER"),
			Password:          getEnv("DB_PASSWORD"),
			Name:              getEnv("DB_NAME"),
			SSLMode:           getEnv("DB_SSL_MODE"),
			MaxConns:          maxConns,
			MinConns:          minConns,
			MaxConnLifetime:   time.Duration(maxConnLifetimeMinutes) * time.Minute,
			MaxConnIdleTime:   time.Duration(maxConnIdleMinutes) * time.Minute,
			ConnectRetries:    connectRetries,
			ConnectRetryDelay: time.Duration(connectRetryDelaySeconds) * time.Second,
		},

		Logger: LoggerConfig{
			Level:  getEnv("LOG_LEVEL"),
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

func getInt32(key string) (int32, error) {
	value, err := getInt(key)
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
