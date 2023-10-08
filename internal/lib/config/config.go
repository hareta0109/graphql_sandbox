package config

import (
	"github.com/caarlos0/env/v6"
	"golang.org/x/xerrors"
)

type Config struct {
	Port       string `env:"API_PORT" envDefault:"8080"`
	DBHost     string `env:"API_DB_HOST" envDefault:"localhost"`
	DBName     string `env:"API_DB_NAME" envDefault:"postgres"`
	DBUser     string `env:"API_DB_USER" envDefault:"postgres"`
	DBPass     string `env:"API_DB_PASS" envDefault:"root"`
	DBPort     string `env:"API_DB_PORT" envDefault:"5432"`
	DBTimeZone string `env:"API_DB_TIMEZONE" envDefault:"Asia/Shanghai"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, xerrors.Errorf("fail to partse cfg: %w", err)
	}
	return cfg, nil
}
