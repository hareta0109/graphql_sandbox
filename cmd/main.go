package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hareta0109/graphql_sandbox/internal/adapter/http/handler"
	"github.com/hareta0109/graphql_sandbox/internal/adapter/http/route"
	infra "github.com/hareta0109/graphql_sandbox/internal/infra/repository"
	"github.com/hareta0109/graphql_sandbox/internal/lib/config"
	"github.com/hareta0109/graphql_sandbox/internal/usecase"
)

func main() {
	// Config
	cfg, _ := config.New()

	// DB
	db, _ := infra.NewDBConnector(cfg)

	// DI
	dr := infra.NewDepartmentRepository(db)
	ur := infra.NewUserRepository(db)
	tr := infra.NewTodoRepository(db)
	du := usecase.NewDepartmentRepository(dr)
	uu := usecase.NewUserUsecase(ur)
	tu := usecase.NewTodoUsecase(tr)
	gh := handler.NewGraphHandler(du, uu, tu)
	ph := playground.Handler("GraphQL", "/query")

	// Routing
	r := route.NewInitRoute(gh, ph)
	r.InitRounting(cfg)

	defer func() {
		if dbConn, err := db.DB(); err != nil {
			dbConn.Close()
		}
	}()
}
