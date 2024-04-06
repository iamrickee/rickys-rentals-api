package router

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"iamricky.com/truck-rental/auth"
	"iamricky.com/truck-rental/config"
	"iamricky.com/truck-rental/location"
	authMiddleware "iamricky.com/truck-rental/middleware"
	"iamricky.com/truck-rental/rental"
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

	authGroup := e.Group("/verify")
	authGroup.Use(authMiddleware.Verify)

	authGroup.GET("/locations", func(c echo.Context) error {
		return location.ListRoute(c)
	})
	authGroup.PUT("/location/save", func(c echo.Context) error {
		return location.SaveRoute(c)
	})
	authGroup.DELETE("/location/delete", func(c echo.Context) error {
		return location.DeleteRoute(c)
	})
	authGroup.GET("/location/get", func(c echo.Context) error {
		return location.GetRoute(c)
	})

	authGroup.GET("/rentals", func(c echo.Context) error {
		return rental.ListRoute(c)
	})
	authGroup.PUT("/rental/save", func(c echo.Context) error {
		return rental.SaveRoute(c)
	})
	authGroup.DELETE("/rental/delete", func(c echo.Context) error {
		return rental.DeleteRoute(c)
	})
	authGroup.GET("/rental/get", func(c echo.Context) error {
		return rental.GetRoute(c)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
