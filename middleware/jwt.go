package middleware

import (
	"echo-golang/model"
	service "echo-golang/service/user"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		config := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(service.JwtCustomClaims)
			},
			SigningKey: []byte("secret"),
			ErrorHandler: func(c echo.Context, err error) error {
				return c.JSON(http.StatusUnauthorized, model.BaseResponseNoData{
					IsSuccess: false,
					Message:   "Unauthorized",
				})
			},
		}
		echojwt.WithConfig(config)
		return next(c)
	}
}