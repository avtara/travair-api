package route

import (
	"github.com/avtara/travair-api/controller"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/users", controller.RegistrationUserController)
	return e
}
