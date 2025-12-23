package http

import (
	"net/http"
	"strconv"
	"youdast/core-public-api/internal/domain"
	"youdast/core-public-api/pkg/utils"

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

// Fetch godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} utils.ApiResponse{data=[]domain.User}
// @Failure 500 {object} utils.ApiResponse
// @Router /api/v1/users [get]
func (h *UserHandler) Fetch(c *fiber.Ctx) error {
	users, err := h.UUsecase.Fetch(c.Context())
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch users", err.Error())
	}
	return utils.SuccessResponse(c, "Users fetched successfully", users)
}

// GetByID godoc
// @Summary Get user by ID
// @Description Get details of a specific user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} utils.ApiResponse{data=domain.User}
// @Failure 400 {object} utils.ApiResponse
// @Failure 404 {object} utils.ApiResponse
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID", nil)
	}

	user, err := h.UUsecase.GetByID(c.Context(), uint(id))
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "User not found", err.Error())
	}
	return utils.SuccessResponse(c, "User fetched successfully", user)
}

// Store godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.User true "User Payload"
// @Success 201 {object} utils.ApiResponse{data=domain.User}
// @Failure 400 {object} utils.ApiResponse
// @Failure 500 {object} utils.ApiResponse
// @Router /api/v1/users [post]
func (h *UserHandler) Store(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid input", err.Error())
	}

	if err := h.UUsecase.Store(c.Context(), &user); err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user", err.Error())
	}
	return utils.SuccessResponse(c, "User created successfully", user)
}
