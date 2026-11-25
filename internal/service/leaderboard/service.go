package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/rasimaghayev/gameforge/internal/domain/leaderboard"
)

type Service struct {
	lb leaderboard.Repository
}

func NewLeaderboardService(lb leaderboard.Repository) *Service {
	return &Service{lb: lb}
}

func (s *Service) GetEntries(ctx context.Context, id uuid.UUID, limit, offset int) ([]leaderboard.Entry, error) {
	return nil, nil
}

func (s *Service) GetRank(ctx context.Context, id, playerID uuid.UUID) (int, error) {
	return 0, nil
}

var _ = leaderboard.Repository(nil)
var _ = context.Background
var _ = fmt.Sprintf
