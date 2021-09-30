package helpers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Base struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(c echo.Context, data interface{}) error {
	res := Base{
		Status:  http.StatusOK,
		Message: "Success",
		Errors:  nil,
		Data:    data,
	}

	return c.JSON(res.Status, res)
}

func BuildErrorResponse(c echo.Context, status int, err error) error {
	res := Base{
		Status:  status,
		Message: "Something not right",
		Errors:  []string{err.Error()},
		Data:    nil,
	}

	return c.JSON(status, res)
}

func BuildErrorValidatorResponse(c echo.Context, err error) error {
	var errs []string
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errs = append(errs,fmt.Sprintf("%s is required",
					err.Field()))
			case "email":
				errs = append(errs,fmt.Sprintf("%s is not valid",
					err.Field()))
			case "gte":
				errs = append(errs,fmt.Sprintf("%s is less than %s",
					err.Field(), err.Param()))
			case "lte":
				errs = append(errs,fmt.Sprintf("%s is greater than %s",
					err.Field(), err.Param()))
			case "password":
				errs = append(errs,fmt.Sprintf("%s is not strong enough",
					err.Field()))
			case "role":
				errs = append(errs,fmt.Sprintf("%s is not valid",
					err.Field()))
			}
		}
	}

	res := Base{
		Status:  http.StatusBadRequest,
		Message: "Something not right",
		Errors:  errs,
		Data:    nil,
	}

	return c.JSON(res.Status, res)
}
