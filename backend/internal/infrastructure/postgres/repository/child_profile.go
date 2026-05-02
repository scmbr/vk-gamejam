package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/infrastructure/postgres/models"
	irepo "github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type childProfileRepository struct {
	db *sqlx.DB
}

func NewChildProfileRepository(db *sqlx.DB) irepo.ChildProfileRepository {
	return &childProfileRepository{db}
}

func (r *childProfileRepository) GetByUserID(ctx context.Context, userID int64) (*domain.ChildProfile, error) {
	var m models.ChildProfileModel

	err := r.db.GetContext(ctx, &m,
		`SELECT id, user_id, child_name, child_gender, parent_pin, has_pet
		 FROM child_profiles WHERE user_id=$1`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return mapChildProfileModelToDomain(&m), nil
}

func (r *childProfileRepository) Create(ctx context.Context, p *domain.ChildProfile) error {
	return r.db.QueryRowxContext(ctx, `
		INSERT INTO child_profiles (user_id, child_name, child_gender, parent_pin, has_pet)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING id`,
		p.UserID,
		p.Name,
		p.Gender,
		p.ParentPin,
		p.HasPet,
	).Scan(&p.ID)
}

func (r *childProfileRepository) Update(ctx context.Context, p *domain.ChildProfile) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE child_profiles SET
			child_name=$1,
			child_gender=$2,
			parent_pin=$3,
			has_pet=$4
		WHERE user_id=$5`,
		p.Name,
		p.Gender,
		p.ParentPin,
		p.HasPet,
		p.UserID,
	)

	return err
}