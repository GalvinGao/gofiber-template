package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Post struct {
	bun.BaseModel `bun:"posts" swaggerignore:"true"`

	ID        int        `bun:",pk,autoincrement" json:"id"`
	CreatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	DeletedAt *time.Time `bun:",soft_delete,nullzero" json:"deleted_at,omitempty"`
	Title     string     `json:"title"`
	Subtitle  string     `json:"subtitle"`
	Content   string     `json:"content"`
}
