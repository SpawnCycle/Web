package services

import (
	"context"
	"smaash-web/internal/models"
	"smaash-web/internal/repository"
)

type UserStats interface {
	ReadAllUsers(context.Context) ([]models.User, error)
	ReadUserByID(c context.Context, id uint) (*models.User, error)
}

type UserStatsService struct {
	userRepo repository.UserRepository
}

func NewUserStatsService(userRepo repository.UserRepository) UserStats {
	return UserStatsService{userRepo: userRepo}
}

func (u UserStatsService) ReadAllUsers(c context.Context) ([]models.User, error) {
	return u.userRepo.ReadAll(c)
}

func (u UserStatsService) ReadUserByID(c context.Context, id uint) (*models.User, error) {
	return u.userRepo.ReadByID(c, id)
}
