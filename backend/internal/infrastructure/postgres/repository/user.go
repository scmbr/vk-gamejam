package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/infrastructure/postgres/models"
	irepo "github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) irepo.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	var m models.UserModel

	err := r.db.GetContext(ctx, &m,
		`SELECT id, email, password_hash, created_at FROM users WHERE id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}

	return mapUserModelToDomain(&m), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var m models.UserModel

	err := r.db.GetContext(ctx, &m,
		`SELECT id, email, password_hash, created_at FROM users WHERE email=$1`,
		email,
	)
	if err != nil {
		return nil, err
	}

	return mapUserModelToDomain(&m), nil
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	return r.db.QueryRowxContext(ctx, `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id, created_at`,
		u.Email,
		u.PasswordHash,
	).Scan(&u.ID, &u.CreatedAt)
}