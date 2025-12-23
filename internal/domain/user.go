package domain

import "context"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Store(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id uint) error
}

type UserUsecase interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	Store(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id uint) error
}
