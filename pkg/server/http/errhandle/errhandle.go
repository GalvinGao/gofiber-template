package errhandle

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GalvinGao/gofiber-template/pkg/x/request/reqerrs"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if e, ok := err.(*fiber.Error); ok {
		return handleFiberError(ctx, e)
	}

	if e, ok := err.(reqerrs.Error); ok {
		// custom reqerrs.Error
		return handleReqErr(ctx, e)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"code":    reqerrs.ErrCodeInternal,
		"message": "internal error",
	})
}

func handleFiberError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(fiber.Map{
		"code":    reqerrs.ErrCodeInvalidRequest,
		"message": err.Message,
	})
}

func handleReqErr(ctx *fiber.Ctx, err reqerrs.Error) error {
	m := fiber.Map{
		"code":    err.ErrorCode(),
		"message": err.Message(),
	}
	if f := err.Fragments(); f != nil {
		for k, v := range f {
			m[k] = v
		}
	}
	return ctx.Status(err.StatusCode()).JSON(m)
}
