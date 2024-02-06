package middleware

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3filter"
)

// NewAuthenticator create an authenticator for the public API
func NewAuthenticator() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		// TODO Call here to your Authenticate function
		// return Authenticate(v, ctx, input)
		return nil
	}
}
