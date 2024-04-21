package router

import (
	"github.com/avisiedo/go-microservice-1/internal/api/http/private"
	"github.com/avisiedo/go-microservice-1/internal/config"
	handler "github.com/avisiedo/go-microservice-1/internal/handler/http"
	"github.com/labstack/echo/v4"
)

func newPrivate(e *echo.Group, cfg *config.Config, app handler.Application) *echo.Group {
	private.RegisterHandlers(e.Group("/internal"), app)
	return e
}
