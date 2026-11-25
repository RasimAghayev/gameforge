package player

import (
	"time"

	"github.com/google/uuid"
)

type PlayerStatus string

const (
	PlayerStatusActive   PlayerStatus = "active"
	PlayerStatusBanned   PlayerStatus = "banned"
	PlayerStatusInactive PlayerStatus = "inactive"
)

type PlayerProfile struct {
	Name    string            `json:"name"`
	Avatar  string            `json:"avatar"`
	Meta    map[string]string `json:"meta"`
}

type PlayerStats struct {
	TotalPoints int64     `json:"total_points"`
	Level       int       `json:"level"`
	Streak      int       `json:"streak"`
}

type Player struct {
	ID          uuid.UUID
	TenantID    uuid.UUID
	ExternalID  string
	Profile     PlayerProfile
	Stats       PlayerStats
	Status      PlayerStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Create(ctx, p *Player) error
	Get(ctx, id uuid.UUID) (*Player, error)
	GetByExternalID(ctx, tenantID uuid.UUID, externalID string) (*Player, error)
	ListByTenant(ctx, tenantID uuid.UUID) ([]*Player, error)
	UpdateStats(ctx, id uuid.UUID, stats PlayerStats) error
}
