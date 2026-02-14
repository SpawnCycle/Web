package services

import (
	"context"
	"smaash-web/internal/models"
	"smaash-web/internal/repository"
)

type playerProfileServiceImpl struct {
	repo repository.PlayerProfileRepository
}

type PlayerProfileService interface {
	CreateProfile(c context.Context, profile *models.PlayerProfile) error
	GetProfileByID(c context.Context, id uint) (*models.PlayerProfile, error)
	GetProfileByUserID(c context.Context, userId uint) (*models.PlayerProfile, error)
	UpdateProfile(ctx context.Context, profile *models.PlayerProfile) error
}

func NewPlayerProfileService(repo repository.PlayerProfileRepository) PlayerProfileService {
	return &playerProfileServiceImpl{repo: repo}
}

func (s *playerProfileServiceImpl) CreateProfile(ctx context.Context, profile *models.PlayerProfile) error {

	// Attempt to create the profile in the database
	err := s.repo.Create(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}

func (s *playerProfileServiceImpl) GetProfileByUserID(ctx context.Context, userID uint) (*models.PlayerProfile, error) {
	return s.repo.ReadByUserId(ctx, userID)
}
func (s *playerProfileServiceImpl) GetProfileByID(c context.Context, id uint) (*models.PlayerProfile, error) {
	return s.repo.ReadById(c, id)
}

func (s *playerProfileServiceImpl) UpdateProfile(ctx context.Context, profile *models.PlayerProfile) error {
	return s.repo.Update(ctx, profile)
}
