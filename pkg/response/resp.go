package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
}

func json(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
func Success(c *gin.Context, data interface{}) {
	json(c, 200, Response{Success: true, Data: data})
}
