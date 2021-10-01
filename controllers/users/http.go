package users

import (
	"github.com/avtara/travair-api/businesses/users"
	"github.com/avtara/travair-api/controllers/users/request"
	"github.com/avtara/travair-api/helpers"
	echo "github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"net/http"
)

type UserController struct {
	userService users.Service
	publishQueue *amqp.Channel
}

func NewUserController(uc users.Service, pubAmpq *amqp.Channel) *UserController {
	return &UserController{
		userService: uc,
		publishQueue: pubAmpq,
	}
}

func (ctrl *UserController) Registration(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.UserRegistration)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("Internal Server Error", err.Error(), helpers.EmptyObj{}))
	}
	if err := c.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("An error occurred while validating the request data", err.Error(), helpers.EmptyObj{}))
	}

	res,err := ctrl.userService.Registration(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("Internal Server Error", err.Error(), helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated, helpers.BuildErrorResponse("Successfully created an account, please check your email to activate!", err.Error(), res))
}