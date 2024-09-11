package middware

import (
	"codearena/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

var whiteList = map[string]bool{
	"/api/user/login": true,
}

func JWTMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if whiteList[ctx.Path()] {
			return ctx.Next()
		}
		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		log.Info(("访问鉴权：" + claims["username"].(string)))
		return jwtware.New(jwtware.Config{
			SigningKey: []byte(config.GetConfig().AuthSecret),
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Invalid token or expired",
				})
			},
		})(ctx)
	}
}
