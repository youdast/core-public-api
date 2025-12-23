package http

import (
	"net/http"
	"strconv"
	"youdast/core-public-api/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UUsecase domain.UserUsecase
}

func NewUserHandler(u domain.UserUsecase) *UserHandler {
	return &UserHandler{
		UUsecase: u,
	}
}

func (h *UserHandler) Fetch(c *fiber.Ctx) error {
	users, err := h.UUsecase.Fetch(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := h.UUsecase.GetByID(c.Context(), uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (h *UserHandler) Store(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.UUsecase.Store(c.Context(), &user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(user)
}
