package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/vk-gamejam/backend/internal/usecase"
)

type SessionHandler struct {
	uc *usecase.SessionUsecase
}

func NewSessionHandler(uc *usecase.SessionUsecase) *SessionHandler {
	return &SessionHandler{uc: uc}
}

func (h *SessionHandler) Ping(c *gin.Context) {
	userID := c.GetInt64("userID")

	if err := h.uc.Ping(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *SessionHandler) End(c *gin.Context) {
	userID := c.GetInt64("userID")

	if err := h.uc.End(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
