package appconfig

import (
	"github.com/GalvinGao/gofiber-template/pkg/app/appcontext"
)

// ConfigSpec is the configuration specification.
type ConfigSpec struct {
	// ServiceListenAddress is the address that the Fiber HTTP server will listen on.
	ServiceListenAddress string `split_words:"true" required:"true" default:":3000"`

	// LogJSONStdout is the flag to enable JSON logging to stdout and disable logging to file.
	LogJSONStdout bool `split_words:"true" required:"true" default:"false"`

	// LogLevel is the log level. Valid values are: trace, debug, info, warn, error, fatal, panic.
	LogLevel ConfigLogLevel `split_words:"true" required:"true" default:"info"`

	// BunPostgresDSN is the DSN for the PostgreSQL database.
	BunPostgresDSN string `split_words:"true" required:"true" default:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`

	// BunDebug is the bun debug mode level. Valid values are: 0 (disabled), 1 (enabled), 2 (enabled with verbose).
	BunDebug int `split_words:"true" required:"true" default:"0"`
}

type Config struct {
	// ConfigSpec is the configuration specification injected to the config.
	ConfigSpec

	// AppContext is the application context
	AppContext appcontext.Ctx
}
