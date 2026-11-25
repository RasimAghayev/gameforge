package redis

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type Leaderboard struct {
	client *redis.Client
	logger *slog.Logger
}

func NewLeaderboard(addr, password string, logger *slog.Logger) (*Leaderboard, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &Leaderboard{client: client, logger: logger}, nil
}

func (l *Leaderboard) UpdateScore(ctx context.Context, key string, member string, score float64) error {
	return l.client.ZAdd(ctx, key, redis.Z{Score: score, Member: member}).Err()
}

func (l *Leaderboard) TopN(ctx context.Context, key string, n int) ([]string, error) {
	return l.client.ZRevRange(ctx, key, 0, int64(n-1)).Result()
}

func (l *Leaderboard) Rank(ctx context.Context, key, member string) (int64, error) {
	return l.client.ZRevRank(ctx, key, member).Result()
}

func (l *Leaderboard) Close() error {
	return l.client.Close()
}

var _ = context.Background
