package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type RouteGroups struct {
	fx.Out

	APIV1    fiber.Router `name:"api-v1"`
	Internal fiber.Router `name:"internal"`
}

func CreateGroups(app *fiber.App) RouteGroups {
	return RouteGroups{
		APIV1:    app.Group("/api/v1"),
		Internal: app.Group("/internal"),
	}
}
