package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	graphql1 "github.com/hantonelli/ghipster/graphql/internal/graphql"
	"github.com/hantonelli/ghipster/graphql/internal/models"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/entgen"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/entgen/user"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*entgen.User, error) {
	client := entgen.FromContext(ctx)
	return client.User.
		Create().
		SetUsername(input.Username).
		Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UpdateUserInput) (*entgen.User, error) {
	client := entgen.FromContext(ctx)
	return client.User.UpdateOneID(input.ID).SetUsername(input.Username).Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	client := entgen.FromContext(ctx)
	err := client.User.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*entgen.User, error) {
	return r.client.User.Get(ctx, id)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, filter *models.UserFilterInput, orderBy *models.UserOrderInput, offset *int, limit int) (*models.UsersPayload, error) {
	q := r.client.User.Query()
	if filter != nil {
		if filter.UsernameLike != nil {
			q = q.Where(
				user.UsernameContains(*filter.UsernameLike),
			)
		}
	}
	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	if orderBy != nil {
		if *orderBy.Direction == models.OrderDirectionDesc {
			q = q.Order(entgen.Desc(string(orderBy.Field)))
		} else {
			q = q.Order(entgen.Asc(string(orderBy.Field)))
		}
	}
	if offset != nil {
		q = q.Offset(*offset)
	}
	nodes, err := q.Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	return &models.UsersPayload{
		TotalCount: totalCount,
		Nodes:      nodes,
	}, nil
}

// User returns graphql1.UserResolver implementation.
func (r *Resolver) User() graphql1.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
