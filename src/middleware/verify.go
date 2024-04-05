package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/auth"
)

func Verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Request().Header.Get("X-username")
		token := c.Request().Header.Get("X-token")
		err := auth.TokenLogin(c, name, token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
		}
		return next(c)
	}
}
