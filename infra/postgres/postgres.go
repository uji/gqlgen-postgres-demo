package postgres

import (
	// test
	"database/sql"
	"fmt"
	"gqlgen-postgres-demo/config"
)

func NewDB() (*sql.DB, error) {
	opt := fmt.Sprintf(
		"host=%s dbname=%s port=%s user=%s password=%s sslmode=disable",
		config.DBHost,
		config.DBName,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
	)
	return sql.Open("postgres", opt)
}
