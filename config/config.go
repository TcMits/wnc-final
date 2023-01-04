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
		PG   `yaml:"pg"`
		Mail `yaml:"mail"`
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
		Port string `env-required:"true" yaml:"port" env:"PORT"`
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

	Mail struct {
		Host                 string        `env-required:"true"                 env:"EMAIL_HOST"`
		SenderName           string        `env-required:"true"                 env:"EMAIL_HOST_SENDER_NAME"`
		Password             string        `env-required:"true"                 env:"EMAIL_HOST_PASSWORD"`
		User                 string        `env-required:"true"                 env:"EMAIL_HOST_USER"`
		Port                 int           `env-required:"true"                 env:"EMAIL_HOST_PORT"`
		ConfirmEmailTemplate string        `yaml:"confirm_email_template" env:"EMAIL_CONFIRM_TEMPLATE"`
		ConfirmEmailSubject  string        `yaml:"confirm_email_subject" env:"EMAIL_CONFIRM_SUBJECT"`
		FrontendURL          string        `yaml:"frontend_url" env:"EMAIL_FRONTEND_URL"`
		OTPTimeout           time.Duration `yaml:"otp_timeout" env:"EMAIL_OTP_TIMEOUT"`
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

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "..")
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewConfigForTest() (*Config, error) {
	cfg := &Config{}
	path := filepath.Join(GetRootDir(), "/config/config.yml")
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	postConfig(cfg)

	return cfg, nil
}

func postConfig(c *Config) {
	path := filepath.Join(GetRootDir(), c.ConfirmEmailTemplate)
	c.Mail.ConfirmEmailTemplate = path
}
