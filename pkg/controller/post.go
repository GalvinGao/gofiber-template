package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/GalvinGao/gofiber-template/pkg/model"
	"github.com/GalvinGao/gofiber-template/pkg/model/dto"
	"github.com/GalvinGao/gofiber-template/pkg/service"
	"github.com/GalvinGao/gofiber-template/pkg/x/request"
)

type Post struct {
	fx.In

	PostService *service.Post
	Route       fiber.Router `name:"api-v1"`
}

func RegisterPost(c Post) {
	c.Route.Get("/posts", c.GetPosts)
	c.Route.Put("/posts", c.CreatePost)
	c.Route.Post("/posts/:id", c.UpdatePostByID)
}

func (c *Post) GetPosts(ctx *fiber.Ctx) error {
	posts, err := c.PostService.GetPosts(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(posts)
}

type CreatePostRequest struct {
	Body dto.CreatePost `body:"json" validate:"dive"`
}

func (c *Post) CreatePost(ctx *fiber.Ctx) error {
	req := new(CreatePostRequest)
	if err := request.Inspect(ctx, req); err != nil {
		return err
	}

	post := &model.Post{
		Title:    req.Body.Title,
		Subtitle: req.Body.Subtitle,
		Content:  req.Body.Content,
	}

	if err := c.PostService.CreatePost(ctx.Context(), post); err != nil {
		return err
	}

	return ctx.JSON(post)
}

type UpdatePostByIDRequest struct {
	PostID int            `path:"id" validate:"required"`
	Body   dto.UpdatePost `body:"json" validate:"dive"`
}

func (c *Post) UpdatePostByID(ctx *fiber.Ctx) error {
	req := new(UpdatePostByIDRequest)
	if err := request.Inspect(ctx, req); err != nil {
		return err
	}

	post := &model.Post{
		ID:       req.PostID,
		Title:    req.Body.Title.String,
		Subtitle: req.Body.Subtitle.String,
		Content:  req.Body.Content.String,
	}

	if err := c.PostService.UpdatePost(ctx.Context(), post); err != nil {
		return err
	}

	return ctx.JSON(post)
}
