package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/vk-gamejam/backend/internal/delivery/http/v1/dto"
	"github.com/scmbr/vk-gamejam/backend/internal/usecase"
)

type UserHandler struct {
	uc      *usecase.UserUsecase
	childUC *usecase.ChildProfileUsecase
}

func NewUserHandler(uc *usecase.UserUsecase, childUC *usecase.ChildProfileUsecase) *UserHandler {
	return &UserHandler{uc: uc, childUC: childUC}
}

func (h *UserHandler) Me(c *gin.Context) {
	userID := c.GetInt64("userID")

	user, err := h.uc.GetByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	profile, err := h.childUC.GetByUserID(c.Request.Context(), userID)

	var hasProfile bool

	if err != nil {
		// важно: отличаем not found от реальной ошибки
		if !isNotFound(err) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}
		hasProfile = false
	} else {
		hasProfile = profile != nil
	}

	c.JSON(http.StatusOK, dto.UserMeResponse{
		ID:              user.ID,
		Email:           user.Email,
		CreatedAt:       user.CreatedAt,
		HasChildProfile: hasProfile,
	})
}
func isNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
