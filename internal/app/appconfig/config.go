package appconfig

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"

	"github.com/GalvinGao/gofiber-template/internal/app/appcontext"
)

func Parse(ctx appcontext.Ctx) (*Config, error) {
	var conf ConfigSpec
	if err := envconfig.Process("app", &conf); err != nil {
		return nil, err
	}

	return &Config{
		ConfigSpec: conf,
		AppContext: ctx,
	}, nil
}
