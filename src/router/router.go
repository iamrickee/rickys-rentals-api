package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/auth"
)

func Route() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Router!")
	})

	e.GET("/register", func(c echo.Context) error {
		return auth.Register(c)
	})

	e.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "Login")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
