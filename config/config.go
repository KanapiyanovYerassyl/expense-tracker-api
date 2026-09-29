package config

import (
	"fmt"

	"github.com/cloudresty/go-env"
)

type Config struct {
	DBHost     string `env:"DB_HOST" default:"localhost"`
	DBName     string `env:"DB_NAME"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBPort     string `env:"DB_PORT"`
}

func New() (*Config, error) {
	var config Config
	err := env.Bind(&config, env.DefaultBindingOptions())
	if err != nil {
		return nil, fmt.Errorf("failed to load config %w", err)
	}
	return &config, nil
}
