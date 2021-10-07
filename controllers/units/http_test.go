package units_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/avtara/travair-api/app/middleware"
	_unit "github.com/avtara/travair-api/businesses/units"
	"github.com/avtara/travair-api/businesses/units/mocks"
	"github.com/avtara/travair-api/controllers/units"
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
	mockUnitService mocks.Service
	claims          *middleware.JwtCustomClaims
	unitCtrl        *units.UnitController
	addErrBind      string
	addErrValidate  string
	addReq          string
)

func TestMain(m *testing.M) {
	uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
	unitCtrl = units.NewUnitController(&mockUnitService)
	claims = &middleware.JwtCustomClaims{UserID: uuidID, Roles: "tenant"}
	addErrBind = `{
    "name":"Hotel Horizon"
    "category":"hotel",
    "price":250000,
    "description":"<b>Fasilitas</b></br><ul><li>Toilet berdiri</li></ul>",
    "policy":"<b>Perturan</b></br><ul><li>Dilarang merokok</li></ul>",
    "thumbnail":"https://asset.kompas.com/crops/bRGakO6FcLC80Cl6w27eUv5jUtU=/19x0:775x504/750x500/data/photo/2019/10/15/5da57133d3be7.jpg",
    "address":{
        "street":"Jalan Gunung mana aja",
        "city":"Purwokerto",
        "state":"Jawa Tengah",
        "country":"Indonesia",
        "postal_code":"53121",
        "latitude": -7.388060,
        "longitude": 109.363890
    }
}`
	addErrValidate = `{
    "category":"hotel",
    "price":250000,
    "description":"<b>Fasilitas</b></br><ul><li>Toilet berdiri</li></ul>",
    "policy":"<b>Perturan</b></br><ul><li>Dilarang merokok</li></ul>",
    "thumbnail":"https://asset.kompas.com/crops/bRGakO6FcLC80Cl6w27eUv5jUtU=/19x0:775x504/750x500/data/photo/2019/10/15/5da57133d3be7.jpg",
    "address":{
        "street":"Jalan Gunung mana aja",
        "city":"Purwokerto",
        "state":"Jawa Tengah",
        "country":"Indonesia",
        "postal_code":"53121",
        "latitude": -7.388060,
        "longitude": 109.363890
    }
}`
	addReq = `{
	"name":"Hotel Horizon",
    "category":"hotel",
    "price":250000,
    "description":"<b>Fasilitas</b></br><ul><li>Toilet berdiri</li></ul>",
    "policy":"<b>Perturan</b></br><ul><li>Dilarang merokok</li></ul>",
    "thumbnail":"https://asset.kompas.com/crops/bRGakO6FcLC80Cl6w27eUv5jUtU=/19x0:775x504/750x500/data/photo/2019/10/15/5da57133d3be7.jpg",
    "address":{
        "street":"Jalan Gunung mana aja",
        "city":"Purwokerto",
        "state":"Jawa Tengah",
        "country":"Indonesia",
        "postal_code":"53121",
        "latitude": -7.388060,
        "longitude": 109.363890
    }
}`
	m.Run()
}

func TestUnitController_AddUnit(t *testing.T) {
	t.Run("fail bind", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/units", strings.NewReader(addErrBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		if assert.NoError(t, unitCtrl.AddUnit(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("fail validate", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/units", strings.NewReader(addErrValidate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		if assert.NoError(t, unitCtrl.AddUnit(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("fail store", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/units", strings.NewReader(addReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockUnitService.On("Add", mock.Anything, mock.Anything, mock.Anything).Return(&_unit.Domain{}, errors.New("adw")).Once()

		if assert.NoError(t, unitCtrl.AddUnit(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			fmt.Println(rec)
		}
	})

	t.Run("fail store", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/units", strings.NewReader(addReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockUnitService.On("Add", mock.Anything, mock.Anything, mock.Anything).Return(&_unit.Domain{}, nil).Once()

		if assert.NoError(t, unitCtrl.AddUnit(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			fmt.Println(rec)
		}
	})
}

func TestUnitController_UpdateThumbnail(t *testing.T) {
	t.Run("fail bind", func(t *testing.T) {
		body := new(bytes.Buffer)

		req := httptest.NewRequest(http.MethodPut, "/units/waodkwd", body)
		req.Header.Add("Content-Type", "multipart/form-data")
		//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockUnitService.On("ChangeThumbnail", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("")).Once()

		if assert.NoError(t, unitCtrl.UpdateThumbnail(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestUnitController_AddPhotos(t *testing.T) {
	t.Run("fail bind", func(t *testing.T) {
		body := new(bytes.Buffer)

		req := httptest.NewRequest(http.MethodPost, "/units/506a1da7-1592-4d28-a067-af036aca07de/photos", body)
		req.Header.Add("Content-Type", "multipart/form-data")
		//req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockUnitService.On("ChangeThumbnail", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("")).Once()

		if assert.NoError(t, unitCtrl.AddPhotos(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestUnitController_GetUnits(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/units", strings.NewReader(addErrBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockUnitService.On("UnitsByGeo", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("")).Once()

		if assert.NoError(t, unitCtrl.GetUnits(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/units", strings.NewReader(addErrBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockUnitService.On("UnitsByGeo", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()

		if assert.NoError(t, unitCtrl.GetUnits(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestUnitController_GetDetail(t *testing.T) {
	t.Run("fail get", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/units/awdkoawdk", strings.NewReader(addErrBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockUnitService.On("Detail", mock.Anything, mock.Anything).Return(&_unit.Domain{}, errors.New("")).Once()

		if assert.NoError(t, unitCtrl.GetDetail(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/units/awdkoawdk", strings.NewReader(addErrBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockUnitService.On("Detail", mock.Anything, mock.Anything).Return(&_unit.Domain{}, nil).Once()

		if assert.NoError(t, unitCtrl.GetDetail(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
