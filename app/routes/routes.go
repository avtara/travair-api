package routes

import (
	"github.com/avtara/travair-api/controllers/users"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/", cl.UserController.Registration)
}