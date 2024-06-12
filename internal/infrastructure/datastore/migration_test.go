package datastore

import (
	"testing"

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
