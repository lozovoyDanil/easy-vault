package service

import (
	"main.go/internal/model"
	"main.go/internal/repository"
)

type SpaceService struct {
	repo repository.Space
}

func NewSpaceService(repo repository.Space) *SpaceService {
	return &SpaceService{repo: repo}
}

func (s *SpaceService) AllSpaces(filter model.SpaceFilter) ([]model.Space, error) {
	return s.repo.AllSpaces(filter)
}

func (s *SpaceService) UserSpaces(id int) ([]model.Space, error) {
	return s.repo.UserSpaces(id)
}

func (s *SpaceService) SpaceById(spaceId int) (model.Space, error) {
	return s.repo.SpaceById(spaceId)
}

func (s *SpaceService) CreateSpace(userId int, space model.Space) (int, error) {
	return s.repo.CreateSpace(userId, space)
}

func (s *SpaceService) UpdateSpace(userId, spaceId int, space model.UpdateSpaceInput) error {
	count, err := s.repo.SpaceBelongsToUser(userId, spaceId)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrOwnershipViolation
	}

	return s.repo.UpdateSpace(spaceId, space)
}

func (s *SpaceService) DeleteSpace(userId, spaceId int) error {
	count, err := s.repo.SpaceBelongsToUser(userId, spaceId)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrOwnershipViolation
	}

	return s.repo.DeleteSpace(spaceId)
}
