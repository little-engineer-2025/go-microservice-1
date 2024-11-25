package echo

import (
	"github.com/avisiedo/go-microservice-1/internal/config"
	presenter "github.com/avisiedo/go-microservice-1/internal/interface/presenter/echo"
	"github.com/labstack/echo/v4"
)

type eventPresenter struct{}

func NewEvent(cfg *config.Config) presenter.Event {
	if cfg == nil {
		panic("'cfg' is nil")
	}
	return &eventPresenter{}
}

// List All events
// (GET /events)
func (p *eventPresenter) ListEvents(ctx echo.Context) error {
	return echo.ErrNotImplemented
}

// Create a Events
// (POST /events)
func (p *eventPresenter) CreateEvent(ctx echo.Context) error {
	return echo.ErrNotImplemented
}

// Delete a Events
// (DELETE /events/{eventsId})
func (p *eventPresenter) DeleteEvent(ctx echo.Context, eventsId string) error {
	return echo.ErrNotImplemented
}

// Get a Events
// (GET /events/{eventsId})
func (p *eventPresenter) GetEvent(ctx echo.Context, eventsId string) error {
	return echo.ErrNotImplemented
}

// Update a Events
// (PUT /events/{eventsId})
func (p *eventPresenter) UpdateEvent(ctx echo.Context, eventsId string) error {
	return echo.ErrNotImplemented
}
