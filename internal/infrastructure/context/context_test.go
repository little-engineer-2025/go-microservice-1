package context

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/avisiedo/go-microservice-1/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestLogFromContext(t *testing.T) {
	var l *slog.Logger

	l = LogFromContext(nil)
	require.NotNil(t, l)

	// Return slog.Default()
	ctx := context.Background()
	l = LogFromContext(ctx)
	require.NotNil(t, l)
	assert.Equal(t, slog.Default(), l)

	ctx = WithLog(ctx, slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})))
	l = LogFromContext(ctx)
	require.NotNil(t, l)
}

func TestDBFromContext(t *testing.T) {
	var (
		db  *gorm.DB
		err error
	)

	db = DBFromContext(nil)
	require.Nil(t, db)

	ctx := context.Background()
	db = DBFromContext(ctx)
	require.Nil(t, db)

	_, db, err = test.NewSqlMock(nil)
	require.NoError(t, err)
	require.NotNil(t, db)
	ctx = WithDB(ctx, db)
	require.NotNil(t, ctx)
	db = DBFromContext(ctx)
	require.NotNil(t, db)
}
