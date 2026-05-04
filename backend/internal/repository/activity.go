package repository

import (
	"context"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
)

type ActivityRepository interface {
	Create(ctx context.Context, a *domain.Activity) error
	GetByChildProfileID(ctx context.Context, childProfileID int64) ([]*domain.Activity, error)
}
