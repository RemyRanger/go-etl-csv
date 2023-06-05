package logs

import (
	"os"

	"github.com/RemyRanger/go-etl-csv/internal/common/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// New : initialize logger
func New() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Set Log level from config file
	switch logsLevel := config.Conf.Logs.Level; logsLevel {
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("Log level is INFO")
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Msg("Log level is DEBUG")
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		log.Info().Msg("Log level is TRACE")
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("Config empty - Log level default is INFO")
	}
}

// LogInfo : LogInfo
func LogInfo(serviceName, msg string) {
	log.Info().Str("service", serviceName).
		Msg(msg)
}

// LogError : LogError
func LogError(serviceName, msg string, err error) {
	log.Debug().Str("service", serviceName).
		Err(err).
		Msg(msg)
}
