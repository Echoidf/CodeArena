package api

import (
	"CodeArena/common"
	"CodeArena/models"
	"CodeArena/utils"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func Register(c echo.Context) (err error) {
	var user = new(models.User)
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, common.InvalidInputError.WithError(err))
	}

	user.Salt = utils.RandomString(6)
	user.Password = utils.GenMd5(user.Salt, user.Password)

	affected, err := models.AddUser(user)
	if err != nil {
		zap.L().Error(fmt.Sprintf("add user failed, err:%v", err.Error()))
		return c.JSON(http.StatusBadRequest, common.InsertError.WithError(err))
	}
	if affected == 0 {
		return c.String(http.StatusOK, "NoAffected")
	}

	return c.String(http.StatusOK, "ok")
}

func GetUserList(c echo.Context) (err error) {
	users, err := models.GetUsers()

	if err != nil {
		zap.L().Error(fmt.Sprintf("get user list failed, err: %v", err.Error()))
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserInfo(c echo.Context) (err error) {
	var user = new(models.User)

	userId := c.Get("userId")
	user.Id = int64(userId.(float64))
	_, err = models.Engine.Cols("username", "email", "created_at", "avatar", "phone").Get(user)

	if err != nil {
		zap.L().Error(fmt.Sprintf("get userInfo failed, err: %v", err.Error()))
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) (err error) {
	var user = new(models.User)
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, common.InvalidInputError.WithError(err))
	}

	if user.Id == 0 {
		return c.JSON(http.StatusBadRequest, common.InvalidInputError.WithError(errors.New("user id is not provided")))
	}

	affected, err := models.Engine.ID(user.Id).Update(user)
	if err != nil {
		zap.L().Error(fmt.Sprintf("update user failed, err:%v", err.Error()))
		return err
	}

	if affected == 0 {
		return c.String(http.StatusOK, "NoAffected")
	}

	return c.String(http.StatusOK, "ok")
}
