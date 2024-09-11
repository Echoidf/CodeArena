package server

import (
	"codearena/consts"
	"codearena/embeds"
	"codearena/internal/api/controllers"
	"codearena/pkg/config"
	"codearena/pkg/middware"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var app *fiber.App
var logFile *os.File

func init() {
	app = fiber.New()
	var err error
	logFile, err = os.OpenFile("./log/codearena.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	app.Use(logger.New(logger.Config{
		Format:     "${time}-${status}-${latency}-${method}-${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Output:     logFile,
	}))
	app.Use(recover.New(), cors.New())

	app.All("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	apiRouter := app.Group("/api")
	apiRouter.Use(middware.JWTMiddleware())
	setupFileRoutes(apiRouter)
	setupUserRoutes(apiRouter)
}

func setupFileRoutes(router fiber.Router) {
	router.Post("/file/upload", controllers.UploadFile)
}

func setupUserRoutes(router fiber.Router) {
	router.Post("/user/login", controllers.Login)
}

func Run() {
	defer logFile.Close()
	fmt.Println(string(embeds.ReadEmbeds(consts.BANNER)))
	app.Listen(fmt.Sprintf(":%d", config.GetConfig().HttpPort))
}
