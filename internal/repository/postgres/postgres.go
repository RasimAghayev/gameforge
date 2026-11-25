package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rasimaghayev/gameforge/internal/domain/tenant"
	"github.com/rasimaghayev/gameforge/internal/domain/player"
	"github.com/rasimaghayev/gameforge/internal/domain/points"
)

type Store struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
}

func NewPostgres(dsn string, logger *slog.Logger) (*Store, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}
	return &Store{pool: pool, logger: logger}, nil
}

func (s *Store) Close() {
	if s.pool != nil {
		s.pool.Close()
	}
}

var _ tenant.Repository = (*Store)(nil)
var _ player.Repository = (*Store)(nil)
var _ points.Repository = (*Store)(nil)

func (s *Store) Create(ctx context.Context, t *tenant.Tenant) error {
	_, err := s.pool.Exec(ctx,
		`INSERT INTO tenants (id, name, slug, config, status, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		t.ID, t.Name, t.Slug, t.Config, t.Status, t.CreatedAt, t.UpdatedAt,
	)
	return err
}

func (s *Store) Get(ctx context.Context, id uuid.UUID) (*tenant.Tenant, error) {
	return nil, nil
}

func (s *Store) GetBySlug(ctx context.Context, slug string) (*tenant.Tenant, error) {
	return nil, nil
}

func (s *Store) List(ctx context.Context, tenantID uuid.UUID) ([]*tenant.Tenant, error) {
	return nil, nil
}

func (s *Store) Update(ctx context.Context, t *tenant.Tenant) error {
	return nil
}

func (s *Store) CreatePlayer(ctx context.Context, p *player.Player) error { return nil }

func (s *Store) GetPlayer(ctx context.Context, id uuid.UUID) (*player.Player, error) { return nil, nil }

func (s *Store) GetByExternalID(ctx context.Context, tenantID uuid.UUID, externalID string) (*player.Player, error) { return nil, nil }

func (s *Store) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]*player.Player, error) { return nil, nil }

func (s *Store) UpdateStats(ctx context.Context, id uuid.UUID, stats player.PlayerStats) error { return nil }

func (s *Store) CreatePointsTransaction(ctx context.Context, tx *points.Transaction) error { return nil }

func (s *Store) GetBalance(ctx context.Context, tenantID, playerID uuid.UUID) (int64, error) { return 0, nil }

func (s *Store) History(ctx context.Context, tenantID, playerID uuid.UUID, limit, offset int) ([]*points.Transaction, error) { return nil, nil }
