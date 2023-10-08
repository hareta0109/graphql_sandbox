package route

import (
	"fmt"
	"net/http"

	"github.com/hareta0109/graphql_sandbox/internal/adapter/http/handler"
	"github.com/hareta0109/graphql_sandbox/internal/lib/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

type Route interface {
	InitRounting(*config.Config) (*echo.Echo, error)
}

type InitRoute struct {
	Gh handler.Graph
	Ph http.HandlerFunc
}

func NewInitRoute(gh handler.Graph, ph http.HandlerFunc) *InitRoute {
	initRoute := InitRoute{gh, ph}
	return &initRoute
}

func (i *InitRoute) InitRounting(cfg *config.Config) (*echo.Echo, error) {
	e := echo.New()

	// Route
	e.POST("/query", i.Gh.QueryHandler())
	e.GET("/playground", func(c echo.Context) error {
		i.Ph.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// Start Server
	if err := e.Start(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		return nil, xerrors.Errorf("fail to start port:%s %w", cfg.Port, err)
	}

	return e, nil
}
