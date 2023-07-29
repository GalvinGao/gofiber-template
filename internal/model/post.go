package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"posts" swaggerignore:"true"`

	ID        int       `bun:",pk,autoincrement" json:"id"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Content   string    `json:"content"`
}
