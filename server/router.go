package server

import (
	"CodeArena/conf"
	"CodeArena/consts"
	"CodeArena/embeds"
	"CodeArena/middware"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Start() {
	e := echo.New()

	// 自定义banner
	e.HideBanner = true
	fmt.Println(string(embeds.ReadEmbeds(consts.BANNER)))

	e.Use(middleware.Recover())

	// 健康检查
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	// 注册swagger
	conf.RegisterSwag(e)

	r := e.Group("/api/v1")
	r.Use(middware.Authorize)

	e.Start(fmt.Sprintf(":%d", conf.V.GetInt("server.port")))
}
