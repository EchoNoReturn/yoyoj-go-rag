package main

import (
	"fmt"
	"yoyoj-go-rag/configs"
	user_interface "yoyoj-go-rag/internal/user/interface"
	app_database "yoyoj-go-rag/pkg/infrastructure/database"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app fiber.Router, db *app_database.PostgresDB) {
	user_interface.RegisterRoutes(app, db.DB)
}

func main() {
	cfg, err := configs.LoadConfig("configs/config.yaml")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}
	fmt.Printf("启动服务 %v\n", cfg)

	db, err := app_database.NewPostgresConnection(&cfg.Database.Postgres)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	v1_app := app.Group("/api/v1")

	// Register user routes
	// user_interface.RegisterRoutes(v1_app, db.DB) --- IGNORE ---
	registerRoutes(v1_app, db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
}
