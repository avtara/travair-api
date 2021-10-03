package users

import (
	"github.com/avtara/travair-api/businesses/users"
	"github.com/avtara/travair-api/controllers/users/request"
	"github.com/avtara/travair-api/controllers/users/response"
	"github.com/avtara/travair-api/helpers"
	echo "github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"net/http"
	"strings"
)

type UserController struct {
	userService  users.Service
	publishQueue *amqp.Channel
}

func NewUserController(uc users.Service, pubAmpq *amqp.Channel) *UserController {
	return &UserController{
		userService:  uc,
		publishQueue: pubAmpq,
	}
}

func (ctrl *UserController) Registration(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.UserRegistration)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	res, err := ctrl.userService.Registration(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created an account, please check your email to activate!",
			response.FromDomain(res)))
}

func (ctrl *UserController) Activation(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Param("id")
	res, err := ctrl.userService.Activation(ctx, userID)
	if err != nil {
		if strings.Contains(err.Error(), "activated") {
			return c.JSON(http.StatusConflict,
				helpers.BuildErrorResponse("Conflict Data",
					err, helpers.EmptyObj{}))
		} else if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound,
				helpers.BuildErrorResponse("Data Not found",
					err, helpers.EmptyObj{}))
		}
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Successfully created an account, please check your email to activate!",
			response.FromDomain(res)))
}
