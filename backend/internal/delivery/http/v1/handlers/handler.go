package handlers

import (
	"github.com/scmbr/vk-gamejam/backend/internal/usecase"
)

type Handler struct {
	Auth  *AuthHandler
	User  *UserHandler
	Child *ChildProfileHandler
	Pet   *PetHandler
}

func NewHandler(
	authUC *usecase.AuthUsecase,
	userUC *usecase.UserUsecase,
	childUC *usecase.ChildProfileUsecase,
	petUC *usecase.PetUsecase,
) *Handler {

	return &Handler{
		Auth:  NewAuthHandler(authUC),
		User:  NewUserHandler(userUC, childUC),
		Child: NewChildProfileHandler(childUC),
		Pet:   NewPetHandler(petUC),
	}
}
