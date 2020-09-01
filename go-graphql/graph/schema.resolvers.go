package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Az3z3l/Ft.-Golang/go-graphql/graph/generated"
	"github.com/Az3z3l/Ft.-Golang/go-graphql/graph/model"
)

func (r *queryResolver) Person(ctx context.Context) ([]*model.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pet(ctx context.Context) ([]*model.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }
