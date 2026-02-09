package services

import (
	"context"
	"smaash-web/internal/models"
	"smaash-web/internal/repository"
)

type LevelService interface {
	Create(context.Context, *models.Level) error
	ReadAll(context.Context) ([]models.Level, error)
	ReadByID(c context.Context, id uint) (*models.Level, error)
	Update(c context.Context, level *models.Level) error
	Delete(c context.Context, id uint) error
}

type LevelServiceImpl struct {
	levelRepo repository.LevelRepository
}

func NewLevelService(levelRepo repository.LevelRepository) LevelService {
	return &LevelServiceImpl{levelRepo: levelRepo}
}

func (l LevelServiceImpl) Create(c context.Context, level *models.Level) error {
	return l.levelRepo.Create(c, level)
}

func (l LevelServiceImpl) ReadAll(c context.Context) ([]models.Level, error) {
	return l.levelRepo.ReadAll(c)
}

func (l LevelServiceImpl) ReadByID(c context.Context, id uint) (*models.Level, error) {
	return l.levelRepo.ReadByID(c, id)
}

func (l LevelServiceImpl) Update(c context.Context, level *models.Level) error {
	return l.levelRepo.Update(c, level)
}

func (l LevelServiceImpl) Delete(c context.Context, id uint) error {
	return l.levelRepo.Delete(c, id)
}
