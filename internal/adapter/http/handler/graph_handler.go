package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/graph-gophers/dataloader"
	"github.com/hareta0109/graphql_sandbox/internal/adapter/http/resolver"
	"github.com/hareta0109/graphql_sandbox/internal/lib/graph/generated"
	"github.com/hareta0109/graphql_sandbox/internal/lib/graph/loader"
	"github.com/hareta0109/graphql_sandbox/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Graph interface {
	QueryHandler() echo.HandlerFunc
}

// usecase を追加
type GraphHandler struct {
	DepartmentUsecase usecase.Department
	UserUsecase       usecase.User
	TodoUsecase       usecase.Todo
}

func NewGraphHandler(
	du usecase.Department,
	uu usecase.User,
	tu usecase.Todo,
) Graph {
	graphHandler := GraphHandler{
		DepartmentUsecase: du,
		UserUsecase:       uu,
		TodoUsecase:       tu,
	}
	return &graphHandler
}

func (g *GraphHandler) QueryHandler() echo.HandlerFunc {
	ldr := &loader.Loaders{
		DepartmentLoader: dataloader.NewBatchedLoader(
			g.DepartmentUsecase.BatchGetDepartments,
			dataloader.WithCache(&dataloader.NoCache{}),
		),
		UserLoader: dataloader.NewBatchedLoader(
			g.UserUsecase.BatchGetUsers,
			dataloader.WithCache(&dataloader.NoCache{}),
		),
	}

	rslvr := resolver.Resolver{
		DepartmentUsecase: g.DepartmentUsecase,
		UserUsecase:       g.UserUsecase,
		TodoUsecase:       g.TodoUsecase,
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &rslvr}),
	)

	return func(c echo.Context) error {
		loader.Middleware(ldr, srv).ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
