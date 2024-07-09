package controllers

import (
	"codearena/internal/db/repository"
	"codearena/pkg/response"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	res, err := repository.GetUserList()
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, res)
}
