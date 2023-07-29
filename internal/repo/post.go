package repo

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/GalvinGao/gofiber-template/internal/model"
)

type Post struct {
	db *bun.DB
}

func NewPost(db *bun.DB) *Post {
	return &Post{db: db}
}

func (r *Post) GetPosts(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post

	err := r.db.NewSelect().
		Model(&posts).
		Scan(ctx)

	return posts, err
}
