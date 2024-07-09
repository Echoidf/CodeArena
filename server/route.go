package server

import (
	"codearena/consts"
	"codearena/embeds"
	"codearena/internal/api/controllers"
	"codearena/pkg/config"
	"codearena/pkg/middleware"
	"codearena/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {
	middleware.InitLogger(&config.GetConfig().Logger)

	r = gin.Default()

	r.Use(middleware.GinLogger(), middleware.GinRecovery(true), GinMiddleware())

	r.Any("/health", func(c *gin.Context) {
		response.Success(c, "OK")
	})

	setupAuthRoutes(r)
	setupUserRoutes(r)
	setupFileRoutes(r)
}

func setupFileRoutes(router *gin.Engine) {
	fileRoutes := router.Group("/api/file")
	fileRoutes.Use(middleware.JWTAuth())
	fileRoutes.POST("/upload", controllers.UploadFile)
}

func setupAuthRoutes(router *gin.Engine) {
	authRoutes := router.Group("/api/auth")
	authRoutes.POST("/login", controllers.Login)
	authRoutes.POST("/signup", controllers.Signup)
	authRoutes.POST("/logout", controllers.Logout)
	authRoutes.POST("/activate", controllers.ActivateAccount)
}

func setupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/api/user")
	userRoutes.Use(middleware.JWTAuth())
	userRoutes.GET("/list", controllers.GetUsers)
}

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func Run() {
	fmt.Println(string(embeds.ReadEmbeds(consts.BANNER)))
	_ = r.Run(fmt.Sprintf(":%d", config.GetConfig().HttpPort))
}
