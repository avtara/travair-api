package helpers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildValidatorErrorResponse(message string, err error, data interface{}) Response {
	var errs []string
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errs = append(errs,fmt.Sprintf("%s is required",
					err.Field()))
			case "email":
				errs = append(errs,fmt.Sprintf("%s is required",
					err.Field()))
			case "gte":
				errs = append(errs,fmt.Sprintf("%s is required",
					err.Field()))
			case "lte":
				errs = append(errs,fmt.Sprintf("%s is required",
					err.Field()))
			case "password":
				errs = append(errs,fmt.Sprintf("%s is not strong enough",
					err.Field()))
			case "role":
				errs = append(errs,fmt.Sprintf("%s is required",
					err.Field()))
			}

		}
	}
	res := Response{
		Status:  false,
		Message: message,
		Errors:  errs,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splitError,
		Data:    data,
	}
	return res
}