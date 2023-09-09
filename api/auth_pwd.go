package api

import (
	"CodeArena/consts"
	"CodeArena/models"
	"CodeArena/utils"
	"errors"
)

func init() {
	authActionMap[consts.LOGIN_BY_PWD] = GetUserByPwd
}

func GetUserByPwd(authForm *AuthForm) (user *models.User, err error) {
	user, err = models.GetModelByFields[models.User]("user", map[string]interface{}{
		"username": authForm.Username,
	})

	if err != nil {
		return nil, err
	}

	// 密码不匹配
	if user.Password != utils.GenMd5(user.Salt, authForm.Password) {
		return nil, errors.New("password didn't match")
	}
	return
}
