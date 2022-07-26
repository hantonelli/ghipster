package main

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	"github.com/hantonelli/ghipster/graphql/internal/resolver"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen"
	"github.com/hantonelli/ghipster/graphql/internal/service/ent/gen/migrate"
	_ "github.com/hantonelli/ghipster/graphql/internal/service/ent/gen/runtime"
	"github.com/hantonelli/ghipster/middleware"
)

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8080" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	log, _ := zap.NewProduction()
	dbClient, err := gen.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
	}
	if err := dbClient.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", zap.Error(err))
	}

	gqlSrv := handler.NewDefaultServer(resolver.NewSchema(dbClient, log))
	// gqlSrv.Use(entgql.Transactioner{TxOpener: dbClient})
	if cli.Debug {
		gqlSrv.Use(&debug.Tracer{})
	}

	router := chi.NewRouter()
	router.Handle("/ui", playground.Handler("GHipster", "/query"))
	router.Handle("/query", middleware.TracingMiddleware(gqlSrv))

	log.Sugar().Infow("listening on", "address", cli.Addr)
	if err := http.ListenAndServe(cli.Addr, router); err != nil {
		log.Error("http server terminated", zap.Error(err))
	}
}
