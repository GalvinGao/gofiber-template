package controller

import (
	"testing"

	"github.com/gofiber/fiber/v3"
)

var _ fiber.Handler = (&Post{}).GetPosts

func TestRegisterPostRegistersFiberV3Handler(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	RegisterPost(Post{Route: app.Group("/api/v1")})

	for _, registeredRoute := range app.GetRoutes(true) {
		if registeredRoute.Method == fiber.MethodGet && registeredRoute.Path == "/api/v1/posts" {
			return
		}
	}

	t.Fatal("GET /api/v1/posts route was not registered")
}
