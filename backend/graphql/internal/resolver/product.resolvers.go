package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graphql1 "github.com/hantonelli/ghipster/graphql/internal/graphql"
	"github.com/hantonelli/ghipster/graphql/internal/models"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input models.CreateProductInput) (*gen.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, input models.UpdateProductInput) (*gen.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Name(ctx context.Context, obj *gen.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id int) (*gen.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, filter *models.ProductFilterInput, orderBy *models.ProductOrderInput, offset *int, limit int) (*models.ProductsPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

// Product returns graphql1.ProductResolver implementation.
func (r *Resolver) Product() graphql1.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }
