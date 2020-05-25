package postgres

import (
	// test
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	return sql.Open("postgres", "dbname=postgres user=postgres sslmode=disable")
}
