package reqerrs

import "github.com/gofiber/fiber/v2"

type ValidationFieldError struct {
	Field     string `json:"field,omitempty"`
	Violation string `json:"violation"`
	Message   string `json:"message"`
}

type ValidationError struct {
	Fields []ValidationFieldError
}

var (
	_ Error = (*ValidationError)(nil)
	_ error = (*ValidationError)(nil)
)

func (e ValidationError) Error() string {
	return e.Message()
}

func (e ValidationError) StatusCode() int {
	return fiber.ErrBadRequest.Code
}

func (e ValidationError) ErrorCode() ErrCode {
	return ErrCodeValidationFailed
}

func (e ValidationError) Message() string {
	return "validation error"
}

func (e ValidationError) Fragments() map[string]any {
	return map[string]any{"fields": e.Fields}
}
