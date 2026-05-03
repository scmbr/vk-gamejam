package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/infrastructure/postgres/models"
	irepo "github.com/scmbr/vk-gamejam/backend/internal/repository"
)

type petRepository struct {
	db *sqlx.DB
}

func NewPetRepository(db *sqlx.DB) irepo.PetRepository {
	return &petRepository{db}
}

func (r *petRepository) GetState(ctx context.Context, userID int64) (*domain.Pet, error) {
	var m models.PetModel

	err := r.db.GetContext(ctx, &m, `
		SELECT user_id, name, type, gender,
		       level, xp,
		       knowledge, energy, creativity, sport,
		       last_online
		FROM pets WHERE user_id=$1`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return mapPetModelToDomain(&m), nil
}

func (r *petRepository) SaveState(ctx context.Context, p *domain.Pet) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO pets (
			user_id, name, type, gender,
			level, xp,
			knowledge, energy, creativity, sport,
			last_online
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		ON CONFLICT (user_id) DO UPDATE SET
			level = EXCLUDED.level,
			xp = EXCLUDED.xp,
			knowledge = EXCLUDED.knowledge,
			energy = EXCLUDED.energy,
			creativity = EXCLUDED.creativity,
			sport = EXCLUDED.sport,
			last_online = EXCLUDED.last_online
	`,
		p.UserID,
		p.Name,
		p.Type,
		p.Gender,
		p.Level,
		p.XP,
		p.Knowledge,
		p.Energy,
		p.Creativity,
		p.Sport,
		p.LastOnline,
	)

	return err
}
func (r *petRepository) Create(ctx context.Context, p *domain.Pet) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO pets (
			user_id, name, type, gender,
			level, xp,
			knowledge, energy, creativity, sport,
			last_online
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`,
		p.UserID,
		p.Name,
		p.Type,
		p.Gender,
		p.Level,
		p.XP,
		p.Knowledge,
		p.Energy,
		p.Creativity,
		p.Sport,
		p.LastOnline,
	)

	return err
}
