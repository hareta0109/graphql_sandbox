package config

import (
	"github.com/caarlos0/env/v6"
	"golang.org/x/xerrors"
)

type Config struct {
	Port string `env:"API_PORT" envDefault:"8080"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, xerrors.Errorf("fail to partse cfg: %w", err)
	}
	return cfg, nil
}
