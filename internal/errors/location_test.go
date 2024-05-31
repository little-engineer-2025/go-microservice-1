package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLocationErrorWithLevel(t *testing.T) {
	e := NewLocationErrorWithLevel(errors.New("test"), 0)
	require.EqualError(t, e, "location_test.go:11 - test")
}

func TestNewLocationError(t *testing.T) {
	e := NewLocationError(errors.New("test"))
	require.EqualError(t, e, "location_test.go:16 - test")
}
