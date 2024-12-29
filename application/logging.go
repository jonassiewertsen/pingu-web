package application

import (
	"log/slog"
	"os"
)

/**
 * Initialize the logging system based on the configuration
 *
 * Fiber has its own logging system, but I decided to use slog,
 * a simple but first party logging library for go.
 * Feel free to use any other logging library you prefer.
 *
 * Fiber Logging: https://docs.gofiber.io/api/log
 * About Slog: https://go.dev/blog/slog
 */
func initializeLogging(config *Config) *slog.Logger {
	var level slog.Level

	switch config.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warning":
		level = slog.LevelWarn
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		panic("Invalid log level. Please use one of: debug, info, warn, error")
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
