package dto

import "github.com/guregu/null/v6"

type UpdatePostDTO struct {
	Title       null.String `json:"title"`
	Description null.String `json:"description"`
}
