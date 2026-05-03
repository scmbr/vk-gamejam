package repository

import (
	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/infrastructure/postgres/models"
)

func mapUserModelToDomain(m *models.UserModel) *domain.User {
	return &domain.User{
		ID:           m.ID,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		CreatedAt:    m.CreatedAt,
	}
}

func mapChildProfileModelToDomain(m *models.ChildProfileModel) *domain.ChildProfile {
	var pin *string
	if m.ParentPin.Valid {
		pin = &m.ParentPin.String
	}

	return &domain.ChildProfile{
		ID:            m.ID,
		UserID:        m.UserID,
		Name:          m.Name,
		Gender:        m.Gender,
		ParentPin:     pin,
		HasPet:        m.HasPet,
		IsFirstLaunch: m.IsFirstLaunch,
		LastLogin:     m.LastLogin,
		LastLogout:    m.LastLogout,
	}
}
func mapPetModelToDomain(m *models.PetModel) *domain.Pet {
	return &domain.Pet{
		UserID:     m.UserID,
		Name:       m.Name,
		Type:       m.Type,
		Gender:     m.Gender,
		Level:      m.Level,
		XP:         m.XP,
		Knowledge:  m.Knowledge,
		Energy:     m.Energy,
		Creativity: m.Creativity,
		Sport:      m.Sport,
		LastOnline: m.LastOnline,
	}
}
