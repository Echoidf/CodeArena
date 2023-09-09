package api

import (
	"CodeArena/consts"
	"CodeArena/models"
	"CodeArena/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Register(c echo.Context) (err error) {
	var user = new(models.User)
	if err = c.Bind(user); err != nil {
		return consts.InvalidInputError
	}

	user.Salt = utils.RandomString(6)
	user.Password = utils.GenMd5(user.Salt, user.Password)

	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	affected, err := models.AddUser(user)
	if err != nil {
		return consts.InsertError
	}
	if affected == 0 {
		return c.String(http.StatusOK, "NoAffected")
	}

	return c.NoContent(http.StatusOK)
}
