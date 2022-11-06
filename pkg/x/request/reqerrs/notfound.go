package reqerrs

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/GalvinGao/gofiber-template/pkg/x/request/reqerrs/predicate"
)

type NotFoundError struct {
	Entity string
}

func NotFoundStr(entity string) NotFoundError {
	return NotFoundError{Entity: entity}
}

func NotFound(pred predicate.Predicate) error {
	return NotFoundStr(pred.Describe())
}

var (
	_ Error = (*NotFoundError)(nil)
	_ error = (*NotFoundError)(nil)
)

func (e NotFoundError) Error() string {
	return e.Message()
}

func (e NotFoundError) StatusCode() int {
	return fiber.ErrBadRequest.Code
}

func (e NotFoundError) ErrorCode() ErrCode {
	return ErrCodeNotFound
}

func (e NotFoundError) Message() string {
	return fmt.Sprintf("cannot find %s", e.Entity)
}

func (e NotFoundError) Fragments() map[string]any {
	return nil
}
