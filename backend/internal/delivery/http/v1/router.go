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
		protected.GET("/users/me", h.User.Me)

		// CHILD PROFILE
		protected.POST("/child/profile", h.Child.Create)
		protected.PUT("/child/profile", h.Child.Update)
		protected.GET("/child/profile", h.Child.Get)

		// PET
		protected.GET("/pets/me", h.Pet.GetState)
		protected.PUT("/pets/me", h.Pet.SaveState)
		protected.POST("/pets", h.Pet.Create)
	}
}
