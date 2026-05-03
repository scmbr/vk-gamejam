package usecase

import (
	"context"
	"time"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type ChildProfileUsecase struct {
	repo repository.ChildProfileRepository
}

func NewChildProfileUsecase(r repository.ChildProfileRepository) *ChildProfileUsecase {
	return &ChildProfileUsecase{r}
}

func (uc *ChildProfileUsecase) GetByUserID(ctx context.Context, userID int64) (*domain.ChildProfile, error) {
	return uc.repo.GetByUserID(ctx, userID)
}

func (uc *ChildProfileUsecase) Create(ctx context.Context, p *domain.ChildProfile) error {
	return uc.repo.Create(ctx, p)
}

func (uc *ChildProfileUsecase) Update(ctx context.Context, p *domain.ChildProfile) error {
	return uc.repo.Update(ctx, p)
}
func (uc *ChildProfileUsecase) MarkLogin(ctx context.Context, userID int64) error {
	profile, err := uc.repo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	profile.LastLogin = time.Now().UTC()
	profile.IsFirstLaunch = false

	return uc.repo.Update(ctx, profile)
}

func (uc *ChildProfileUsecase) MarkLogout(ctx context.Context, userID int64) error {
	profile, err := uc.repo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	profile.LastLogout = time.Now().UTC()

	return uc.repo.Update(ctx, profile)
}