package logger

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/avisiedo/go-microservice-1/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	assert.PanicsWithValue(t, "'cfg' cannot be nil", func() {
		InitLogger(nil)
	})

	cfg := config.Config{
		Logging: config.Logging{
			Location: true,
			Level:    "info",
			Console:  false,
		},
	}
	assert.NotPanics(t, func() {
		InitLogger(&cfg)
	})

	cfg.Logging.Console = true
	assert.NotPanics(t, func() {
		InitLogger(&cfg)
	})

	levels := []string{"TRACE", "DEBUG", "INFO", "NOTICE", "WARN", "ERROR", ""}
	for _, level := range levels {
		cfg.Logging.Level = level
		InitLogger(&cfg)
	}
}

func TestLogBuildInfo(t *testing.T) {
	b := bytes.Buffer{}
	sh := slog.NewTextHandler(&b, &slog.HandlerOptions{
		Level: LevelInfo,
	})
	sl := slog.New(sh)
	slogDefault := slog.Default()
	defer slog.SetDefault(slogDefault)

	slog.SetDefault(sl)
	LogBuildInfo("test")
	slog.SetDefault(slogDefault)

	// https://regex101.com/
	const regularExp = "^" + timeExp + " " + levelExp + " " + `msg=(.*)\n$`

	assert.Regexp(t, regularExp, b.String())
}
