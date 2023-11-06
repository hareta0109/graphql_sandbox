package infra

import (
	"fmt"

	"github.com/hareta0109/graphql_sandbox/internal/lib/config"
	"golang.org/x/xerrors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnector(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort, cfg.DBTimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, xerrors.Errorf("db connection failed: %w", err)
	}

	return db, nil
}
