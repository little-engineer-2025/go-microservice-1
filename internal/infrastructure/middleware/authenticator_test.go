package middleware

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAuthenticator(t *testing.T) {
	m := NewAuthenticator()
	require.NotNil(t, m)

	// Currently it is not used; update with your changes
	err := m(nil, nil)
	require.NoError(t, err)
}
