package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/usecase"
)

type ChildProfileHandler struct {
	uc *usecase.ChildProfileUsecase
}

func NewChildProfileHandler(uc *usecase.ChildProfileUsecase) *ChildProfileHandler {
	return &ChildProfileHandler{uc}
}
func (h *ChildProfileHandler) Create(c *gin.Context) {
	userID := c.GetInt64("userID")

	var req struct {
		Name      string  `json:"name" binding:"required"`
		Gender    string  `json:"gender" binding:"required"`
		ParentPin *string `json:"parentPin"` // nullable
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	profile := &domain.ChildProfile{
		UserID:        userID,
		Name:          req.Name,
		Gender:        req.Gender,
		ParentPin:     req.ParentPin,
		HasPet:        false,
		IsFirstLaunch: true,
	}

	if err := h.uc.Create(c.Request.Context(), profile); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(201)
}

func (h *ChildProfileHandler) Update(c *gin.Context) {
	userID := c.GetInt64("userID")

	var req struct {
		Name      *string `json:"name"`
		Gender    *string `json:"gender"`
		ParentPin *string `json:"parentPin"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.uc.GetByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "profile not found"})
		return
	}

	if req.Name != nil {
		profile.Name = *req.Name
	}

	if req.Gender != nil {
		profile.Gender = *req.Gender
	}

	if req.ParentPin != nil {
		profile.ParentPin = req.ParentPin
	}

	if err := h.uc.Update(c.Request.Context(), profile); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}
