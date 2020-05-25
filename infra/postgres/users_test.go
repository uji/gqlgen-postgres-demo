package postgres_test

import (
	"gqlgen-postgres-demo/infra/postgres"
	"testing"
)

func TestQueryUsers(t *testing.T) {
	db, err := postgres.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	_, err = postgres.QueryUsers(db)
	if err != nil {
		t.Fatal(err)
	}
}
