package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

type (
	Config struct {
		Postgres PostgresConfig
		Auth     AuthConfig
		HTTP     HTTPConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	PostgresConfig struct {
		Username string
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string
		SSLMode  string `mapstructure:"sslmode"`
		Password string
	}

	AuthConfig struct {
		JWTSecret       string
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
	}
)

func Init(configsDir string) (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = EnvLocal
	}

	if env == EnvLocal {
		if err := godotenv.Load(".env"); err != nil {
			log.Println("warning: .env file not found, relying on environment variables")
		}
	}

	if err := parseConfigFile(configsDir, env); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return fmt.Errorf("failed to unmarshal postgres config: %w", err)
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return fmt.Errorf("failed to unmarshal http config: %w", err)
	}

	if err := viper.UnmarshalKey("auth", &cfg.Auth); err != nil {
		return fmt.Errorf("failed to unmarshal auth config: %w", err)
	}

	return nil
}

func setFromEnv(cfg *Config) {
	cfg.Postgres.Username = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Name = os.Getenv("POSTGRES_DB")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Auth.JWTSecret = os.Getenv("JWT_SECRET")
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("base")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read main config: %w", err)
	}

	viper.SetConfigName(env)

	if err := viper.MergeInConfig(); err != nil {
		return fmt.Errorf("failed to read %s config: %w", env, err)
	}

	return nil
}
