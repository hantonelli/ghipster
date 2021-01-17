package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen"
	graphql1 "github.com/hantonelli/ghipster/graphql/internal/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *gen.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *gen.Client) graphql.ExecutableSchema {
	return graphql1.NewExecutableSchema(graphql1.Config{
		Resolvers: &Resolver{client},
	})
}
