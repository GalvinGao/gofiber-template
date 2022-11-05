package http

import (
	"go.uber.org/fx"

	"github.com/GalvinGao/gofiber-template/pkg/server/http/route"
)

func Module() fx.Option {
	return fx.Module("http", fx.Provide(Create), route.Module())
}
