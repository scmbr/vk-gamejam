package usecase

import (
	"context"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{r}
}

func (uc *UserUsecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}

func (uc *UserUsecase) Update(ctx context.Context, u *domain.User) error {
	return uc.userRepo.Create(ctx, u)
}