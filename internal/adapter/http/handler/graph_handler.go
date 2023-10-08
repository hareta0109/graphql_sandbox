package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/hareta0109/graphql_sandbox/internal/adapter/http/resolver"
	"github.com/hareta0109/graphql_sandbox/internal/lib/graph/generated"
	"github.com/hareta0109/graphql_sandbox/internal/lib/graph/loader"
	"github.com/labstack/echo/v4"
)

type Graph interface {
	QueryHandler() echo.HandlerFunc
}

// usecase を追加
type GraphHandler struct{}

func NewGraphHandler() Graph {
	graphHandler := GraphHandler{}
	return &graphHandler
}

func (g *GraphHandler) QueryHandler() echo.HandlerFunc {
	ldr := loader.NewLoaders()

	rslvr := resolver.Resolver{}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &rslvr}),
	)

	return func(c echo.Context) error {
		loader.Middleware(ldr, srv).ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
