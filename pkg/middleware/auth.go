package middleware

import (
	"content-assistant/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return JWT(jwt.AccessTokenType)
}

func JWTRefresh() gin.HandlerFunc {
	return JWT(jwt.RefreshTokenType)
}

func JWT(tokenType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("Authorization")
		token, _ := c.Cookie("auth-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, "token invalid")
			c.Abort()
			return
		}
		payload, err := jwt.ValidateToken(token)
		if err != nil || payload == nil || payload["type"] != tokenType {
			c.JSON(http.StatusUnauthorized, "token invalid")
			c.Abort()
			return
		}
		c.Set("userId", payload["id"])
		c.Next()
	}
}
