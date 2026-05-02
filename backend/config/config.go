package config

import (
	"os"
	"time"
)

type Config struct {
	DBUrl           string
	JWTSecret       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func Load() *Config {
	return &Config{
		DBUrl:           getEnv("DATABASE_URL", "postgres://postgres:postgres@postgres:5432/game?sslmode=disable"),
		JWTSecret:       getEnv("JWT_SECRET", "secret"),
		AccessTokenTTL:  getDuration("ACCESS_TOKEN_TTL", 15*time.Minute),
		RefreshTokenTTL: getDuration("REFRESH_TOKEN_TTL", 720*time.Hour),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getDuration(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		panic("invalid duration for " + key)
	}

	return d
}