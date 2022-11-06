package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/GalvinGao/gofiber-template/pkg/server/http/errhandle"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errhandle.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: time.Second * 60,
	}))

	return app
}
