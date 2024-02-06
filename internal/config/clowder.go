//go:build ignore

// build ignore
package config

import (
	clowder "github.com/redhatinsights/app-common-go/pkg/api/v1"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

func setClowderConfiguration(v *viper.Viper, clowderConfig *clowder.AppConfig) {
	if !clowder.IsClowderEnabled() {
		return
	}
	if v == nil {
		panic("'v' is nil")
	}
	if clowderConfig == nil {
		panic("'clowderConfig' is nil")
	}

	// Web
	v.Set("web.port", clowderConfig.PublicPort)

	// Database
	var rdsCertPath string
	if clowderConfig.Database != nil && clowderConfig.Database.RdsCa != nil {
		var err error
		if rdsCertPath, err = clowderConfig.RdsCa(); err != nil {
			slog.Warn("Cannot read RDS CA cert", slog.Any("error", err))
		}
	}
	if clowderConfig.Database != nil {
		v.Set("database.host", clowderConfig.Database.Hostname)
		v.Set("database.port", clowderConfig.Database.Port)
		v.Set("database.user", clowderConfig.Database.Username)
		v.Set("database.password", clowderConfig.Database.Password)
		v.Set("database.name", clowderConfig.Database.Name)
		if rdsCertPath != "" {
			v.Set("database.ca_cert_path", rdsCertPath)
		}
	}

	// Clowdwatch
	if clowderConfig.Logging.Cloudwatch != nil {
		v.Set("cloudwatch.region", clowderConfig.Logging.Cloudwatch.Region)
		v.Set("cloudwatch.group", clowderConfig.Logging.Cloudwatch.LogGroup)
		v.Set("cloudwatch.secret", clowderConfig.Logging.Cloudwatch.SecretAccessKey)
		v.Set("cloudwatch.key", clowderConfig.Logging.Cloudwatch.AccessKeyId)
	}

	// Metrics configuration
	v.Set("metrics.path", clowderConfig.MetricsPath)
	v.Set("metrics.port", clowderConfig.MetricsPort)
}
