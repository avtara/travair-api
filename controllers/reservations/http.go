package reservations

import (
	"github.com/avtara/travair-api/app/middleware"
	"github.com/avtara/travair-api/businesses/reservations"
	"github.com/avtara/travair-api/controllers/reservations/request"
	"github.com/avtara/travair-api/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ReservationController struct {
	reservationService  reservations.Service
}

func NewReservationController(uc reservations.Service) *ReservationController {
	return &ReservationController{
		reservationService:  uc,
	}
}

func (ctrl *ReservationController) Reservation(c echo.Context) error {
	claim := middleware.GetUser(c)
	ctx := c.Request().Context()

	req := new(request.Reservation)
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

	res, err := ctrl.reservationService.Reservation(ctx, request.ToDomain(req, claim.UserID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Success reservation hotel!",
			map[string]int{"price":res.Price}))
}