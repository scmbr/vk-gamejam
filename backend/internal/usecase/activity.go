package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type ActivityUsecase struct {
	repo repository.ActivityRepository
}

func NewActivityUsecase(r repository.ActivityRepository) *ActivityUsecase {
	return &ActivityUsecase{repo: r}
}

func (uc *ActivityUsecase) Create(ctx context.Context, a *domain.Activity) error {
	a.ID = uuid.NewString()
	a.CreatedAt = time.Now().UTC()

	return uc.repo.Create(ctx, a)
}

func (uc *ActivityUsecase) GetByChildProfileID(ctx context.Context, userID int64) ([]*domain.Activity, error) {
	return uc.repo.GetByChildProfileID(ctx, userID)
}
