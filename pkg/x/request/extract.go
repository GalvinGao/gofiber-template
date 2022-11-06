package request

import (
	"reflect"

	"github.com/gofiber/fiber/v2"

	"github.com/GalvinGao/gofiber-template/pkg/x/request/reqerrs"
)

func extract(ctx *fiber.Ctx, f reflect.StructField, v reflect.Value) error {
	key, ok := f.Tag.Lookup("path")
	if ok {
		if err := extractPath(ctx, key, v); err != nil {
			return reqerrs.Parsing("path", key, err)
		}
	}

	key, ok = f.Tag.Lookup("query")
	if ok {
		if err := extractQuery(ctx, key, v); err != nil {
			return reqerrs.Parsing("query", key, err)
		}
	}

	key, ok = f.Tag.Lookup("body")
	if ok {
		if key == "json" || key == "form" {
			if err := extractBody(ctx, v); err != nil {
				return reqerrs.Parsing("body", key, err)
			}
		}
	}
	return nil
}

func extractPath(ctx *fiber.Ctx, key string, v reflect.Value) error {
	param := ctx.Params(key)
	return convert(param, v, v.Type())
}

func extractQuery(ctx *fiber.Ctx, key string, v reflect.Value) error {
	param := ctx.Query(key)
	return convert(param, v, v.Type())
}

func extractBody(ctx *fiber.Ctx, v reflect.Value) error {
	return ctx.BodyParser(v.Addr().Interface())
}
