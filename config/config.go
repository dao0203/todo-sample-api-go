package config

import "github.com/caarlos0/env/v11"

type (
	// Config -.
	Config struct {
		HTTP HTTP
	}

	// HTTP -.
	HTTP struct {
		Port string `env:"HTTP_PORT", required`
	}
)

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
