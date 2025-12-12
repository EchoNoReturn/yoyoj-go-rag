package main

import (
	user_interface "yoyoj-go-rag/internal/user/interface"
	app_database "yoyoj-go-rag/pkg/infrastructure/database"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app fiber.Router, db *app_database.PostgresDB) {
	user_interface.RegisterRoutes(app, db.DB)
}
