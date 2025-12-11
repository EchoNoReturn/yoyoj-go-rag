package http

import (
	"yoyoj-go-rag/internal/user/infrastructure"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app fiber.Router, db *gorm.DB) {
	handler := UserHandler{
		userRepo: infrastructure.NewPostgresUserRepository(db),
	}

	users := app.Group("/users")

	users.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Success")
	})

	users.Get("/register", handler.Register)
	// users.Post("/login", handler.Login)
}
