package echo

import (
	"net/http"

	api_healthcheck "github.com/avisiedo/go-microservice-1/internal/api/http/healthcheck"
	"github.com/avisiedo/go-microservice-1/internal/infrastructure/logger/slogctx"
	interactor "github.com/avisiedo/go-microservice-1/internal/interface/interactor"
	presenter "github.com/avisiedo/go-microservice-1/internal/interface/presenter/echo"
	"github.com/labstack/echo/v4"
)

type healthcheckPresenter struct {
	interactor interactor.HealthcheckInteractor
}

func NewHealthcheck(i interactor.HealthcheckInteractor) presenter.Healthcheck {
	return &healthcheckPresenter{
		interactor: i,
	}
}

// Liveness kubernetes probe endpoint
// (GET /livez)
func (p *healthcheckPresenter) GetLivez(ctx echo.Context) error {
	c := ctx.Request().Context()
	l := slogctx.FromCtx(c)
	if err := p.interactor.IsLive(); err != nil {
		l.ErrorContext(c, err.Error())
		return ctx.String(http.StatusInternalServerError, api_healthcheck.NotHealthy)
	}
	return ctx.String(http.StatusOK, api_healthcheck.Healthy)
}

// Readiness kubernetes probe endpoint
// (GET /readyz)
func (p *healthcheckPresenter) GetReadyz(ctx echo.Context) error {
	c := ctx.Request().Context()
	l := slogctx.FromCtx(c)
	if err := p.interactor.IsReady(); err != nil {
		l.ErrorContext(c, err.Error())
		return ctx.String(http.StatusInternalServerError, api_healthcheck.NotReady)
	}
	return ctx.String(http.StatusOK, api_healthcheck.Ready)
}
