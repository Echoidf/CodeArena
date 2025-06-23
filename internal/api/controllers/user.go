package controllers

import (
	"content-assistant/internal/db/repository"
	"content-assistant/pkg/common"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	res, err := repository.GetUserList()
	if err != nil {
		common.Error(c, 500, err.Error())
		return
	}

	common.Success(c, res)
}
