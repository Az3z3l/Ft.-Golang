package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Az3z3l/Ft.-Golang/go-graphql/graph/generated"
	"github.com/Az3z3l/Ft.-Golang/go-graphql/graph/model"
)

func (r *queryResolver) Person(ctx context.Context) ([]*model.Person, error) {
	// panic(fmt.Errorf("not implemented"))
	var result []*model.Person
	Adam := &model.Person{
		ID:   1,
		Name: "Adam",
		Pet: &model.Pet{
			ID:   1,
			Name: "Stew",
		},
	}
	result = append(result, Adam)
	return result, nil
}

func (r *queryResolver) Pet(ctx context.Context) ([]*model.Pet, error) {
	var result []*model.Pet
	Stew := &model.Pet{
		ID:   1,
		Name: "Stew",
	}
	result = append(result, Stew)
	return result, nil
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "HEll0 W0R1d", nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
