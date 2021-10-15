package model

import (
	"sync"
	"gopkg.in/go-playground/validator.v9"
)

var once sync.Once

var validate *validator.Validate

func getValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})
	return validate
}
