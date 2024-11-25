package echo

import (
	"github.com/labstack/echo/v4"
)

// Event presenter for the events resources
type Event interface {
	// List All events
	// (GET /events)
	ListEvents(ctx echo.Context) error
	// Create a Events
	// (POST /events)
	CreateEvent(ctx echo.Context) error
	// Delete a Events
	// (DELETE /events/{eventsId})
	DeleteEvent(ctx echo.Context, eventsId string) error
	// Get a Events
	// (GET /events/{eventsId})
	GetEvent(ctx echo.Context, eventsId string) error
	// Update a Events
	// (PUT /events/{eventsId})
	UpdateEvent(ctx echo.Context, eventsId string) error
}
