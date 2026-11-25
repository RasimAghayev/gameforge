package gamification

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/rasimaghayev/gameforge/internal/domain/points"
)

type Service struct {
	points points.Repository
	logger *slog.Logger
}

func NewPointsService(p points.Repository, logger *slog.Logger) *Service {
	return &Service{points: p, logger: logger}
}

func (s *Service) AwardPoints(ctx context.Context, tenantID, playerID uuid.UUID, amount int64, reason, source string) (*points.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	tx := &points.Transaction{
		ID:        uuid.New(),
		TenantID:  tenantID,
		PlayerID:  playerID,
		Amount:    amount,
		Reason:    reason,
		Source:    source,
		CreatedAt: time.Now(),
	}
	if err := s.points.CreatePointsTransaction(ctx, tx); err != nil {
		return nil, fmt.Errorf("award points: %w", err)
	}
	s.logger.Info("points awarded", "tenant", tenantID, "player", playerID, "amount", amount, "reason", reason)
	return tx, nil
}

func (s *Service) DeductPoints(ctx context.Context, tenantID, playerID uuid.UUID, amount int64, reason string) (*points.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	balance, err := s.points.GetBalance(ctx, tenantID, playerID)
	if err != nil {
		return nil, fmt.Errorf("get balance: %w", err)
	}
	if balance-amount < 0 {
		return nil, errors.New("insufficient balance")
	}
	tx := &points.Transaction{
		ID:        uuid.New(),
		TenantID:  tenantID,
		PlayerID:  playerID,
		Amount:    -amount,
		Reason:    reason,
		Source:    "api",
		CreatedAt: time.Now(),
	}
	if err := s.points.CreatePointsTransaction(ctx, tx); err != nil {
		return nil, fmt.Errorf("deduct points: %w", err)
	}
	s.logger.Info("points deducted", "tenant", tenantID, "player", playerID, "amount", amount, "reason", reason)
	return tx, nil
}

func (s *Service) GetBalance(ctx context.Context, tenantID, playerID uuid.UUID) (int64, error) {
	return s.points.GetBalance(ctx, tenantID, playerID)
}
