package usecase

import (
	"context"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type ChildProfileUsecase struct {
	repo repository.ChildProfileRepository
}

func NewChildProfileUsecase(r repository.ChildProfileRepository) *ChildProfileUsecase {
	return &ChildProfileUsecase{r}
}

func (uc *ChildProfileUsecase) Get(ctx context.Context, userID int64) (*domain.ChildProfile, error) {
	return uc.repo.GetByUserID(ctx, userID)
}

func (uc *ChildProfileUsecase) Create(ctx context.Context, p *domain.ChildProfile) error {
	return uc.repo.Create(ctx, p)
}

func (uc *ChildProfileUsecase) Update(ctx context.Context, p *domain.ChildProfile) error {
	return uc.repo.Update(ctx, p)
}