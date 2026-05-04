package usecase

import (
	"context"
	"time"

	"github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type SessionUsecase struct {
	petRepo repository.PetRepository
}

func NewSessionUsecase(petRepo repository.PetRepository) *SessionUsecase {
	return &SessionUsecase{petRepo: petRepo}
}

func (uc *SessionUsecase) Ping(ctx context.Context, userID int64) error {
	return uc.petRepo.UpdateLastOnline(ctx, userID, time.Now().UTC())
}

func (uc *SessionUsecase) End(ctx context.Context, userID int64) error {
	return uc.petRepo.UpdateLastOnline(ctx, userID, time.Now().UTC())
}
