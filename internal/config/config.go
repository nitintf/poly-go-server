package config

import (
	"github.com/friendsofgo/errors"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	DBUser     string `envconfig:"DB_USER" default:"postgres"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"postgres"`
	DBName     string `envconfig:"DB_NAME" default:"gograph"`
	DBPort     string `envconfig:"DB_PORT"`
	LocalDbURL string `envconfig:"API_LOCAL_DB_URL"`
}

type LoggingConfig struct {
	LogFormat string `envconfig:"LOG_FORMAT" default:"text"`
	LogLevel  string `envconfig:"LOG_LEVEL" default:"info"`
}

type MonitoringConfig struct {
	DatadogAgentHost string `envconfig:"DD_AGENT_HOST"`
	SentryDSN        string `envconfig:"SENTRY_DSN"`
}

type Config struct {
	DatabaseConfig
	LoggingConfig
	MonitoringConfig
	ServiceName string `envconfig:"SERVICE_NAME" default:"graph-go"`
	Port        string `envconfig:"SERVICE_PORT" default:"8080"`
}

func New() (*Config, error) {
	cfg := new(Config)

	err := envconfig.Process("go-graph", cfg)

	if err != nil {
		return nil, errors.Wrap(err, "error loading environment variables")
	}

	return cfg, err
}
