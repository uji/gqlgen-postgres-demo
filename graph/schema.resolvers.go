package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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
		res = append(res, &model.User{
			ID:   strconv.Itoa(u.ID),
			Name: u.Name,
		})
	}
	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
