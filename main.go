package main

import (
	"codearena/pkg/config"
	"codearena/server"

	"github.com/gin-gonic/gin"
)

func main() {
	if config.GetConfig().Environment == config.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	server.Run()
}
