package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/GalvinGao/gofiber-template/internal/app/appconfig"
)

func Configure(conf *appconfig.Config) {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.DurationFieldUnit = time.Nanosecond
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	var writer io.Writer

	if conf.LogJSONStdout {
		writer = os.Stdout
	} else {
		writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339Nano,
		}
	}

	log.Logger = zerolog.New(writer).
		With().
		Timestamp().
		Caller().
		Logger().
		Level(zerolog.Level(conf.LogLevel))
}
