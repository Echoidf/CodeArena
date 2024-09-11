package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
}

func Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Success: false,
		Error:   err.Error(),
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Success: false,
		Error:   "Not Found",
	})
}

func Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Success: false,
		Error:   "Unauthorized",
	})
}

func BadRequest(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Success: false,
		Error:   err.Error(),
	})
}

func ErrorWithStatus(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Error:   err.Error(),
	})
}
