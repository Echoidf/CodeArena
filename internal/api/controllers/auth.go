package controllers

import (
	"codearena/internal/db/repository"
	"codearena/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// Login handles the login request
func Login(c *fiber.Ctx) error {
	req := new(repository.AuthRequest)
	if err := c.BodyParser(req); err != nil {
		return response.BadRequest(c, err)
	}
	resp, err := repository.Login(req)
	if err != nil {
		return response.Error(c, err)
	}
	return response.Success(c, resp)
}
