package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	irepo "github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) irepo.AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) SaveRefreshToken(ctx context.Context, userID int64, tokenHash string, exp time.Time) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO refresh_tokens (user_id, token_hash, expires_at)
		VALUES ($1,$2,$3)`,
		userID,
		tokenHash,
		exp,
	)
	return err
}

func (r *authRepository) GetUserIDByRefresh(ctx context.Context, tokenHash string) (int64, error) {
	var userID int64

	err := r.db.GetContext(ctx, &userID, `
		SELECT user_id FROM refresh_tokens
		WHERE token_hash=$1 AND expires_at > NOW()`,
		tokenHash,
	)

	return userID, err
}

func (r *authRepository) DeleteRefreshToken(ctx context.Context, tokenHash string) error {
	_, err := r.db.ExecContext(ctx,
		`DELETE FROM refresh_tokens WHERE token_hash=$1`,
		tokenHash,
	)
	return err
}