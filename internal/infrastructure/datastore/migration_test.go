package datastore

import (
	"testing"

	"github.com/avisiedo/go-microservice-1/internal/config"
	"github.com/stretchr/testify/require"
)

func TestCreateMigrationFile(t *testing.T) {
	oldDBMigrationPath := DBMigrationPath()
	SetDBMigrationPath("/tmp-does-not-exist")
	err := CreateMigrationFile("test-demo")
	SetDBMigrationPath(oldDBMigrationPath)
	require.Error(t, err)
	require.Regexp(t, "open /tmp-does-not-exist/([[:digit:]]{14})_test-demo.up.sql: no such file or directory", err.Error())
}

func TestMigrateDb(t *testing.T) {
	err := MigrateDb(nil, "none", 0)
	require.EqualError(t, err, "'cfg' is nil")

	cfg := config.Get()
	oldDBPort := cfg.Database.Port
	cfg.Database.Port = 2345
	err = MigrateDb(cfg, "none", 0)
	cfg.Database.Port = oldDBPort
	require.EqualError(t, err, "could not get database driver: dial tcp 127.0.0.1:2345: connect: connection refused")

	SetDBMigrationPath("/tmp")
	err = MigrateDb(cfg, "none", 0)
	SetDBMigrationPath("")
	require.EqualError(t, err, "'direction' should be 'up' or 'down' but was found 'none'")

	SetDBMigrationPath("/tmp")
	err = MigrateDb(cfg, "up", 0)
	SetDBMigrationPath("")
	require.NoError(t, err)

	SetDBMigrationPath("/tmp")
	err = MigrateDb(cfg, "down", 0)
	SetDBMigrationPath("")
	require.NoError(t, err)
}

func TestMigrateUp(t *testing.T) {
	cfg := config.Get()
	require.NotNil(t, cfg)

	SetDBMigrationPath("/tmp")
	err := MigrateUp(cfg)
	SetDBMigrationPath("")
	require.NoError(t, err)
}

func TestMigrateDown(t *testing.T) {
	cfg := config.Get()
	require.NotNil(t, cfg)

	SetDBMigrationPath("/tmp")
	err := MigrateDown(cfg)
	SetDBMigrationPath("")
	require.NoError(t, err)
}
