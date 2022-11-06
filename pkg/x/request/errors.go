package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/GalvinGao/gofiber-template/pkg/x/request/reqerrs"
	"github.com/GalvinGao/gofiber-template/pkg/x/translator"
)

func convertErrors(ctx *fiber.Ctx, err error) reqerrs.Error {
	errs := err.(validator.ValidationErrors)

	fieldErrs := make([]reqerrs.ValidationFieldError, 0, len(errs))
	t := translator.FromFiberCtx(ctx)
	for _, err := range errs {
		fieldErrs = append(fieldErrs, reqerrs.ValidationFieldError{
			Field:     err.Field(),
			Violation: err.Tag(),
			Message:   err.Translate(t),
		})
	}
	return reqerrs.ValidationError{Fields: fieldErrs}
}
