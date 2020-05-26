package postgres_test

import (
	"gqlgen-postgres-demo/infra/postgres"
	"testing"
)

func TestGetUsers(t *testing.T) {
	db, err := postgres.NewDB()
	if err != nil {
		t.Fatal(err)
	}

	g := postgres.NewUserGetter(db)
	_, err = g.GetUsers()
	if err != nil {
		t.Fatal(err)
	}
}
