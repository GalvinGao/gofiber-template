package reqerrs

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RequestParseError struct {
	Location string `json:"location,omitempty"`
	Key      string `json:"key,omitempty"`
	Msg      string `json:"message,omitempty"`
}

func Parsing(location, key string, err error) RequestParseError {
	return RequestParseError{
		Location: location,
		Key:      key,
		Msg:      err.Error(),
	}
}

var (
	_ Error = (*RequestParseError)(nil)
	_ error = (*RequestParseError)(nil)
)

func (e RequestParseError) Error() string {
	return e.Message()
}

func (e RequestParseError) StatusCode() int {
	return fiber.ErrBadRequest.Code
}

func (e RequestParseError) ErrorCode() ErrCode {
	return ErrCodeParseFailed
}

func (e RequestParseError) Message() string {
	return fmt.Sprintf("unable to parse parameter `%s` in `%s`: %s", e.Key, e.Location, e.Msg)
}

func (e RequestParseError) Fragments() map[string]any {
	return map[string]any{
		"location": e.Location,
		"key":      e.Key,
	}
}
