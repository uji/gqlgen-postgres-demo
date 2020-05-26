package graph_test

import (
	"context"
	"gqlgen-postgres-demo/graph"
	"gqlgen-postgres-demo/usecase"
	"testing"
)

type mockUsecase struct {
	graph.Usecase
}

func (m *mockUsecase) GetUsers() ([]usecase.UserColumns, error) {
	return nil, nil
}

func TestSimple(t *testing.T) {
	m := &mockUsecase{}
	r := graph.NewResolver(m)
	us, err := r.Query().Users(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if us != nil {
		t.Fatal(us)
	}
}
