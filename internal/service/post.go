package service

import (
	"context"

	"github.com/GalvinGao/gofiber-template/internal/model"
	"github.com/GalvinGao/gofiber-template/internal/repo"
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
	posts, err := s.postRepo.GetPosts(ctx)

	return posts, err
}
