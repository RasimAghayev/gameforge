package badge

import (
	"time"

	"github.com/google/uuid"
)

type Rarity string

const (
	RarityCommon    Rarity = "common"
	RarityRare      Rarity = "rare"
	RarityEpic      Rarity = "epic"
	RarityLegendary Rarity = "legendary"
)

type BadgeStatus string

const (
	BadgeStatusActive    BadgeStatus = "active"
	BadgeStatusDeprecated BadgeStatus = "deprecated"
)

type BadgeCriteria struct {
	Expression string            `json:"expression"`
	Meta       map[string]string `json:"meta"`
}

type Badge struct {
	ID          uuid.UUID
	TenantID    uuid.UUID
	Name        string
	Description string
	IconURL     string
	Category    string
	Rarity      Rarity
	Criteria    BadgeCriteria
	Status      BadgeStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Create(ctx, b *Badge) error
	Get(ctx, id uuid.UUID) (*Badge, error)
	ListByTenant(ctx, tenantID uuid.UUID) ([]*Badge, error)
}
