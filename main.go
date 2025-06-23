package main

import (
	"content-assistant/pkg/config"
	"content-assistant/server"

	"github.com/gin-gonic/gin"
)

func main() {
	if config.GetConfig().Environment == config.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	server.Run()
}
