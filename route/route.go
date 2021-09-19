package route

import (
	"github.com/avtara/travair-api/controller"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	return e
}
