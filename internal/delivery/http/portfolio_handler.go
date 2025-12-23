package http

import (
	"net/http"
	"youdast/core-public-api/internal/domain"
	"youdast/core-public-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type PortfolioHandler struct {
	PUsecase domain.PortfolioUsecase
}

func NewPortfolioHandler(p domain.PortfolioUsecase) *PortfolioHandler {
	return &PortfolioHandler{
		PUsecase: p,
	}
}

// GetProfile godoc
// @Summary Get portfolio profile
// @Description Get the main profile information
// @Tags portfolio
// @Accept json
// @Produce json
// @Success 200 {object} utils.ApiResponse{data=domain.Profile}
// @Failure 404 {object} utils.ApiResponse
// @Router /api/v1/portfolio/profile [get]
func (h *PortfolioHandler) GetProfile(c *fiber.Ctx) error {
	profile, err := h.PUsecase.GetProfile(c.Context())
	if err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Profile not found", nil)
	}
	return utils.SuccessResponse(c, "Profile fetched successfully", profile)
}

// GetSkills godoc
// @Summary Get skills
// @Description Get list of skills
// @Tags portfolio
// @Accept json
// @Produce json
// @Success 200 {object} utils.ApiResponse{data=[]domain.Skill}
// @Failure 500 {object} utils.ApiResponse
// @Router /api/v1/portfolio/skills [get]
func (h *PortfolioHandler) GetSkills(c *fiber.Ctx) error {
	skills, err := h.PUsecase.GetSkills(c.Context())
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch skills", err.Error())
	}
	return utils.SuccessResponse(c, "Skills fetched successfully", skills)
}

// GetProjects godoc
// @Summary Get projects
// @Description Get list of projects
// @Tags portfolio
// @Accept json
// @Produce json
// @Success 200 {object} utils.ApiResponse{data=[]domain.Project}
// @Failure 500 {object} utils.ApiResponse
// @Router /api/v1/portfolio/projects [get]
func (h *PortfolioHandler) GetProjects(c *fiber.Ctx) error {
	projects, err := h.PUsecase.GetProjects(c.Context())
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch projects", err.Error())
	}
	return utils.SuccessResponse(c, "Projects fetched successfully", projects)
}
