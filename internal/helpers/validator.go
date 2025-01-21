package helpers

import "github.com/go-playground/validator/v10"

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	Validator struct {
		validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse
	return validationErrors
}
