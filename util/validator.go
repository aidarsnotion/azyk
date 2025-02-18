package util

import "github.com/go-playground/validator/v10"

var validate = validator.New()

// Функция для валидации структуры
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
