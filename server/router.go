package server

import (
	"CodeArena/api"
	"CodeArena/conf"
	"CodeArena/consts"
	"CodeArena/embeds"
	"CodeArena/middware"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func Start() {
	e := echo.New()

	// 自定义banner
	e.HideBanner = true
	fmt.Println(string(embeds.ReadEmbeds(consts.BANNER)))

	e.Use(middleware.Recover())
	e.Use(middware.CORS)
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "time=${time_rfc3339}, remoteIp=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
			Output: conf.GetLogFile(),
		}))

	// 健康检查
	e.Any("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	// 注册swagger
	conf.RegisterSwag(e)

	r := e.Group("/api/v1")
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)

	// 以下接口需要token认证
	r.Use(middware.Authorize)

	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.V.GetInt("server.port"))))
}
