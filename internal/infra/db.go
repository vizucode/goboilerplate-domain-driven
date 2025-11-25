package infra

import (
	"database/sql"
	"fmt"
	"goboilerplate-domain-driven/pkg/utils"
	"time"

	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Driver          string
	MaxIdlePool     int
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxOpenPool     int
	MaxIdleTime     time.Duration
	MaxPoolLifetime time.Duration
}

func NewInitDB(cfg DBConfig) (resp *sql.DB, err error) {

	db, err := sql.Open(cfg.Driver, utils.GetPostgresDsn(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("failed open db: %w", err)
	}

	db.SetMaxIdleConns(cfg.MaxIdlePool)
	db.SetMaxOpenConns(cfg.MaxOpenPool)
	db.SetConnMaxLifetime(cfg.MaxPoolLifetime)
	db.SetConnMaxIdleTime(cfg.MaxIdleTime)

	if err := goose.Up(db, utils.GetMigrationDir()); err != nil {
		return nil, fmt.Errorf("failed open migration: %w", err)
	}

	return db, nil
}
