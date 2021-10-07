package reservations_test

import (
	"errors"
	"fmt"
	"github.com/avtara/travair-api/app/middleware"
	_reservationB "github.com/avtara/travair-api/businesses/reservations"
	"github.com/avtara/travair-api/businesses/reservations/mocks"
	"github.com/avtara/travair-api/controllers/reservations"
	"github.com/avtara/travair-api/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	mockReservationService mocks.Service
	claims                 *middleware.JwtCustomClaims
	errBindReq             string
	errValidateReq         string
	req                    string
	reservationsCtrl       *reservations.ReservationController
)

func TestMain(m *testing.M) {
	uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
	reservationsCtrl = reservations.NewReservationController(&mockReservationService)
	claims = &middleware.JwtCustomClaims{UserID: uuidID, Roles: "guest"}
	errBindReq = `{
    "price": 20000
    "unit_id": "3699a5b9-5fd0-4dd6-9dbb-f8484c711a1b",
    "check_in_date": "2006-01-10",
    "check_out_date": "2006-01-20"
}`
	errValidateReq = `{
    "unit_id": "3699a5b9-5fd0-4dd6-9dbb-f8484c711a1b",
    "check_in_date": "2006-01-10"
}`
	req = `{
    "price": 20000,
    "unit_id": "3699a5b9-5fd0-4dd6-9dbb-f8484c711a1b",
    "check_in_date": "2006-01-10",
    "check_out_date": "2006-01-20"
}`
	m.Run()
}

func TestReservationController_Reservation(t *testing.T) {
	t.Run("fail bind", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/reservations", strings.NewReader(errBindReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		if assert.NoError(t, reservationsCtrl.Reservation(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("fail validate", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/reservations", strings.NewReader(errValidateReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		if assert.NoError(t, reservationsCtrl.Reservation(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			fmt.Println(rec)
		}
	})

	t.Run("fail get data from service", func(t *testing.T) {
		mockReservationService.On("Reservation", mock.Anything,mock.Anything).Return(nil, errors.New("")).Once()
		req := httptest.NewRequest(http.MethodPost, "/reservations", strings.NewReader(req))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		if assert.NoError(t, reservationsCtrl.Reservation(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			fmt.Println(rec)
		}
	})

	t.Run("success", func(t *testing.T) {
		mockReservationService.On("Reservation", mock.Anything,mock.Anything).Return(&_reservationB.Domain{}, nil).Once()
		req := httptest.NewRequest(http.MethodPost, "/reservations", strings.NewReader(req))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		if assert.NoError(t, reservationsCtrl.Reservation(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			fmt.Println(rec)
		}
	})
}
