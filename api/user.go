package api

import (
	"CodeArena/common"
	"CodeArena/models"
	"CodeArena/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Register(c echo.Context) (err error) {
	var user = new(models.User)
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, common.InvalidInputError.WithError(err))
	}

	user.Salt = utils.RandomString(6)
	user.Password = utils.GenMd5(user.Salt, user.Password)

	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

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
