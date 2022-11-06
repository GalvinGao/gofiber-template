package request

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func populateDefault(ctx *fiber.Ctx, f reflect.StructField, v reflect.Value) error {
	if !v.CanSet() {
		return nil
	}

	if v.IsZero() {
		if f.Tag.Get("default") == "" {
			return nil
		}

		if err := convert(f.Tag.Get("default"), v, f.Type); err != nil {
			return err
		}
	}

	return nil
}
