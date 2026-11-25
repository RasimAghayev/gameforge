package mongo

import (
	"context"
	"log/slog"
)

type EventStore struct {
	logger *slog.Logger
}

func NewEventStore(logger *slog.Logger) *EventStore {
	return &EventStore{logger: logger}
}

func (s *EventStore) Insert(ctx context.Context, tenantID, playerID, eventType, source string, payload map[string]interface{}) error {
	s.logger.Info("event inserted", "tenant", tenantID, "player", playerID, "type", eventType, "source", source)
	return nil
}

var _ = context.Background
