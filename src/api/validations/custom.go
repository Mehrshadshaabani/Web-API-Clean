package validations

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    any    `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = fe.Field()
			el.Tag = fe.Tag()
			el.Value = fe.Value()
			el.Message = fe.Error()
			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}

	return nil

}
