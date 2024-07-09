package jwt

import (
	"codearena/pkg/config"
	"codearena/pkg/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/quangdangfit/gocommon/logger"
)

const (
	AccessTokenExpiredTime = 24 * 60 * 60
	RefreshTokenExpireTime = 30 * 24 * 60 * 60
	AccessTokenType        = "access"
	RefreshTokenType       = "x-refresh"
)

func GenerateAccessToken(payload map[string]interface{}) string {
	cfg := config.GetConfig()
	payload["type"] = AccessTokenType
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * AccessTokenExpiredTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate access token", err)
		return ""
	}
	return token
}

func GenerateRefreshToken(payload map[string]interface{}) string {
	cfg := config.GetConfig()
	payload["type"] = RefreshTokenType
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * RefreshTokenExpireTime).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecret))
	if err != nil {
		logger.Error("Failed to generate refresh token: ", err)
		return ""
	}
	return token
}

func ValidateToken(jwtToken string) (map[string]interface{}, error) {
	cfg := config.GetConfig()
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AuthSecret), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}
	var data map[string]interface{}
	utils.Copy(&data, tokenData["payload"])

	return data, nil
}
