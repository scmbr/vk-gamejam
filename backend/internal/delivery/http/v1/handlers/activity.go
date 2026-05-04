package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/vk-gamejam/backend/internal/delivery/http/v1/dto"
	"github.com/scmbr/vk-gamejam/backend/internal/domain"
	"github.com/scmbr/vk-gamejam/backend/internal/usecase"
)

type ActivityHandler struct {
	uc      *usecase.ActivityUsecase
	childUC *usecase.ChildProfileUsecase
}

func NewActivityHandler(uc *usecase.ActivityUsecase, childUC *usecase.ChildProfileUsecase) *ActivityHandler {
	return &ActivityHandler{uc: uc, childUC: childUC}
}

func (h *ActivityHandler) Create(c *gin.Context) {
	userID := c.GetInt64("userID")

	var req dto.CreateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	switch req.Type {
	case "reading", "art", "sport":
	default:
		c.JSON(400, gin.H{"error": "invalid activity type"})
		return
	}
	childProfile, err := h.childUC.GetByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "child profile not found"})
		return
	}
	a := &domain.Activity{
		ChildProfileID:    childProfile.ID,
		Type:              domain.ActivityType(req.Type),
		ActivityID:        req.ActivityID,
		ConfirmedByParent: req.ConfirmedByParent,
	}

	if err := h.uc.Create(c.Request.Context(), a); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *ActivityHandler) Get(c *gin.Context) {
	userID := c.GetInt64("userID")
	childProfile, err := h.childUC.GetByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(404, gin.H{"error": "child profile not found"})
		return
	}
	activities, err := h.uc.GetByChildProfileID(c.Request.Context(), childProfile.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp := make([]dto.ActivityResponse, 0, len(activities))

	for _, a := range activities {
		resp = append(resp, dto.ActivityResponse{
			ID:                a.ID,
			Type:              string(a.Type),
			ActivityID:        a.ActivityID,
			ConfirmedByParent: a.ConfirmedByParent,
			CreatedAt:         a.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, resp)
}
