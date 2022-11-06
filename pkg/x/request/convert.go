package request

import (
	"reflect"
	"strconv"
)

func convert(s string, v reflect.Value, t reflect.Type) error {
	switch t.Kind() {
	case reflect.String:
		v.SetString(s)
	case reflect.Int:
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		v.SetInt(int64(i))
	}
	return nil
}
