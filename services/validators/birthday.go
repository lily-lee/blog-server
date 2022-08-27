package validators

import (
	"github.com/go-playground/validator/v10"

	"github.com/lily-lee/blog-server/services/types"
)

func birthday(fl validator.FieldLevel) bool {
	if fl.Field().Interface() == nil {
		return true
	}

	t := fl.Field().Interface().(types.Birthday)
	_, err := t.Value()

	return err == nil
}
