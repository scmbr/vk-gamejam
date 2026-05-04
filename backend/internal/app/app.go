package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/scmbr/vk-gamejam/backend/internal/config"
	"github.com/scmbr/vk-gamejam/backend/internal/delivery/http/server"
	router "github.com/scmbr/vk-gamejam/backend/internal/delivery/http/v1"
	"github.com/scmbr/vk-gamejam/backend/internal/delivery/http/v1/handlers"
	"github.com/scmbr/vk-gamejam/backend/internal/infrastructure/postgres/client"
	"github.com/scmbr/vk-gamejam/backend/internal/infrastructure/postgres/repository"
	"github.com/scmbr/vk-gamejam/backend/internal/middleware"
	"github.com/scmbr/vk-gamejam/backend/internal/usecase"
	"github.com/scmbr/vk-gamejam/backend/pkg/logger"
)

func Run(configsDir string) {
	logger.Init()
	cfg, err := config.Init(configsDir)
	if err != nil {
		logger.Error("config initialization error", err, nil)
		os.Exit(1)
	}
	db, err := client.NewPostgresDB(client.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.Name,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		logger.Error("database initialization error", err, nil)
		os.Exit(1)
	}
	logger.Info("database connected successfully", nil)

	userRepo := repository.NewUserRepository(db)
	childRepo := repository.NewChildProfileRepository(db)
	petRepo := repository.NewPetRepository(db)
	authRepo := repository.NewAuthRepository(db)
	activityRepo := repository.NewActivityRepository(db)
	userUC := usecase.NewUserUsecase(userRepo, childRepo)
	petUC := usecase.NewPetUsecase(petRepo)
	childUC := usecase.NewChildProfileUsecase(childRepo)
	authUC := usecase.NewAuthUsecase(userRepo, authRepo, cfg.Auth.JWTSecret, cfg.Auth.AccessTokenTTL, cfg.Auth.RefreshTokenTTL)
	activityUC := usecase.NewActivityUsecase(activityRepo)
	sessionUC := usecase.NewSessionUsecase(petRepo)
	handler := handlers.NewHandler(authUC, userUC, childUC, petUC, activityUC, sessionUC)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	authMW := middleware.AuthMiddleware(cfg.Auth.JWTSecret)

	router.RegisterRoutes(r, handler, authMW)

	server := server.NewServer(cfg, r)
	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error("error occurred while running server", err, nil)
		}
	}()

	logger.Info("server started", nil)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		logger.Error("failed to stop server", err, nil)
	}

}
