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

func (s *SpaceService) AllSpaces() ([]model.Space, error) {
	return s.repo.AllSpaces()
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
	if err := s.repo.SpaceBelongsToUser(userId, spaceId); err != nil {
		return err
	}

	return s.repo.UpdateSpace(userId, spaceId, space)
}

func (s *SpaceService) DeleteSpace(userId, spaceId int) error {
	if err := s.repo.SpaceBelongsToUser(userId, spaceId); err != nil {
		return err
	}

	return s.repo.DeleteSpace(userId, spaceId)
}
