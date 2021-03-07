package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	graphql1 "github.com/hantonelli/ghipster/graphql/internal/graphql"
	"github.com/hantonelli/ghipster/graphql/internal/models"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen/review"
)

func (r *mutationResolver) CreateReview(ctx context.Context, input models.CreateReviewInput) (*gen.Review, error) {
	client := gen.FromContext(ctx)
	return client.Review.
		Create().
		SetBody(input.Body).
		SetAuthorID(input.AuthorID).
		SetProductID(input.ProductID).
		Save(ctx)
}

func (r *mutationResolver) UpdateReview(ctx context.Context, input models.UpdateReviewInput) (*gen.Review, error) {
	client := gen.FromContext(ctx)
	return client.Review.UpdateOneID(input.ID).SetBody(input.Body).Save(ctx)
}

func (r *mutationResolver) DeleteReview(ctx context.Context, id int) (bool, error) {
	client := gen.FromContext(ctx)
	err := client.Review.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Review(ctx context.Context, id int) (*gen.Review, error) {
	return r.client.Review.Get(ctx, id)
}

func (r *queryResolver) Reviews(ctx context.Context, filter *models.ReviewFilterInput, orderBy *models.ReviewOrderInput, offset *int, limit int) (*models.ReviewsPayload, error) {
	q := r.client.Review.Query()
	if filter != nil {
		if filter.BodyLike != nil {
			q = q.Where(
				review.BodyContains(*filter.BodyLike),
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
	return &models.ReviewsPayload{
		TotalCount: totalCount,
		Nodes:      nodes,
	}, nil
}

func (r *reviewResolver) Author(ctx context.Context, obj *gen.Review) (*gen.User, error) {
	result, err := obj.Edges.AuthorOrErr()
	if gen.IsNotLoaded(err) {
		result, err = obj.QueryAuthor().Only(ctx)
	}
	return result, err
}

func (r *reviewResolver) Product(ctx context.Context, obj *gen.Review) (*gen.Product, error) {
	result, err := obj.Edges.ProductOrErr()
	if gen.IsNotLoaded(err) {
		result, err = obj.QueryProduct().Only(ctx)
	}
	return result, err
}

func (r *productResolver) Reviews(ctx context.Context, obj *gen.Product) ([]*gen.Review, error) {
	// return r.client.Review.Query().Where(review.HasProductWith(product.ID(obj.ID))).All(ctx)
	result, err := obj.Edges.ReviewsOrErr()
	if gen.IsNotLoaded(err) {
		result, err = obj.QueryReviews().All(ctx)
	}
	return result, err
}

func (r *userResolver) Reviews(ctx context.Context, obj *gen.User) ([]*gen.Review, error) {
	result, err := obj.Edges.ReviewsOrErr()
	if gen.IsNotLoaded(err) {
		result, err = obj.QueryReviews().All(ctx)
	}
	return result, err
}

// Review returns graphql1.ReviewResolver implementation.
func (r *Resolver) Review() graphql1.ReviewResolver { return &reviewResolver{r} }

type reviewResolver struct{ *Resolver }
