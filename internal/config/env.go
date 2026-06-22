package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	loadEnvFile()

	requiredEnvs := []string{
		"APP_NAME",
		"ENVIRONMENT",

		"GRPC_HOST",
		"GRPC_PORT",

		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_SSL_MODE",

		"DB_MAX_CONNS",
		"DB_MIN_CONNS",
		"DB_MAX_CONN_LIFETIME_MINUTES",
		"DB_MAX_CONN_IDLE_TIME_MINUTES",

		"DB_CONNECT_RETRIES",
		"DB_CONNECT_RETRY_DELAY_SECONDS",

		"LOG_LEVEL",
		"LOG_PRETTY",

		"SHUTDOWN_TIMEOUT_SECONDS",
	}

	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			return fmt.Errorf(
				"missing required environment variable: %s",
				env,
			)
		}
	}

	return nil
}

func loadEnvFile() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(
			"[WARNING] .env file not found, using system environment variables",
		)
	}
}
