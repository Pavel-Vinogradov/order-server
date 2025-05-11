package pkg

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	return validator.New().Struct(&payload)
}
