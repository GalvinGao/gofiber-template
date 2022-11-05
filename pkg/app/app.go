package app

import (
	"go.uber.org/fx"

	"github.com/GalvinGao/gofiber-template/pkg/app/appconfig"
	"github.com/GalvinGao/gofiber-template/pkg/app/appcontext"
	"github.com/GalvinGao/gofiber-template/pkg/controller"
	"github.com/GalvinGao/gofiber-template/pkg/infra"
	"github.com/GalvinGao/gofiber-template/pkg/repo"
	"github.com/GalvinGao/gofiber-template/pkg/server"
	"github.com/GalvinGao/gofiber-template/pkg/service"
	"github.com/GalvinGao/gofiber-template/pkg/x/logger"
	"github.com/GalvinGao/gofiber-template/pkg/x/logger/fxlogger"
)

func New(ctx appcontext.Ctx, additionalOpts ...fx.Option) *fx.App {
	conf, err := appconfig.Parse(ctx)
	if err != nil {
		panic(err)
	}

	// logger and configuration are the only two things that are not in the fx graph
	// because some other packages need them to be initialized before fx starts
	logger.Configure(conf)

	baseOpts := []fx.Option{
		fx.WithLogger(fxlogger.Logger),
		fx.Supply(conf),
		controller.Module(),
		infra.Module(),
		repo.Module(),
		service.Module(),
		server.Module(),
	}

	return fx.New(append(baseOpts, additionalOpts...)...)
}
