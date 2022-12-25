package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
		AuthUseCase
		TransactionUseCase
	}

	// App -.
	App struct {
		Name      string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version   string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		SecretKey string `env-required:"true" yaml:"secret_key" env:"APP_SECRET_KEY"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX" env-default:"90"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}

	// Usecases.
	AuthUseCase struct {
		AccessTTL  time.Duration `env-required:"true" env:"AUTH_ACCESS_TTL"`
		RefreshTTL time.Duration `env-required:"true" env:"AUTH_REFRESH_TTL"`
	}
	TransactionUseCase struct {
		FeeAmount float64 `env-required:"true" env:"FEE_AMOUNT"`
		FeeDesc   string  `env-required:"true" env:"FEE_DESC"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(b), "config.yml")
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
