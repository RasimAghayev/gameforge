package leaderboard

import (
	"time"

	"github.com/google/uuid"
)

type LeaderboardType string

const (
	LeaderboardTypeGlobal   LeaderboardType = "global"
	LeaderboardTypeWeekly   LeaderboardType = "weekly"
	LeaderboardTypeMonthly  LeaderboardType = "monthly"
	LeaderboardTypeCustom   LeaderboardType = "custom"
)

type LeaderboardScope string

const (
	LeaderboardScopeGlobal  LeaderboardScope = "global"
	LeaderboardScopeFriends LeaderboardScope = "friends"
	LeaderboardScopeRegion  LeaderboardScope = "region"
)

type ResetPeriod string

const (
	ResetPeriodNone   ResetPeriod = "none"
	ResetPeriodDaily  ResetPeriod = "daily"
	ResetPeriodWeekly ResetPeriod = "weekly"
	ResetPeriodMonthly ResetPeriod = "monthly"
)

type LeaderboardStatus string

const (
	LeaderboardStatusActive LeaderboardStatus = "active"
	LeaderboardStatusInactive LeaderboardStatus = "inactive"
)

type Leaderboard struct {
	ID          uuid.UUID
	TenantID    uuid.UUID
	Name        string
	Type        LeaderboardType
	Metric      string
	Scope       LeaderboardScope
	ResetPeriod ResetPeriod
	MaxEntries  int
	Status      LeaderboardStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Entry struct {
	PlayerID uuid.UUID
	Score    int64
	Rank     int
}

type Repository interface {
	Create(ctx, l *Leaderboard) error
	Get(ctx, id uuid.UUID) (*Leaderboard, error)
	ListByTenant(ctx, tenantID uuid.UUID) ([]*Leaderboard, error)
	Update(ctx, l *Leaderboard) error
}
