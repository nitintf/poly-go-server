package logging

import (
	"github.com/nitintf/graph-go/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func New(cfg *config.Config) *zerolog.Logger {
	logger := log.
		With().
		Caller().
		Str("app", cfg.ServiceName).
		Str("version", "").
		Logger()

	level, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.Warn().Str("log_level_config", cfg.LogLevel).Msg("unable to parse log level")
	}
	logger = logger.Level(level)

	// Set this as default logger and context logger
	zerolog.DefaultContextLogger = &logger
	log.Logger = logger

	return &logger
}
