package middware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Claims struct {
	TokenType string `json:"tokenType,omitempty"`
	jwt.RegisteredClaims
}

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

		_, err := ParseToken(formatToken(header))

		//c.Set("user", claims.User)

		if err != nil {
			// 验证不通过，不再调用后续的函数处理
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		//zap.L().Info("访问鉴权：{%s}", claims.User.Name)

		// 设置userinfo
		go func() {
		}()
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

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		certificate, err := jwt.ParseRSAPublicKeyFromPEM(cert.Certificate)

		if err != nil {
			return nil, err
		}

		return certificate, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
