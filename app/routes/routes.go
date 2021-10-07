package routes

import (
	_middleware "github.com/avtara/travair-api/app/middleware"
	"github.com/avtara/travair-api/controllers/reservations"
	"github.com/avtara/travair-api/controllers/units"
	"github.com/avtara/travair-api/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController users.UserController
	JWTMiddleware  middleware.JWTConfig
	UnitController units.UnitController
	ReservationController reservations.ReservationController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("", cl.UserController.Registration)
	users.POST("/activation/:id", cl.UserController.Activation)

	auth := e.Group("authentications")
	auth.POST("", cl.UserController.Login)

	unit := e.Group("units")
	unit.POST("", cl.UnitController.AddUnit, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("tenant"))
	unit.PUT("/:id", cl.UnitController.UpdateThumbnail, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("tenant"))
	unit.POST("/:id/photos", cl.UnitController.AddPhotos, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("tenant"))
	unit.GET("/:id", cl.UnitController.GetDetail)
	unit.GET("", cl.UnitController.GetUnits)

	resv := e.Group("reservations")
	resv.POST("", cl.ReservationController.Reservation, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation("guest"))
}