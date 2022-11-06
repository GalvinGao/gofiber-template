package dto

import "gopkg.in/guregu/null.v3"

type UpdatePost struct {
	Title    null.String `json:"title" validate:"max=255"`
	Subtitle null.String `json:"subtitle" validate:"max=255"`
	Content  null.String `json:"content"`
}

type CreatePost struct {
	Title    string `json:"title" validate:"required,max=255"`
	Subtitle string `json:"subtitle" validate:"required,max=255"`
	Content  string `json:"content" validate:"required"`
}
