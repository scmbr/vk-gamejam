package repository

import (
	"context"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
}