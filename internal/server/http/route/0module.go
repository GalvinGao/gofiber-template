package route

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("route", fx.Provide(CreateGroups))
}
