package users_test

import (
	"errors"
	"fmt"
	"github.com/avtara/travair-api/app/middleware"
	_user "github.com/avtara/travair-api/businesses/users"
	"github.com/avtara/travair-api/businesses/users/mocks"
	"github.com/avtara/travair-api/controllers/users"
	"github.com/avtara/travair-api/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	mockUsersService mocks.Service
	queueConn        *amqp.Channel
	claims           *middleware.JwtCustomClaims
	usersCtrl        *users.UserController
	regisErrBind     string
	regisErrValidate string
	regisReq         string
	loginErrBind     string
	loginErrValidate string
	loginReq         string
)

func TestMain(m *testing.M) {
	uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
	usersCtrl = users.NewUserController(&mockUsersService, queueConn)
	claims = &middleware.JwtCustomClaims{UserID: uuidID, Roles: "guest"}
	regisErrBind = `{
    "email": "tenant@testing.com"
    "password": "Adwawdaw!23!",
    "name":"Alexa",
    "role": "tenant"
}`
	regisErrValidate = `{
    "email": "tenant@testing.com",
    "password": "Adwawdaw!23!",
    "name":"Alexa"
}`
	regisReq = `{
    "email": "tenant@testing.com",
    "password": "Adwawdaw!23!",
    "name":"Alexa",
    "role": "tenant"
}`
	loginErrBind = `{
    "email": "wadawd"
    "password" : "dawjdwaidj"
}`
	loginErrValidate = `{
	}`
	loginReq = `{
    "email": "wadawd@s.com",
    "password" : "dawjdwaidj"
}`
	m.Run()
}

func TestUserController_Registration(t *testing.T) {
	t.Run("fail bind", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(regisErrBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		if assert.NoError(t, usersCtrl.Registration(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("fail validate", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(regisErrValidate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		if assert.NoError(t, usersCtrl.Registration(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("fail store", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(regisReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Registration", mock.Anything, mock.Anything).Return(&_user.Domain{}, errors.New("")).Once()

		if assert.NoError(t, usersCtrl.Registration(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("success store", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(regisReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Registration", mock.Anything, mock.Anything).Return(&_user.Domain{}, nil).Once()

		if assert.NoError(t, usersCtrl.Registration(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			fmt.Println(rec)
		}
	})
}

func TestUserController_Login(t *testing.T) {
	t.Run("fail bind", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/authentications", strings.NewReader(loginErrBind))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		if assert.NoError(t, usersCtrl.Login(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("fail validate", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/authentications", strings.NewReader(loginErrValidate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		if assert.NoError(t, usersCtrl.Login(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("wrong credential", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/authentications", strings.NewReader(loginReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Login", mock.Anything, mock.Anything, mock.Anything).Return(&_user.Domain{}, errors.New("not match")).Once()

		if assert.NoError(t, usersCtrl.Login(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}
	})

	t.Run("unactivated account", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/authentications", strings.NewReader(loginReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Login", mock.Anything, mock.Anything, mock.Anything).Return(&_user.Domain{}, errors.New("not been activated")).Once()

		if assert.NoError(t, usersCtrl.Login(c)) {
			assert.Equal(t, http.StatusForbidden, rec.Code)
		}
	})

	t.Run("error login", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/authentications", strings.NewReader(loginReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Login", mock.Anything, mock.Anything, mock.Anything).Return(&_user.Domain{}, errors.New("")).Once()

		if assert.NoError(t, usersCtrl.Login(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/authentications", strings.NewReader(loginReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Login", mock.Anything, mock.Anything, mock.Anything).Return(&_user.Domain{}, nil).Once()

		if assert.NoError(t, usersCtrl.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestUserController_Activation(t *testing.T) {
	t.Run("conflict activated", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users/activation/291239128", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Activation", mock.Anything, mock.Anything).Return(&_user.Domain{}, errors.New("activated")).Once()
		if assert.NoError(t, usersCtrl.Activation(c)) {
			assert.Equal(t, http.StatusConflict, rec.Code)
		}
	})

	t.Run("not found account", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users/activation/291239128", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Activation", mock.Anything, mock.Anything).Return(&_user.Domain{}, errors.New("not found")).Once()
		if assert.NoError(t, usersCtrl.Activation(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("error activating", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users/activation/291239128", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Activation", mock.Anything, mock.Anything).Return(&_user.Domain{}, errors.New("err")).Once()
		if assert.NoError(t, usersCtrl.Activation(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users/activation/291239128", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockUsersService.On("Activation", mock.Anything, mock.Anything).Return(&_user.Domain{}, nil).Once()
		if assert.NoError(t, usersCtrl.Activation(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
