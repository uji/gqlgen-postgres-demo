package usecase_test

import (
	"gqlgen-postgres-demo/usecase"
	"testing"
)

type mockUserGetter struct {
	usecase.UserGetter
}

func (m *mockUserGetter) GetUsers() ([]usecase.UserColumns, error) {
	return nil, nil
}

func TestSimple(t *testing.T) {
	getter := &mockUserGetter{}
	user := usecase.NewUser(getter)
	us, err := user.GetUsers()
	if us != nil || err != nil {
		t.Fatal("error")
	}
}
