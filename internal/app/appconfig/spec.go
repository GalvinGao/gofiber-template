package appconfig

import (
	"github.com/GalvinGao/gofiber-template/internal/app/appcontext"
)

// ConfigSpec is the configuration specification.
type ConfigSpec struct {
	// ServiceListenAddress is the address that the Fiber HTTP server will listen on.
	ServiceListenAddress string `split_words:"true" required:"true" default:":3000"`

	// LogJSONStdout is the flag to enable JSON logging to stdout and disable logging to file.
	LogJSONStdout bool `split_words:"true" required:"true" default:"false"`

	// LogLevel is the log level. Valid values are: trace, debug, info, warn, error, fatal, panic.
	LogLevel ConfigLogLevel `split_words:"true" required:"true" default:"info"`
}

type Config struct {
	// ConfigSpec is the configuration specification injected to the config.
	ConfigSpec

	// AppContext is the application context
	AppContext appcontext.Ctx
}
