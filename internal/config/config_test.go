package config

import (
	"testing"

	validator "github.com/go-playground/validator/v10"
	"github.com/openlyinc/pointy"
	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"
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

func TestHasKafkaBrokerConfig(t *testing.T) {
	assert.False(t, hasKafkaBrokerConfig(nil))
	cfg := clowder.AppConfig{}
	assert.False(t, hasKafkaBrokerConfig(&cfg))
	cfg.Kafka = &clowder.KafkaConfig{}
	assert.False(t, hasKafkaBrokerConfig(&cfg))
	cfg.Kafka.Brokers = []clowder.BrokerConfig{}
	assert.False(t, hasKafkaBrokerConfig(&cfg))
	cfg.Kafka.Brokers = append(cfg.Kafka.Brokers, clowder.BrokerConfig{})
	assert.False(t, hasKafkaBrokerConfig(&cfg))
	cfg.Kafka.Brokers[0].Hostname = "test.kafka.svc.localdomain"
	assert.False(t, hasKafkaBrokerConfig(&cfg))
	cfg.Kafka.Brokers[0].Port = pointy.Int(3000)
	assert.True(t, hasKafkaBrokerConfig(&cfg))
}

func TestAddEventConfigDefaults(t *testing.T) {
	assert.PanicsWithValue(t, "'options' is nil", func() {
		addEventConfigDefaults(nil)
	})

	v := viper.New()
	addEventConfigDefaults(v)
	assert.Equal(t, 10000, v.Get("kafka.timeout"))
	assert.Equal(t, DefaultAppName, v.Get("kafka.group.id"))
	assert.Equal(t, "latest", v.Get("kafka.auto.offset.reset"))
	assert.Equal(t, 5000, v.Get("kafka.auto.commit.interval.ms"))
	assert.Equal(t, -1, v.Get("kafka.request.required.acks"))
	assert.Equal(t, 15, v.Get("kafka.message.send.max.retries"))
	assert.Equal(t, 100, v.Get("kafka.retry.backoff.ms"))

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
