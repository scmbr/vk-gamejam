package repository

import (
	"context"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
)

type PetRepository interface {
	GetState(ctx context.Context, userID int64) (*domain.Pet, error)
	SaveState(ctx context.Context, p *domain.Pet) error
}