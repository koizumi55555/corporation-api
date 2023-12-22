package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App         `yaml:"app"`
		Logger      `yaml:"logger"`
		QueueConfig `yaml:"queue_config"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// Log -.
	Logger struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// QueueConfig -.
	QueueConfig struct {
		Region   string `env-required:"true" yaml:"region" env:"SQS_REGION"`
		Endpoint string `env-required:"false" yaml:"endpoint" env:"SQS_ENDPOINT"`
		Url      string `env-required:"false" yaml:"url" env:"QUEUE_URL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("github.com/koizumi55555/corporation-api/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
