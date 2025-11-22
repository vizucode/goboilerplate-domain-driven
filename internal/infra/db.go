package infra

import (
	"database/sql"
	"fmt"
	"time"
)

type DBConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	SSLMode         string
	MaxIdlePool     int
	MaxOpenPool     int
	MaxIdleTime     time.Duration
	MaxPoolLifetime time.Duration
}

func NewPostgresDB(cfg DBConfig) (resp *sql.DB, err error) {

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed open db: %w", err)
	}

	db.SetMaxIdleConns(cfg.MaxIdlePool)
	db.SetMaxOpenConns(cfg.MaxOpenPool)
	db.SetConnMaxLifetime(cfg.MaxPoolLifetime)
	db.SetConnMaxIdleTime(cfg.MaxIdleTime)

	return resp, nil
}
