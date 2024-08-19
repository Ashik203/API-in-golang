package utils

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func Validator(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
