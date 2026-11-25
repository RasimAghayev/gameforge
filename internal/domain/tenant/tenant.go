package tenant

import (
	"time"

	"github.com/google/uuid"
)

type TenantStatus string

const (
	TenantStatusActive    TenantStatus = "active"
	TenantStatusSuspended TenantStatus = "suspended"
	TenantStatusArchived  TenantStatus = "archived"
)

type TenantConfig map[string]interface{}

type Tenant struct {
	ID        uuid.UUID
	Name      string
	Slug      string
	Config    TenantConfig
	Status    TenantStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	Create(ctx, t *Tenant) error
	Get(ctx, id uuid.UUID) (*Tenant, error)
	GetBySlug(ctx, slug string) (*Tenant, error)
	List(ctx, tenantID uuid.UUID) ([]*Tenant, error)
	Update(ctx, t *Tenant) error
}
