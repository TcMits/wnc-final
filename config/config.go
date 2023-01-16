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
		App                `yaml:"app"`
		HTTP               `yaml:"http"`
		Log                `yaml:"logger"`
		DB                 `yaml:"db"`
		Redis              `yaml:"redis"`
		Mail               `yaml:"mail"`
		AuthUseCase        `yaml:"auth_usecase"`
		TransactionUseCase `yaml:"transaction_usecase"`
		TPBank             `yaml:"tp_bank"`
	}

	// App -.
	App struct {
		Name      string `env-required:"true" env:"APP_NAME"`
		Version   string `yaml:"version"`
		SecretKey string `yaml:"secret_key"`
		Debug     bool   `yaml:"debug" env-required:"true" env:"APP_DEBUG"`
	}

	TPBank struct {
		Name                 string `yaml:"name"`
		ApiKey               string `yaml:"api_key"`
		SecretKey            string `yaml:"secret_key"`
		PrivateKey           string `yaml:"private_key"`
		PublicKey            string `yaml:"public_key"`
		BaseURL              string `yaml:"base_url" env-required:"true" env:"TPBANK_BASE_URL"`
		AuthAPI              string `yaml:"auth_api"`
		BankAccountAPI       string `yaml:"bank_account_api"`
		ValidateAPI          string `yaml:"validate_transaction_api"`
		CreateTransactionAPI string `yaml:"create_transaction_api"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" env:"PORT"`
	}

	// Log -.
	Log struct {
		Level string `yaml:"log_level"`
	}

	// PG -.
	DB struct {
		PoolMax int    `yaml:"pool_max"`
		URL     string `env-required:"true" env:"DB_URL"`
	}

	// Redis
	Redis struct {
		URL      string `env-required:"true"                 env:"REDIS_DB_URL"`
		Password string `env-required:"true"                 env:"REDIS_DB_PASSWORD"`
		DB       int    `env-required:"true"                 env:"REDIS_DB"`
	}

	Mail struct {
		Host                 string        `yaml:"email_host"`
		SenderName           string        `yaml:"email_host_sender_name"`
		Password             string        `yaml:"email_host_password"`
		User                 string        `yaml:"email_host_user"`
		Port                 int           `yaml:"email_host_port"`
		ConfirmEmailTemplate string        `yaml:"confirm_email_template"`
		ConfirmEmailSubject  string        `yaml:"confirm_email_subject"`
		FrontendURL          string        `yaml:"frontend_url"`
		EmailTester          string        `yaml:"your_email"`
		OTPTimeout           time.Duration `yaml:"otp_timeout"`
	}

	// Usecases.
	AuthUseCase struct {
		AccessTTL  time.Duration `yaml:"access_ttl"`
		RefreshTTL time.Duration `yaml:"refresh_ttl"`
	}
	TransactionUseCase struct {
		FeeAmount float64 `yaml:"amount"`
		FeeDesc   string  `yaml:"desc"`
		Layout    string  `yaml:"layout"`
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
