package usecase

import (
	"context"
	"time"
	"youdast/core-public-api/internal/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (a *userUsecase) Fetch(c context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.userRepo.Fetch(ctx)
}

func (a *userUsecase) GetByID(c context.Context, id uint) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.userRepo.GetByID(ctx, id)
}

func (a *userUsecase) Store(c context.Context, u *domain.User) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	// Here you can add business logic, validation, hashing password, etc.
	return a.userRepo.Store(ctx, u)
}

func (a *userUsecase) Update(c context.Context, u *domain.User) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.userRepo.Update(ctx, u)
}

func (a *userUsecase) Delete(c context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	return a.userRepo.Delete(ctx, id)
}
