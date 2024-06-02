package main

import (
	"log/slog"
	"os"

	"github.com/2Cheetah/was_your_pass_leaked/cmd/api"
)

func getLogLevelFromEnv() slog.Level {
	levelStr := os.Getenv("LOGLEVEL")
	switch levelStr {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo // Default log level
	}
}

func main() {

	logLevel := getLogLevelFromEnv()
	currentLogLevel := slog.SetLogLoggerLevel(logLevel)
	slog.Any("Log level", logLevel)
	defer slog.SetLogLoggerLevel(currentLogLevel) // revert changes after the example

	server := api.NewAPIServer("localhost:8080")
	if err := server.Run(); err != nil {
		slog.Error("Error, while trying to start server:", "host", server.Addr)
		panic(err)
	}
}
