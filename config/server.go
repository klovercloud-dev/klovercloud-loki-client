package config

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	InitEnvironmentVariables()
	echoInstance := echo.New()
	echoInstance.Use(middleware.Recover())
	return echoInstance
}
