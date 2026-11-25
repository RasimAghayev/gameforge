package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/rasimaghayev/gameforge/internal/config"
	"github.com/rasimaghayev/gameforge/internal/service/analytics"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	sched, err := analytics.NewScheduler(cfg, logger)
	if err != nil {
		logger.Error("failed to create scheduler", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := sched.Run(ctx); err != nil && err != context.Canceled {
		logger.Error("scheduler error", "error", err)
		os.Exit(1)
	}
	logger.Info("scheduler stopped")
}
