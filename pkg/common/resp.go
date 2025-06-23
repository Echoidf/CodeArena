package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func json(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

func Success(c *gin.Context, data any) {
	json(c, http.StatusOK, Response{Success: true, Data: data})
}

func Error(c *gin.Context, statusCode int, message string) {
	json(c, statusCode, Response{Error: message})
}

func BadRequest(c *gin.Context, message string) {
	json(c, http.StatusBadRequest, Response{Error: message})
}

func InternalError(c *gin.Context, message string) {
	json(c, http.StatusInternalServerError, Response{Error: message})
}

func NotFound(c *gin.Context, message string) {
	json(c, http.StatusNotFound, Response{Error: message})
}

func Forbidden(c *gin.Context, message string) {
	json(c, http.StatusForbidden, Response{Error: message})
}

func Unauthorized(c *gin.Context, message string) {
	json(c, http.StatusUnauthorized, Response{Error: message})
}

func UnprocessableEntity(c *gin.Context, message string) {
	json(c, http.StatusUnprocessableEntity, Response{Error: message})
}
