package common

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type Validate interface {
	Exec(req interface{}) []*ErrorResponse
	SingleValue(req interface{}, tag string) error
}

type Validator struct {
	lib *validator.Validate
}

func (v *Validator) Exec(req interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := v.lib.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

func (v *Validator) SingleValue(req interface{}, tag string) error {
	return v.lib.Var(req, tag)
}

var validate = validator.New()

var ValidatorAdapter = Validator{validate}
