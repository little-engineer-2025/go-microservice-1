package echo

import (
	"errors"
	"log/slog"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/avisiedo/go-microservice-1/internal/api/http/public"
	"github.com/avisiedo/go-microservice-1/internal/config"
	"github.com/avisiedo/go-microservice-1/internal/domain/model"
	presenter "github.com/avisiedo/go-microservice-1/internal/interface/presenter/sync/echo"
	"github.com/avisiedo/go-microservice-1/internal/test"
	http_builder "github.com/avisiedo/go-microservice-1/internal/test/builder/api/http"
	model_builder "github.com/avisiedo/go-microservice-1/internal/test/builder/model"
	echo_helper "github.com/avisiedo/go-microservice-1/internal/test/helper/http/echo"
	"github.com/avisiedo/go-microservice-1/internal/test/mock/interface/interactor"
	presenter_mock "github.com/avisiedo/go-microservice-1/internal/test/mock/interface/presenter/echo"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNewTodo(t *testing.T) {
	assert.PanicsWithValue(t, ErrPanicCfgIsNil, func() {
		NewTodo(nil, nil, nil)
	})

	cfg := config.Get()
	assert.PanicsWithValue(t, ErrPanicInteractorIsNil, func() {
		NewTodo(cfg, nil, nil)
	})

	i := interactor.NewTodo(t)
	assert.PanicsWithValue(t, ErrPanicDbIsNil, func() {
		NewTodo(cfg, i, nil)
	})

	var p presenter.Todo
	dbMock, db, err := test.NewSqlMock(nil)
	require.NotNil(t, dbMock)
	require.NotNil(t, db)
	require.NoError(t, err)
	assert.NotPanics(t, func() {
		p = NewTodo(cfg, i, db)
	})
	assert.NotNil(t, p)
}

func TestGetAllTodos(t *testing.T) {
	const path = "/todo/v1/todos"
	var (
		db      *gorm.DB
		sqlMock sqlmock.Sqlmock
		err     error
	)
	// inbound error
	cfg := config.Get()
	i := interactor.NewTodo(t)
	inputMock := presenter_mock.NewTodoInput(t)
	outputMock := presenter_mock.NewTodoOutput(t)
	if sqlMock, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
		require.NoError(t, err)
	}
	sqlMock.ExpectBegin()
	sqlMock.ExpectRollback()
	expectedErr := "inbound data error"
	inputMock.On("GetAll", mock.Anything).Return(errors.New(expectedErr))
	e := echo.New()
	require.NotNil(t, e)
	p := newTodo(cfg, inputMock, outputMock, i, db)
	require.NotNil(t, p)
	ctx := echo_helper.NewContext(e, http.MethodGet, path, http.Header{}, nil, slog.Default())
	require.NotNil(t, ctx)
	err = p.GetAllTodos(ctx)
	assert.EqualError(t, err, expectedErr)

	// interactor error
	cfg = config.Get()
	i = interactor.NewTodo(t)
	inputMock = presenter_mock.NewTodoInput(t)
	outputMock = presenter_mock.NewTodoOutput(t)
	if sqlMock, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
		require.NoError(t, err)
	}
	sqlMock.ExpectBegin()
	sqlMock.ExpectRollback()
	expectedErr = "interactor error"
	inputMock.On("GetAll", mock.Anything).Return(nil)
	i.On("GetAll", mock.Anything).Return(nil, errors.New(expectedErr))
	p = newTodo(cfg, inputMock, outputMock, i, db)
	ctx = echo_helper.NewContext(e, http.MethodGet, path, http.Header{}, nil, slog.Default())
	require.NotNil(t, ctx)
	err = p.GetAllTodos(ctx)
	assert.EqualError(t, err, expectedErr)

	// outbound error
	cfg = config.Get()
	i = interactor.NewTodo(t)
	inputMock = presenter_mock.NewTodoInput(t)
	outputMock = presenter_mock.NewTodoOutput(t)
	if sqlMock, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
		require.NoError(t, err)
	}
	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()
	expectedErr = "outbound error"
	inputMock.On("GetAll", mock.Anything).Return(nil)
	i.On("GetAll", mock.Anything).Return([]model.Todo{}, nil)
	outputMock.On("GetAll", mock.Anything, []model.Todo{}).Return(nil, errors.New(expectedErr))
	p = newTodo(cfg, inputMock, outputMock, i, db)
	ctx = echo_helper.NewContext(e, http.MethodGet, path, http.Header{}, nil, slog.Default())
	require.NotNil(t, ctx)
	err = p.GetAllTodos(ctx)
	assert.EqualError(t, err, expectedErr)

	// Success case
	cfg = config.Get()
	i = interactor.NewTodo(t)
	inputMock = presenter_mock.NewTodoInput(t)
	outputMock = presenter_mock.NewTodoOutput(t)
	if sqlMock, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
		require.NoError(t, err)
	}
	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()
	inputMock.On("GetAll", mock.Anything).Return(nil)
	i.On("GetAll", mock.Anything).Return([]model.Todo{}, nil)
	outputMock.On("GetAll", mock.Anything, []model.Todo{}).Return([]public.ToDo{}, nil)
	p = newTodo(cfg, inputMock, outputMock, i, db)
	ctx = echo_helper.NewContext(e, http.MethodGet, path, http.Header{}, nil, slog.Default())
	require.NotNil(t, ctx)
	err = p.GetAllTodos(ctx)
	assert.NoError(t, err)
}

func TestCreateTodo(t *testing.T) {
	var (
		db        *gorm.DB
		sqlMock   sqlmock.Sqlmock
		err       error
		apiInData *public.ToDo
		data      *model.Todo
		// apiOutData *public.ToDo
	)

	const (
		path = "/todo/v1/todos"
	)

	// inbound error
	cfg := config.Get()
	i := interactor.NewTodo(t)
	inputMock := presenter_mock.NewTodoInput(t)
	outputMock := presenter_mock.NewTodoOutput(t)
	// if sqlMock, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
	// 	require.NoError(t, err)
	// }
	if _, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
		require.NoError(t, err)
	}
	// sqlMock.ExpectBegin()
	// sqlMock.ExpectRollback()
	expectedErr := "inbound data error"
	inputMock.On("Create", mock.Anything).Return(nil, errors.New(expectedErr))
	e := echo.New()
	require.NotNil(t, e)
	p := newTodo(cfg, inputMock, outputMock, i, db)
	require.NotNil(t, p)
	apiInData = http_builder.NewToDo().Build()
	ctx := echo_helper.NewContext(e, http.MethodPost, path, http.Header{}, apiInData, slog.Default())
	require.NotNil(t, ctx)
	err = p.CreateTodo(ctx)
	assert.EqualError(t, err, expectedErr)

	// inbound empty error
	cfg = config.Get()
	i = interactor.NewTodo(t)
	inputMock = presenter_mock.NewTodoInput(t)
	outputMock = presenter_mock.NewTodoOutput(t)
	// if sqlMock, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
	// 	require.NoError(t, err)
	// }
	if _, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
		require.NoError(t, err)
	}
	// sqlMock.ExpectBegin()
	// sqlMock.ExpectRollback()
	expectedErr = "empty todo data"
	inputMock.On("Create", mock.Anything).Return(nil, nil)
	e = echo.New()
	require.NotNil(t, e)
	p = newTodo(cfg, inputMock, outputMock, i, db)
	require.NotNil(t, p)
	apiInData = http_builder.NewToDo().Build()
	ctx = echo_helper.NewContext(e, http.MethodPost, path, http.Header{}, apiInData, slog.Default())
	require.NotNil(t, ctx)
	err = p.CreateTodo(ctx)
	assert.EqualError(t, err, echo.NewHTTPError(http.StatusBadRequest, expectedErr).Error())

	// title is an empty string
	cfg = config.Get()
	i = interactor.NewTodo(t)
	inputMock = presenter_mock.NewTodoInput(t)
	outputMock = presenter_mock.NewTodoOutput(t)
	// if sqlMock, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
	// 	require.NoError(t, err)
	// }
	if _, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
		require.NoError(t, err)
	}
	// sqlMock.ExpectBegin()
	// sqlMock.ExpectRollback()
	expectedErr = "title is an empty string"
	apiInData = http_builder.NewToDo().WithTitle("").Build()
	data = model_builder.NewTodo().WithTitle("").Build()
	inputMock.On("Create", mock.Anything).Return(data, nil)
	e = echo.New()
	require.NotNil(t, e)
	p = newTodo(cfg, inputMock, outputMock, i, db)
	require.NotNil(t, p)
	ctx = echo_helper.NewContext(e, http.MethodPost, path, http.Header{}, apiInData, slog.Default())
	require.NotNil(t, ctx)
	err = p.CreateTodo(ctx)
	assert.EqualError(t, err, echo.NewHTTPError(http.StatusBadRequest, expectedErr).Error())

	// description is an empty string
	cfg = config.Get()
	i = interactor.NewTodo(t)
	inputMock = presenter_mock.NewTodoInput(t)
	outputMock = presenter_mock.NewTodoOutput(t)
	// if sqlMock, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
	// 	require.NoError(t, err)
	// }
	if _, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
		require.NoError(t, err)
	}
	// sqlMock.ExpectBegin()
	// sqlMock.ExpectRollback()
	expectedErr = "title is an empty string"
	apiInData = http_builder.NewToDo().WithTitle("title").WithDescription("").Build()
	data = model_builder.NewTodo().WithTitle("title").WithDescription("").Build()
	inputMock.On("Create", mock.Anything).Return(data, nil)
	e = echo.New()
	require.NotNil(t, e)
	p = newTodo(cfg, inputMock, outputMock, i, db)
	require.NotNil(t, p)
	ctx = echo_helper.NewContext(e, http.MethodPost, path, http.Header{}, apiInData, slog.Default())
	require.NotNil(t, ctx)
	err = p.CreateTodo(ctx)
	assert.EqualError(t, err, echo.NewHTTPError(http.StatusBadRequest, expectedErr).Error())

	// interactor error
	cfg = config.Get()
	i = interactor.NewTodo(t)
	inputMock = presenter_mock.NewTodoInput(t)
	outputMock = presenter_mock.NewTodoOutput(t)
	if sqlMock, db, err = test.NewSqlMock(&gorm.Session{}); err != nil {
		require.NoError(t, err)
	}
	sqlMock.ExpectBegin()
	sqlMock.ExpectRollback()
	expectedErr = "title is an empty string"
	apiInData = http_builder.NewToDo().WithTitle("title").WithDescription("description").Build()
	data = model_builder.NewTodo().WithTitle("title").WithDescription("description").Build()
	inputMock.On("Create", mock.Anything).Return(data, nil)
	e = echo.New()
	require.NotNil(t, e)
	p = newTodo(cfg, inputMock, outputMock, i, db)
	require.NotNil(t, p)
	ctx = echo_helper.NewContext(e, http.MethodPost, path, http.Header{}, apiInData, slog.Default())

	require.NotNil(t, ctx)
	err = p.CreateTodo(ctx)
	assert.EqualError(t, err, echo.NewHTTPError(http.StatusBadRequest, expectedErr).Error())
}
