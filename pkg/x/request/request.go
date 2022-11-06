package request

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"github.com/GalvinGao/gofiber-template/pkg/x/validator"
)

// Inspect extracts the request into dest, populates empty fields with default values,
// then validates it using the `validate` tag.
//
// dest must be a pointer to a struct. The struct fields must be exported.
//
// The first stage is default value population.
// The supported struct fields for populating the default values are:
// - `default:"value"`: populates the value with the default value.
//
// The second stage is extraction.
// The supported struct fields for extracting the values are:
// - path category: `path:"name"`: the value is extracted from the path parameter.
// - query category: `query:"name"`: the value is extracted from the query parameter.
// - body category: `body:"type"`: the value is extracted from the request body.
//   - `json`: the value is extracted from the request body as JSON.
//   - `form`: the value is extracted from the request body as form.
//
// There shall be no more than one field in the struct for each category.
//
// The third stage is validation.
// The supported struct fields for validating the values are:
// - `validate:"validation"`: the value is validated using the validation tag.
//
// If validation fails, an error is returned.
func Inspect(ctx *fiber.Ctx, dest any) error {
	t := reflect.TypeOf(dest)
	v := reflect.ValueOf(dest)
	if t.Kind() != reflect.Ptr {
		return errors.New("request must be a pointer")
	}

	if t.Elem().Kind() != reflect.Struct {
		return errors.New("request must be a pointer to a struct")
	}

	// first stage: default value population
	for i := 0; i < t.Elem().NumField(); i++ {
		if err := populateDefault(ctx, t.Elem().Field(i), v.Elem().Field(i)); err != nil {
			return err
		}
	}

	// second stage: extraction
	for i := 0; i < t.Elem().NumField(); i++ {
		if err := extract(ctx, t.Elem().Field(i), v.Elem().Field(i)); err != nil {
			return err
		}
	}

	// third stage: validation
	if err := validator.Struct(dest); err != nil {
		return convertErrors(ctx, err)
	}

	return nil
}
