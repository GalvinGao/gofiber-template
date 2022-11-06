package repo

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/GalvinGao/gofiber-template/pkg/model"
	"github.com/GalvinGao/gofiber-template/pkg/x/request/reqerrs"
	"github.com/GalvinGao/gofiber-template/pkg/x/request/reqerrs/predicate"
)

type Post struct {
	db *bun.DB
}

func NewPost(db *bun.DB) *Post {
	return &Post{db: db}
}

func (r *Post) GetPosts(ctx context.Context) ([]model.Post, error) {
	posts := make([]model.Post, 0)

	err := r.db.NewSelect().
		Model(&posts).
		Scan(ctx)

	return posts, err
}

func (r *Post) GetPostByID(ctx context.Context, id int) (*model.Post, error) {
	post := new(model.Post)

	err := r.db.NewSelect().
		Model(post).
		Where("id = ?", id).
		Scan(ctx)

	return post, err
}

func (r *Post) CreatePost(ctx context.Context, post *model.Post) error {
	_, err := r.db.NewInsert().
		Model(post).
		Exec(ctx)

	return err
}

func (r *Post) UpdatePost(ctx context.Context, post *model.Post) error {
	res, err := r.db.NewUpdate().
		Model(post).
		OmitZero().
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return reqerrs.NotFound(predicate.E("post").F("id").IEq(post.ID))
	}

	return nil
}

func (r *Post) DeletePost(ctx context.Context, id int) error {
	res, err := r.db.NewDelete().
		Model((*model.Post)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return reqerrs.NotFound(predicate.E("post").F("id").IEq(id))
	}

	return err
}
