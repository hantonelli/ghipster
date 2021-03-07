package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	graphql1 "github.com/hantonelli/ghipster/graphql/internal/graphql"
	"github.com/hantonelli/ghipster/graphql/internal/models"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*gen.User, error) {
	client := gen.FromContext(ctx)
	return client.User.
		Create().
		SetUsername(input.Username).
		Save(ctx)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UpdateUserInput) (*gen.User, error) {
	client := gen.FromContext(ctx)
	return client.User.UpdateOneID(input.ID).SetUsername(input.Username).Save(ctx)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	client := gen.FromContext(ctx)
	err := client.User.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*gen.User, error) {
	return r.client.User.Get(ctx, id)
}

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
			q = q.Order(gen.Desc(string(orderBy.Field)))
		} else {
			q = q.Order(gen.Asc(string(orderBy.Field)))
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
