package config

import (
	"testing"

	validator "github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSetDefaults(t *testing.T) {
	assert.Panics(t, func() {
		setDefaults(nil)
	})

	v := viper.New()
	assert.NotPanics(t, func() {
		setDefaults(v)
	})

	assert.Equal(t, DefaultWebPort, v.Get("web.port"))
	assert.Equal(t, "info", v.Get("logging.level"))
	assert.Equal(t, "http://localhost:8010/api/inventory/v1", v.Get("clients.host_inventory_base_url"))
	assert.Equal(t, DefaultTokenExpirationTimeSeconds, v.Get("app.token_expiration_seconds"))
	assert.Equal(t, PaginationDefaultLimit, v.Get("app.pagination_default_limit"))
	assert.Equal(t, PaginationMaxLimit, v.Get("app.pagination_max_limit"))
}

func TestLoad(t *testing.T) {
	// 'cfg' is nil panic
	assert.Panics(t, func() {
		Load(nil)
	}, "'cfg' is nil")

	// Success Load
	cfg := Config{}
	assert.NotPanics(t, func() {
		Load(&cfg)
	})
}

func TestValidateConfig(t *testing.T) {
	cfg := Config{}

	err := Validate(&cfg)
	assert.Error(t, err)

	cfg.Application.PathPrefix = DefaultPathPrefix
	ve, ok := err.(validator.ValidationErrors)
	assert.True(t, ok)
	assert.Equal(t, len(ve), 1)
	assert.Equal(t, ve[0].Namespace(), "Config.Application.PathPrefix")
	assert.Equal(t, "required", ve[0].Tag())
}
