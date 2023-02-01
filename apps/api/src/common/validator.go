package common

import "github.com/go-playground/validator/v10"

type Validator struct {
	Package *validator.Validate
}

type ValidateError = map[string]string

// func (v *Validator) Execute(object *interface{}) error  {

// }
