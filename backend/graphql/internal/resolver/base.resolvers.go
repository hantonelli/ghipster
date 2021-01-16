package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graphql1 "github.com/hantonelli/ghipster/graphql/internal/graphql"
)

func (r *mutationResolver) Version(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Version(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
