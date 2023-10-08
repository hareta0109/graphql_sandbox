package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hareta0109/graphql_sandbox/internal/adapter/http/handler"
	"github.com/hareta0109/graphql_sandbox/internal/adapter/http/route"
	"github.com/hareta0109/graphql_sandbox/internal/lib/config"
)

func main() {
	// Config
	cfg, _ := config.New()

	// DI
	gh := handler.NewGraphHandler()
	ph := playground.Handler("GraphQL", "/query")

	// Routing
	r := route.NewInitRoute(gh, ph)
	r.InitRounting(cfg)

	defer func() {

	}()
}
