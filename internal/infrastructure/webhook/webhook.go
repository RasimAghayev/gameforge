package webhook

import (
	"context"
	"log/slog"
)

type Dispatcher struct {
	logger *slog.Logger
}

func NewDispatcher(logger *slog.Logger) *Dispatcher {
	return &Dispatcher{logger: logger}
}

func (d *Dispatcher) Dispatch(ctx context.Context, url, event string, payload []byte) error {
	d.logger.Info("webhook dispatch", "url", url, "event", event)
	return nil
}

var _ = context.Background
