package echo

import (
	"testing"

	"github.com/avisiedo/go-microservice-1/internal/config"
	presenter "github.com/avisiedo/go-microservice-1/internal/interface/presenter/echo"
	"github.com/avisiedo/go-microservice-1/internal/test"
	"github.com/avisiedo/go-microservice-1/internal/test/mock/interface/interactor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTodo(t *testing.T) {
	assert.PanicsWithValue(t, "'cfg' is nil", func() {
		NewTodo(nil, nil, nil)
	})

	cfg := config.Get()
	assert.PanicsWithValue(t, "interactor is nil", func() {
		NewTodo(cfg, nil, nil)
	})

	i := interactor.NewTodo(t)
	assert.PanicsWithValue(t, "'db' is nil", func() {
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
