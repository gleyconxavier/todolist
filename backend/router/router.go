package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New instance of echo
func New() *echo.Echo {
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	return e
}
