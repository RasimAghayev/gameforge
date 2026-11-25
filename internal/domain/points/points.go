package points

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	PlayerID  uuid.UUID
	Amount    int64
	Reason    string
	Source    string
	Metadata  map[string]interface{}
	CreatedAt time.Time
}

type Repository interface {
	Create(ctx, tx *Transaction) error
	GetBalance(ctx, tenantID, playerID uuid.UUID) (int64, error)
	History(ctx, tenantID, playerID uuid.UUID, limit, offset int) ([]*Transaction, error)
}
