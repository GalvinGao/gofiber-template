package service

import (
	"context"

	"github.com/GalvinGao/gofiber-template/pkg/model"
	"github.com/GalvinGao/gofiber-template/pkg/repo"
)

type Post struct {
	postRepo *repo.Post
}

func NewPost(postRepo *repo.Post) *Post {
	return &Post{
		postRepo: postRepo,
	}
}

func (s *Post) GetPosts(ctx context.Context) ([]model.Post, error) {
	return s.postRepo.GetPosts(ctx)
}

func (s *Post) GetPostByID(ctx context.Context, id int) (*model.Post, error) {
	return s.postRepo.GetPostByID(ctx, id)
}

func (s *Post) CreatePost(ctx context.Context, post *model.Post) error {
	return s.postRepo.CreatePost(ctx, post)
}

func (s *Post) UpdatePost(ctx context.Context, post *model.Post) error {
	return s.postRepo.UpdatePost(ctx, post)
}

func (s *Post) DeletePost(ctx context.Context, id int) error {
	return s.postRepo.DeletePost(ctx, id)
}
