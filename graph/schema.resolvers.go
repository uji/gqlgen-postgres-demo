package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqlgen-postgres-demo/graph/generated"
	"gqlgen-postgres-demo/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{
		{
			ID:   "user-id",
			Name: "user1",
			Todes: []*model.Todo{
				{
					ID:   "todo1",
					Done: false,
					Text: "todo1",
				},
			},
		},
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
