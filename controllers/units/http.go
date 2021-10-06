package units

import (
	"github.com/avtara/travair-api/app/middleware"
	"github.com/avtara/travair-api/businesses/units"
	"github.com/avtara/travair-api/controllers/units/request"
	"github.com/avtara/travair-api/controllers/units/response"
	"github.com/avtara/travair-api/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UnitController struct {
	unitService  units.Service
}

func NewUnitController(uc units.Service) *UnitController {
	return &UnitController{
		unitService:  uc,
	}
}

func (ctrl *UnitController) AddUnit(c echo.Context) error {
	claim := middleware.GetUser(c)
	ctx := c.Request().Context()

	req := new(request.Unit)
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

	res, err := ctrl.unitService.Add(ctx, req.AddUnitToDomain(), claim.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Success add units!",
			response.FromDomain(res)))
}

func (ctrl *UnitController) UpdateThumbnail(c echo.Context) error {
	ctx := c.Request().Context()
	unitID := c.Param("id")
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}


	if err = ctrl.unitService.ChangeThumbnail(ctx, unitID, file); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Success update thumbnail!",
			map[string]string{"unit_id": unitID}))
}

func (ctrl *UnitController) AddPhotos(c echo.Context) error {
	ctx := c.Request().Context()
	unitID := c.Param("id")
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	if err = ctrl.unitService.AddPhoto(ctx, unitID, file); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildResponse("Success add photos!",
			map[string]string{"unit_id": unitID}))
}

func (ctrl *UnitController) GetDetail(c echo.Context) error {
	ctx := c.Request().Context()
	unitID := c.Param("id")

	res, err := ctrl.unitService.Detail(ctx, unitID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Internal Server Error",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildResponse("Success get detail unit!",
			response.DetailFromDomain(res)))
}
