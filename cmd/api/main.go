package main

import (
	"log"
	"time"
	"youdast/core-public-api/config"
	"youdast/core-public-api/internal/delivery/http"
	"youdast/core-public-api/internal/domain"
	"youdast/core-public-api/internal/repository"
	"youdast/core-public-api/internal/usecase"
	"youdast/core-public-api/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Connect to Database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migrate (for demo purposes)
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// 3. Init Layers
	userRepo := repository.NewMysqlUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, 2*time.Second)
	userHandler := http.NewUserHandler(userUsecase)

	// 4. Init Fiber
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// 5. Setup Routes
	http.NewUserHttpHandler(app, userHandler)

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// 6. Start Server
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
