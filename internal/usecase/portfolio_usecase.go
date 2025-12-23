package usecase

import (
	"context"
	"time"
	"youdast/core-public-api/internal/domain"
)

type portfolioUsecase struct {
	repo           domain.PortfolioRepository
	contextTimeout time.Duration
}

func NewPortfolioUsecase(r domain.PortfolioRepository, timeout time.Duration) domain.PortfolioUsecase {
	return &portfolioUsecase{
		repo:           r,
		contextTimeout: timeout,
	}
}

func (a *portfolioUsecase) GetProfile(c context.Context) (domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.repo.GetProfile(ctx)
}

func (a *portfolioUsecase) GetSkills(c context.Context) ([]domain.Skill, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.repo.GetSkills(ctx)
}

func (a *portfolioUsecase) GetProjects(c context.Context) ([]domain.Project, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.repo.GetProjects(ctx)
}
