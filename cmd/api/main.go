package main

import (
	"log"
	"time"
	"youdast/core-public-api/config"
	_ "youdast/core-public-api/docs" // Import generated docs
	"youdast/core-public-api/internal/delivery/http"
	"youdast/core-public-api/internal/domain"
	"youdast/core-public-api/internal/repository"
	"youdast/core-public-api/internal/usecase"
	"youdast/core-public-api/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title Core Public API
// @version 1.0
// @description This is a Clean Architecture REST API service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
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
	if err := db.AutoMigrate(&domain.User{}, &domain.Profile{}, &domain.Skill{}, &domain.Project{}); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	// 3. Init Layers
	userRepo := repository.NewMysqlUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, 2*time.Second)
	userHandler := http.NewUserHandler(userUsecase)

	portfolioRepo := repository.NewPortfolioRepository(db)
	portfolioUsecase := usecase.NewPortfolioUsecase(portfolioRepo, 2*time.Second)
	portfolioHandler := http.NewPortfolioHandler(portfolioUsecase)

	// 4. Init Fiber
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Swagger Route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// 5. Setup Routes
	http.NewUserHttpHandler(app, userHandler)
	http.NewPortfolioHttpHandler(app, portfolioHandler)

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// 6. Start Server
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
