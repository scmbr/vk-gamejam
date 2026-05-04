package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scmbr/vk-gamejam/backend/internal/delivery/http/v1/dto"
	"github.com/scmbr/vk-gamejam/backend/internal/usecase"
)

type AuthHandler struct {
	uc *usecase.AuthUsecase
}

func NewAuthHandler(uc *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	access, refresh, err := h.uc.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	setRefreshCookie(c, refresh)

	c.JSON(http.StatusOK, gin.H{
		"accessToken": access,
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	access, refresh, err := h.uc.Register(
		c.Request.Context(),
		req.Email,
		req.Password,
	)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	setRefreshCookie(c, refresh)

	c.JSON(200, gin.H{
		"accessToken": access,
	})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(401, gin.H{"error": "missing refresh token"})
		return
	}

	access, newRefresh, err := h.uc.Refresh(c.Request.Context(), refreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid refresh token"})
		return
	}

	setRefreshCookie(c, newRefresh)

	c.JSON(200, gin.H{
		"accessToken": access,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, _ := c.Cookie("refreshToken")

	_ = h.uc.Logout(c.Request.Context(), refreshToken)

	// удалить cookie
	c.SetCookie(
		"refreshToken",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	c.Status(200)
}
func setRefreshCookie(c *gin.Context, token string) {
	c.SetCookie(
		"refreshToken",
		token,
		60*60*24*30, // 30 дней
		"/",
		"",
		false, // secure (true в prod)
		true,  // httpOnly
	)
}
