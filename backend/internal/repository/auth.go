package repository

import (
	"context"
	"time"
)

type AuthRepository interface {
	SaveRefreshToken(ctx context.Context, userID int64, tokenHash string, exp time.Time) error
	GetUserIDByRefresh(ctx context.Context, tokenHash string) (int64, error)
	DeleteRefreshToken(ctx context.Context, tokenHash string) error
}