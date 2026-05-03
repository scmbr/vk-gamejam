package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/repository"
	"github.com/scmbr/vk-gamejam/backend/pkg/hasher"
	"github.com/scmbr/vk-gamejam/backend/pkg/jwt"
)

type AuthUsecase struct {
	userRepo repository.UserRepository
	authRepo repository.AuthRepository

	jwtSecret string
	accessTTL time.Duration
	refreshTTL time.Duration
}

func NewAuthUsecase(
	userRepo repository.UserRepository,
	authRepo repository.AuthRepository,
	jwtSecret string,
	accessTTL, refreshTTL time.Duration,
) *AuthUsecase {
	return &AuthUsecase{
		userRepo:  userRepo,
		authRepo:  authRepo,
		jwtSecret: jwtSecret,
		accessTTL: accessTTL,
		refreshTTL: refreshTTL,
	}
}
func (uc *AuthUsecase) Register(ctx context.Context, email, rawPassword string) (access string, refresh string, err error) {
	_, err = uc.userRepo.GetByEmail(ctx, email)
	if err == nil {
		return "", "", ErrUserAlreadyExists
	}
	passwordHash, err := hasher.Hash(rawPassword)
	if err != nil {
		return "", "", err
	}
	user := &domain.User{
		Email:        email,
		PasswordHash: passwordHash,

	}

	err = uc.userRepo.Create(ctx, user)
	if err != nil {
		return "", "", err
	}



	access, err = jwt.Generate(user.ID, uc.jwtSecret, uc.accessTTL)
	if err != nil {
		return "", "", err
	}

	refresh = uuid.New().String()
	hash := hashToken(refresh)

	err = uc.authRepo.SaveRefreshToken(ctx, user.ID, hash, time.Now().Add(uc.refreshTTL))
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (uc *AuthUsecase) Login(ctx context.Context, email, rawPassword string) (access string, refresh string, err error) {
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	
	err = hasher.Compare(user.PasswordHash, rawPassword)
	if err != nil {
		return "", "", ErrInvalidCredentials
	}
	access, err = jwt.Generate(user.ID, uc.jwtSecret, uc.accessTTL)
	if err != nil {
		return "", "", err
	}

	refresh = uuid.New().String()

	hash := hashToken(refresh)

	err = uc.authRepo.SaveRefreshToken(ctx, user.ID, hash, time.Now().Add(uc.refreshTTL))
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}


func (uc *AuthUsecase) Refresh(ctx context.Context, refreshToken string) (string, string, error) {
	hash := hashToken(refreshToken)

	userID, err := uc.authRepo.GetUserIDByRefresh(ctx, hash)
	if err != nil {
		return "", "", err
	}


	if err := uc.authRepo.DeleteRefreshToken(ctx, hash); err != nil {
		return "", "", err
	}

	newRefresh := uuid.New().String()
	newHash := hashToken(newRefresh)

	err = uc.authRepo.SaveRefreshToken(ctx, userID, newHash, time.Now().Add(uc.refreshTTL))
	if err != nil {
		return "", "", err
	}

	access, err := jwt.Generate(userID, uc.jwtSecret, uc.accessTTL)
	if err != nil {
		return "", "", err
	}

	return access, newRefresh, nil
}

func (uc *AuthUsecase) Logout(ctx context.Context, refreshToken string) error {
	hash := hashToken(refreshToken)
	return uc.authRepo.DeleteRefreshToken(ctx, hash)
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}