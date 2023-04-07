package helper

import (
	"errors"
	"ocra_server/model/response"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	registerAuthValidator()
}

func registerAuthValidator() {
	Validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		inputString := fl.Field().String()

		isContainUpper := false
		isContainLower := false
		isContainsNumber := false
		for _, v := range inputString {
			if unicode.IsUpper(v) {
				isContainUpper = true
			}

			if unicode.IsLower(v) {
				isContainLower = true
			}

			if unicode.IsNumber(v) {
				isContainsNumber = true
			}

			if isContainLower && isContainUpper && isContainsNumber {
				break
			}
		}

		return isContainUpper && isContainLower && isContainsNumber
	})

}

func ValidateUserInput(model any) error {
	if err := Validate.Struct(model); err != nil {
		return errors.New(response.MessageInvalidJsonInput)
	}

	return nil
}
