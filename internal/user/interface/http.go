package http

import (
	"yoyoj-go-rag/internal/user/domain/repository"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
	})
}
