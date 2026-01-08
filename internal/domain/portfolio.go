package domain

import "context"

type Profile struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
	Github      string `json:"github"`
	LinkedIn    string `json:"linkedin"`
	Resume      string `json:"resume"`
	Email       string `json:"email"`
}

type Skill struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Project struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Role        string  `json:"role"`
	Description string  `json:"description"`
	Skills      []Skill `json:"tech_stacks" gorm:"many2many:tech_stacks;"`
	SourceLink  string  `json:"source_link"`
	DemoLink    string  `json:"demo_link"`
	Order       int     `json:"order"`
}

type PortfolioRepository interface {
	GetProfile(ctx context.Context) (Profile, error)
	GetSkills(ctx context.Context) ([]Skill, error)
	GetProjects(ctx context.Context) ([]Project, error)

	// Admin methods (optional for now, but good to have in interface)
	UpdateProfile(ctx context.Context, p *Profile) error
	AddSkill(ctx context.Context, s *Skill) error
	AddProject(ctx context.Context, p *Project) error
}

type PortfolioUsecase interface {
	GetProfile(ctx context.Context) (Profile, error)
	GetSkills(ctx context.Context) ([]Skill, error)
	GetProjects(ctx context.Context) ([]Project, error)
}
