package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//IsLoggedIn Evaluate the config of the JWT
var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

//IsAdmin Check if an user have the right to do an action
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)
		if isAdmin == false {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
