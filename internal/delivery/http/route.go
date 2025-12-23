package http

import (
	"github.com/gofiber/fiber/v2"
)

func NewUserHttpHandler(app *fiber.App, handler *UserHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")

	users.Get("/", handler.Fetch)
	users.Get("/:id", handler.GetByID)
	users.Post("/", handler.Store)
}
