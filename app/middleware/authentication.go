package middleware

import (
	"github.com/avtara/travair-api/helpers"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

type JwtCustomClaims struct {
	UserID uuid.UUID `json:"id"`
	Roles  string    `json:"roles"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return c.JSON(http.StatusForbidden, helpers.BuildErrorResponse("Forbidden Access!", e, helpers.EmptyObj{}))
		},
	}
}

func (jwtConf *ConfigJWT) GenerateToken(userID uuid.UUID, roles string) string {
	claims := JwtCustomClaims{
		userID,
		roles,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token
}

func GetClaimsToken(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
