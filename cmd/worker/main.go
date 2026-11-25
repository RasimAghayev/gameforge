package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/rasimaghayev/gameforge/internal/config"
	"github.com/rasimaghayev/gameforge/internal/service/eventprocessor"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	ep, err := eventprocessor.New(cfg, logger)
	if err != nil {
		logger.Error("failed to create event processor", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := ep.Run(ctx); err != nil && err != context.Canceled {
		logger.Error("event processor error", "error", err)
		os.Exit(1)
	}
	logger.Info("event processor stopped")
}
