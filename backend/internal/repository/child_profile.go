package repository

import (
	"context"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
)

type ChildProfileRepository interface {
	GetByUserID(ctx context.Context, userID int64) (*domain.ChildProfile, error)
	Update(ctx context.Context, profile *domain.ChildProfile) error
	Create(ctx context.Context, profile *domain.ChildProfile) error
}