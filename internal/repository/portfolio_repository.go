package repository

import (
	"context"
	"youdast/core-public-api/internal/domain"

	"gorm.io/gorm"
)

type portfolioRepository struct {
	Conn *gorm.DB
}

func NewPortfolioRepository(Conn *gorm.DB) domain.PortfolioRepository {
	return &portfolioRepository{Conn}
}

func (m *portfolioRepository) GetProfile(ctx context.Context) (res domain.Profile, err error) {
	// Assuming single profile for now, get the first one
	err = m.Conn.WithContext(ctx).First(&res).Error
	return
}

func (m *portfolioRepository) GetSkills(ctx context.Context) (res []domain.Skill, err error) {
	err = m.Conn.WithContext(ctx).Where("hidden = ?", false).Order("id asc").Find(&res).Error
	return
}

func (m *portfolioRepository) GetProjects(ctx context.Context) (res []domain.Project, err error) {
	err = m.Conn.WithContext(ctx).Preload("Skills").Order("\"order\" asc").Find(&res).Error
	return
}

func (m *portfolioRepository) UpdateProfile(ctx context.Context, p *domain.Profile) (err error) {
	// Upsert based on ID=1 usually for single profile
	if p.ID == 0 {
		p.ID = 1
	}
	err = m.Conn.WithContext(ctx).Save(p).Error
	return
}

func (m *portfolioRepository) AddSkill(ctx context.Context, s *domain.Skill) (err error) {
	err = m.Conn.WithContext(ctx).Create(s).Error
	return
}

func (m *portfolioRepository) AddProject(ctx context.Context, p *domain.Project) (err error) {
	err = m.Conn.WithContext(ctx).Create(p).Error
	return
}
