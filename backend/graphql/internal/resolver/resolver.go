package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/hantonelli/ghipster/graphql/internal/directives"
	graphql1 "github.com/hantonelli/ghipster/graphql/internal/graphql"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/entgen"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *entgen.Client
	logger *zap.SugaredLogger
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *entgen.Client, logger *zap.Logger) graphql.ExecutableSchema {
	return graphql1.NewExecutableSchema(graphql1.Config{
		Resolvers: &Resolver{
			client: client,
			logger: logger.Sugar(),
		},
		Directives: graphql1.DirectiveRoot{
			Binding: directives.Binding,
		},
	})
}
