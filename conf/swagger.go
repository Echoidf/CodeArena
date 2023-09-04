package conf

import (
	_ "CodeArena/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterSwag(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
