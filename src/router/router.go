package router

import (
	"crypto/subtle"
	"net/http"

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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Router!")
	})

	e.POST("/register", func(c echo.Context) error {
		return auth.Register(c)
	})

	e.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "Login")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
