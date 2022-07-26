package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	graphql1 "github.com/hantonelli/ghipster/graphql/internal/graphql"
	"github.com/hantonelli/ghipster/graphql/internal/models"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen/product"
	"github.com/hantonelli/ghipster/middleware"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input models.CreateProductInput) (*gen.Product, error) {
	log := middleware.AddTraceToLog(ctx, r.logger)
	log.Infow("Calling CreateProduct")
	client := gen.FromContext(ctx)
	return client.Product.
		Create().
		SetName(input.Name).
		Save(ctx)
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input models.UpdateProductInput) (*gen.Product, error) {
	client := gen.FromContext(ctx)
	return client.Product.UpdateOneID(input.ID).SetName(input.Name).Save(ctx)
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id int) (bool, error) {
	client := gen.FromContext(ctx)
	err := client.Product.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id int) (*gen.Product, error) {
	return r.client.Product.Get(ctx, id)
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, filter *models.ProductFilterInput, orderBy *models.ProductOrderInput, offset *int, limit int) (*models.ProductsPayload, error) {
	q := r.client.Product.Query()
	if filter != nil {
		if filter.NameLike != nil && *filter.NameLike != "" {
			q = q.Where(
				product.NameEQ(*filter.NameLike),
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
	return &models.ProductsPayload{
		TotalCount: totalCount,
		Nodes:      nodes,
	}, nil
}

// Product returns graphql1.ProductResolver implementation.
func (r *Resolver) Product() graphql1.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }
