package api

import (
	"CodeArena/common"
	"CodeArena/consts"
	"CodeArena/models"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type AuthForm struct {
	Username  string           `json:"username,omitempty"`
	Password  string           `json:"password,omitempty"`
	Captcha   string           `json:"captcha,omitempty"`
	Phone     string           `json:"phone,omitempty"`
	PhoneCode string           `json:"phoneCode,omitempty"`
	LoginType consts.LoginType `json:"loginType"`
}

var authActionMap map[consts.LoginType]func(authForm *AuthForm) (user *models.User, err error)

func init() {
	authActionMap = make(map[consts.LoginType]func(authForm *AuthForm) (user *models.User, err error))
}

func Login(c echo.Context) (err error) {
	authForm := new(AuthForm)
	if err = c.Bind(authForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fc, ok := authActionMap[authForm.LoginType]
	if !ok {
		zap.L().Error("loginType does not match")
		return c.JSON(http.StatusBadRequest, common.InvalidInputError)
	}

	// 登录校验具体逻辑执行
	user, err := fc(authForm)

	if err != nil {
		zap.L().Error("auth failed", zap.Error(err))
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := models.GenerateToken(user)

	if err != nil {
		zap.L().Error("generated token failed", zap.Error(err))
		return err
	}

	return c.JSON(http.StatusOK, models.Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: res,
	})
}
