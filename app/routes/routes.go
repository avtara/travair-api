package routes

import (
	"github.com/avtara/travair-api/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController users.UserController
	JWTMiddleware  middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", cl.UserController.Registration)
	users.POST("/activation/:id", cl.UserController.Activation)

	auth := e.Group("authentications")
	auth.POST("", cl.UserController.Login)
}