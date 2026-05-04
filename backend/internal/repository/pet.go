package repository

import (
	"context"
	"time"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
)

type PetRepository interface {
	GetState(ctx context.Context, userID int64) (*domain.Pet, error)
	Create(ctx context.Context, p *domain.Pet) error
	SaveState(ctx context.Context, p *domain.Pet) error
	UpdateLastOnline(ctx context.Context, userID int64, t time.Time) error
}
