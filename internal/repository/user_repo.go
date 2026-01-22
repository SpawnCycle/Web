package repository

import (
	"context"
	"smaash-web/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	ReadAll(context.Context) ([]models.User, error)
	ReadByID(c context.Context, id uint) (*models.User, error)
	Create(c context.Context, user models.User) error
	Update(c context.Context, user models.User) error
	Delete(c context.Context, id uint) error
}

type GormUserRepo struct {
	DB *gorm.DB
}

func NewGormUserRepo() UserRepository {
	db := NewGormDBConn()
	return &GormUserRepo{DB: db}
}

func (u GormUserRepo) ReadAll(c context.Context) ([]models.User, error) {
	users, err := gorm.G[models.User](u.DB).Find(c)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u GormUserRepo) ReadByID(c context.Context, id uint) (*models.User, error) {
	return nil, nil
}

func (u GormUserRepo) Create(c context.Context, user models.User) error {
	return nil
}

func (u GormUserRepo) Update(c context.Context, user models.User) error {
	return nil
}

func (u GormUserRepo) Delete(c context.Context, id uint) error {
	return nil
}
