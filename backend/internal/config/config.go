package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	EnvLocal = "local"
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
	if err := parseConfigFile(configsDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}
	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	setFromEnv(&cfg)
	return &cfg, nil
}
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("auth", &cfg.Auth); err != nil {
		return err
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
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)

	return nil
}
