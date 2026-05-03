package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/vk-gamejam/backend/internal/delivery/http/v1/dto"
	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/usecase"
)

type PetHandler struct {
	uc *usecase.PetUsecase
}

func NewPetHandler(uc *usecase.PetUsecase) *PetHandler {
	return &PetHandler{uc}
}

func (h *PetHandler) GetState(c *gin.Context) {
	userID := c.GetInt64("userID")

	pet, err := h.uc.GetState(c.Request.Context(), userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "pet not found"})
		return
	}

	c.JSON(http.StatusOK, dto.PetResponse{
		PetName:    pet.Name,
		PetType:    pet.Type,
		PetGender:  pet.Gender,
		PetLevel:   pet.Level,
		CurrentXP:  pet.XP,
		Knowledge:  pet.Knowledge,
		Energy:     pet.Energy,
		Creativity: pet.Creativity,
		Sport:      pet.Sport,
		LastOnline: pet.LastOnline,
	})
}

func (h *PetHandler) SaveState(c *gin.Context) {
	userID := c.GetInt64("userID")

	var req dto.UpdatePetStateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pet := &domain.Pet{
		UserID:     userID,
		Level:      req.PetLevel,
		XP:         req.CurrentXP,
		Knowledge:  req.Knowledge,
		Energy:     req.Energy,
		Creativity: req.Creativity,
		Sport:      req.Sport,
		LastOnline: req.LastOnline,
	}

	if err := h.uc.SaveState(c.Request.Context(), pet); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}

func (h *PetHandler) Create(c *gin.Context) {
	userID := c.GetInt64("userID")

	var req dto.CreatePetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pet := &domain.Pet{
		UserID: userID,
		Name:   req.PetName,
		Type:   req.PetType,
		Gender: req.PetGender,
	}

	if err := h.uc.Create(c.Request.Context(), pet); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(201)
}
