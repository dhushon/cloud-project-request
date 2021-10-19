package model

import (
	"sync"
	"github.com/go-playground/validator/v10"
)

var once sync.Once

var validate *validator.Validate

func getValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})
	return validate
}
