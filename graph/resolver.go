package graph

import "gqlgen-postgres-demo/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Usecase interface {
	GetUsers() ([]usecase.UserColumns, error)
}

type Resolver struct {
	usecase Usecase
}

func NewResolver(usecase Usecase) *Resolver {
	return &Resolver{usecase}
}
