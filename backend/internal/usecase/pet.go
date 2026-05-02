package usecase

import (
	"context"
	"time"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type PetUsecase struct {
	repo repository.PetRepository
}

func NewPetUsecase(r repository.PetRepository) *PetUsecase {
	return &PetUsecase{r}
}

func (uc *PetUsecase) GetState(ctx context.Context, userID int64) (*domain.Pet, error) {
	return uc.repo.GetState(ctx, userID)
}

func (uc *PetUsecase) SaveState(ctx context.Context, p *domain.Pet) error {
	p.LastOnline = time.Now().UTC()
	return uc.repo.SaveState(ctx, p)
}