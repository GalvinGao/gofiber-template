package appconfig

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

type ConfigLogLevel zerolog.Level

// ensure ConfigLogLevel implements envconfig.Decoder
var _ envconfig.Decoder = (*ConfigLogLevel)(nil)

func (c *ConfigLogLevel) Decode(value string) error {
	level, err := zerolog.ParseLevel(value)
	if err != nil {
		return err
	}

	*c = ConfigLogLevel(level)
	return nil
}
