package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/infrastructure/postgres/models"
	irepo "github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type activityRepo struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) irepo.ActivityRepository {
	return &activityRepo{db: db}
}

func (r *activityRepo) Create(ctx context.Context, a *domain.Activity) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO activities (
			child_profile_id, type, activity_id, confirmed_by_parent, created_at
		)
		VALUES ($1,$2,$3,$4,NOW())
	`,
		a.ChildProfileID,
		a.Type,
		a.ActivityID,
		a.ConfirmedByParent,
	)

	return err
}

func (r *activityRepo) GetByChildProfileID(ctx context.Context, childProfileID int64) ([]*domain.Activity, error) {
	rows, err := r.db.QueryxContext(ctx, `
		SELECT id, child_profile_id, type, activity_id, confirmed_by_parent, created_at
		FROM activities
		WHERE child_profile_id = $1
		ORDER BY created_at DESC
	`, childProfileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.Activity

	for rows.Next() {
		var m models.Activity

		if err := rows.StructScan(&m); err != nil {
			return nil, err
		}

		result = append(result, mapActivityModelToDomain(&m))
	}

	return result, nil
}
