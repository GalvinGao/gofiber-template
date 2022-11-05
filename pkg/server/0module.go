package server

import (
	"go.uber.org/fx"

	"github.com/GalvinGao/gofiber-template/pkg/server/http"
)

func Module() fx.Option {
	return fx.Module("server", http.Module())
}
