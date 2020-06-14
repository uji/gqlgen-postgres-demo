package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-postgres-demo/graph/generated"
	"gqlgen-postgres-demo/graph/model"
	"strconv"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	us, err := r.Resolver.usecase.GetUsers()
	if err != nil {
		return nil, err
	}

	var res []*model.User
	for _, u := range us {
		ts := make([]*model.Todo, len(u.Todos))
		for i, t := range u.Todos {
			ts[i] = &model.Todo{
				ID:   strconv.Itoa(t.ID),
				Done: t.Done,
				Text: t.Text,
			}
		}
		res = append(res, &model.User{
			ID:    strconv.Itoa(u.ID),
			Name:  u.Name,
			Todes: ts,
		})
	}
	return res, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
