package common

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNil(t *testing.T) {
	e := NewNil("ctx")
	require.EqualError(t, e, "'ctx' is nil")

	assert.Panic(t, func() {
		e = NewNil("")
	}, "'fieldName' is empty")
}

func TestNewEmpty(t *testing.T) {
	e := NewEmpty("ctx")
	require.EqualError(t, e, "'ctx' is empty")

	assert.Panic(t, func() {
		e = NewEmpty("")
	}, "'fieldName' is empty")
}
