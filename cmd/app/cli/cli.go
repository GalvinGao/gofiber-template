package cli

import (
	"context"

	"go.uber.org/fx"

	"github.com/GalvinGao/gofiber-template/internal/app"
	"github.com/GalvinGao/gofiber-template/internal/app/appcontext"
)

func Start(module fx.Option) {
	app.New(appcontext.Declare(appcontext.EnvCLI), module).Start(context.Background())
}
