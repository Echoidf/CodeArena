package middware

import (
	"CodeArena/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type TokenRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireIn     int64  `json:"expireIn"`
}

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return c.JSON(http.StatusUnauthorized, "please login first")
		}

		claims, err := models.ParseToken(formatToken(header))

		c.Set("userId", claims.UserId)

		if err != nil {
			// 验证不通过，不再调用后续的函数处理
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		zap.L().Info(fmt.Sprintf("访问鉴权：{%s}", claims.UserName))

		return next(c)
	}
}

func formatToken(header string) string {
	tokens := strings.Split(header, " ")
	if len(tokens) != 2 {
		return ""
	}
	if tokens[0] != "Bearer" {
		return ""
	}
	return tokens[1]
}
