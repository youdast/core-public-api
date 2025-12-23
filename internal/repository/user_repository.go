package repository

import (
	"context"
	"youdast/core-public-api/internal/domain"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{Conn}
}

func (m *mysqlUserRepository) Fetch(ctx context.Context) (res []domain.User, err error) {
	err = m.Conn.WithContext(ctx).Find(&res).Error
	return
}

func (m *mysqlUserRepository) GetByID(ctx context.Context, id uint) (res domain.User, err error) {
	err = m.Conn.WithContext(ctx).First(&res, id).Error
	return
}

func (m *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (res domain.User, err error) {
	err = m.Conn.WithContext(ctx).Where("email = ?", email).First(&res).Error
	return
}

func (m *mysqlUserRepository) Store(ctx context.Context, u *domain.User) (err error) {
	err = m.Conn.WithContext(ctx).Create(u).Error
	return
}

func (m *mysqlUserRepository) Update(ctx context.Context, u *domain.User) (err error) {
	err = m.Conn.WithContext(ctx).Save(u).Error
	return
}

func (m *mysqlUserRepository) Delete(ctx context.Context, id uint) (err error) {
	err = m.Conn.WithContext(ctx).Delete(&domain.User{}, id).Error
	return
}
