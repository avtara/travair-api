package controller

import (
	"github.com/avtara/travair-api/lib/database"
	"github.com/avtara/travair-api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegistrationUserController(c echo.Context) (err error) {
	u := new(RegistrationUserRequest)
	if err := c.Bind(u); err != nil {
		return nil
	}
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ValidatorErrorResponse("An error occurred while validating the request data", err, utils.EmptyObj{}))
	}
	res, e := database.InsertUser(RegistrationUserRequestToModelUser(u))
	if e != nil {
		return c.JSON(http.StatusBadRequest, utils.BuildErrorResponse("Conflict", e.Error(), utils.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated, utils.BuildResponse(true,"successfully created an account, please check your email to activate!",ModelUserToRegistrationUserResponse(res)))
}
