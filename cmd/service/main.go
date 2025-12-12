package main

import (
	"fmt"
	"yoyoj-go-rag/configs"
	app_database "yoyoj-go-rag/pkg/infrastructure/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg, err := configs.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Errorf("Failed to load config: %v\n", err)
		return
	}
	log.Infof("启动服务 %v\n", cfg)

	db, err := app_database.NewPostgresConnection(&cfg.Database.Postgres)
	if err != nil {
		log.Errorf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ErrorHandler:  globalErrorHandler(),
	})

	app.Use(recover.New())
	app.Use(logAllRequests())
	// app.Use(catchMiddleware) --- IGNORE ---
	// 其他的错误处理交给每个领域和模块自己去处理

	v1_app := app.Group("/api/v1")

	// Register user routes
	// user_interface.RegisterRoutes(v1_app, db.DB) --- IGNORE ---
	registerRoutes(v1_app, db)

	app.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
}
