package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/GalvinGao/gofiber-template/internal/service"
)

type Post struct {
	fx.In

	PostService *service.Post
	Route       fiber.Router `name:"api-v1"`
}

func RegisterPost(c Post) {
	c.Route.Get("/posts", c.GetPosts)
}

func (c *Post) GetPosts(ctx *fiber.Ctx) error {
	posts, err := c.PostService.GetPosts(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(posts)
}
