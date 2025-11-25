package utils

import (
	"fmt"
	"os"
)

func GetPostgresDsn(
	host string,
	port int,
	user string,
	password string,
	name string,
	sslmode string,
) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, name, sslmode,
	)
}

func GetMigrationDir() string {
	return fmt.Sprintf("./internal/infra/migrations/%s", os.Getenv("DB_DRIVER"))
}
