package router

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"iamricky.com/truck-rental/auth"
	"iamricky.com/truck-rental/config"
)

func Route() {
	e := echo.New()

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(config.Load("API_USER"))) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(config.Load("API_PASSWORD"))) == 1 {
			return true, nil
		}
		return false, nil
	}))

	e.POST("/register", func(c echo.Context) error {
		return auth.RegisterRoute(c)
	})

	e.POST("/login", func(c echo.Context) error {
		return auth.LoginRoute(c)
	})

	e.POST("/token", func(c echo.Context) error {
		return auth.TokenRoute(c)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
