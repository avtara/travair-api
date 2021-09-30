package users

import (
	"github.com/avtara/travair-api/businesses/users"
	"github.com/avtara/travair-api/controllers/users/request"
	"github.com/avtara/travair-api/controllers/users/response"
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
		return helpers.BuildErrorResponse(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return helpers.BuildErrorValidatorResponse(c, err)
	}

	res,err := ctrl.userService.Registration(ctx, req.ToDomain())
	if err != nil {
		return helpers.BuildErrorResponse(c, http.StatusInternalServerError, err)
	}

	return helpers.BuildResponse(c, response.FromDomain(res))
}