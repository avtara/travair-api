package middleware

import (
	"errors"
	"github.com/avtara/travair-api/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RoleValidation(role string) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := GetUser(c)

			if claims.Roles == role {
				return hf(c)
			} else {
				return c.JSON(http.StatusForbidden, helpers.BuildErrorResponse("Forbidden Access!",
					errors.New("invalid roles"), helpers.EmptyObj{}))
			}
		}
	}
}
