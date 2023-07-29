package dto

import "gopkg.in/guregu/null.v4"

type UpdatePostDTO struct {
	Title       null.String `json:"title"`
	Description null.String `json:"description"`
}
