package router

import (
	"github.com/gin-gonic/gin"
	"github.com/scmbr/vk-gamejam/backend/internal/delivery/http/v1/handlers"
)

func RegisterRoutes(r *gin.Engine, h *handlers.Handler, authMW gin.HandlerFunc) {

	api := r.Group("/api")

	// ───────── AUTH (public) ─────────
	auth := api.Group("/auth")
	{
		auth.POST("/login", h.Auth.Login)
		auth.POST("/register", h.Auth.Register)
		auth.POST("/refresh", h.Auth.Refresh)
		auth.POST("/logout", h.Auth.Logout)
	}

	// ───────── PROTECTED ─────────
	protected := api.Group("/")
	protected.Use(authMW)
	{
		// USER
		protected.GET("/user/me", h.User.Me)

		// CHILD PROFILE
		protected.POST("/child/profile", h.Child.Create)
		protected.PUT("/child/profile", h.Child.Update)

		// PET
		protected.GET("/pet/state", h.Pet.GetState)
		protected.PUT("/pet/state", h.Pet.SaveState)
		protected.POST("/pet/create", h.Pet.Create)
	}
}
