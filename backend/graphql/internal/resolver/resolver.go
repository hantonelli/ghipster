package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	graphql1 "github.com/hantonelli/ghipster/graphql/internal/graphql"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *gen.Client
	logger *zap.SugaredLogger
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *gen.Client, logger *zap.Logger) graphql.ExecutableSchema {
	return graphql1.NewExecutableSchema(graphql1.Config{
		Resolvers: &Resolver{
			client: client,
			logger: logger.Sugar(),
		},
	})
}
