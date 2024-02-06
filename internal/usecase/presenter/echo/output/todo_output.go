package output

import (
	"errors"

	"github.com/avisiedo/go-microservice-1/internal/api/http/public"
	"github.com/avisiedo/go-microservice-1/internal/domain/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
)

type TodoOutput struct{}

func (o TodoOutput) createGuards(ctx echo.Context, data *model.Todo) error {
	if ctx == nil {
		return errors.New("ctx is nil")
	}
	if data == nil {
		return errors.New("dataOutput is nil")
	}
	return nil
}

func (o TodoOutput) Create(ctx echo.Context, data *model.Todo) (*public.ToDo, error) {
	if err := o.createGuards(ctx, data); err != nil {
		return nil, err
	}
	output := &public.ToDo{}
	uuidTemp := &uuid.UUID{}
	*uuidTemp = data.UUID
	output.TodoId = uuidTemp
	output.Title = data.Title
	output.Description = data.Description
	if data.DueDate == nil {
		output.DueDate = nil
	} else {
		output.DueDate.Time = *data.DueDate
	}
	return output, nil
}

func (o TodoOutput) getAllGuards(ctx echo.Context, data []model.Todo) error {
	if ctx == nil {
		return errors.New("ctx is nil")
	}
	if data == nil {
		return errors.New("dataOutput is nil")
	}
	return nil
}

func (o TodoOutput) GetAll(ctx echo.Context, data []model.Todo) ([]public.ToDo, error) {
	if err := o.getAllGuards(ctx, data); err != nil {
		return []public.ToDo{}, err
	}
	output := make([]public.ToDo, len(data))
	for i := range data {
		output[i].TodoId = &data[i].UUID
		output[i].Title = data[i].Title
		output[i].Description = data[i].Description
		output[i].DueDate = &types.Date{Time: *data[i].DueDate}
	}
	return output, nil
}

func (o TodoOutput) getGuards(ctx echo.Context, data *model.Todo) error {
	if ctx == nil {
		return errors.New("'ctx' is nil")
	}
	if data == nil {
		return errors.New("'data' is nil")
	}
	if (data.UUID == uuid.UUID{}) {
		return errors.New("'data.UUID' is empty")
	}
	return nil
}

func (o TodoOutput) Get(ctx echo.Context, data *model.Todo) (*public.ToDo, error) {
	if err := o.getGuards(ctx, data); err != nil {
		return nil, err
	}
	out := &public.ToDo{}
	outUUID := &uuid.UUID{}
	*outUUID = data.UUID
	out.TodoId = outUUID
	out.Title = data.Title
	out.Description = data.Description
	if data.DueDate != nil {
		dueDate := &types.Date{}
		dueDate.Time = *data.DueDate
		out.DueDate = dueDate
	}
	return out, nil
}
