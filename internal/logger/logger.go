package logger

import (
	"io"
	"os"
	"strings"
	"time"

	"db-apeiron/internal/config"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func Initialize(cfg config.LoggerConfig) {
	output := getOutput(cfg.Pretty)

	level := parseLogLevel(cfg.Level)

	logger := zerolog.New(output).
		Level(level).
		With().
		Timestamp().
		Logger()

	Log = logger

	Log.Info().
		Str("system", "logger").
		Str("level", strings.ToLower(cfg.Level)).
		Bool("pretty", cfg.Pretty).
		Msg("logger initialized")
}

func getOutput(pretty bool) io.Writer {
	if !pretty {
		return os.Stdout
	}

	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
}

func parseLogLevel(level string) zerolog.Level {
	switch strings.ToLower(level) {
	case "trace":
		return zerolog.TraceLevel

	case "debug":
		return zerolog.DebugLevel

	case "info":
		return zerolog.InfoLevel

	case "warn":
		return zerolog.WarnLevel

	case "error":
		return zerolog.ErrorLevel

	case "fatal":
		return zerolog.FatalLevel

	case "panic":
		return zerolog.PanicLevel

	default:
		return zerolog.InfoLevel
	}
}

func WithComponent(component string) zerolog.Logger {
	return Log.With().
		Str("component", component).
		Logger()
}
