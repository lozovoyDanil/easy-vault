package service

import (
	m "main.go/internal/model"
	"main.go/internal/repository"
)

type SpaceService struct {
	repo repository.Space
}

func NewSpaceService(repo repository.Space) *SpaceService {
	return &SpaceService{repo: repo}
}

func (s *SpaceService) AllSpaces(filter m.SpaceFilter) ([]m.Space, error) {
	return s.repo.AllSpaces(filter)
}

func (s *SpaceService) UserSpaces(id int) ([]m.Space, error) {
	return s.repo.UserSpaces(id)
}

func (s *SpaceService) SpaceById(spaceId int) (m.Space, error) {
	return s.repo.SpaceById(spaceId)
}

func (s *SpaceService) CreateSpace(userId int, space m.Space) (int, error) {
	return s.repo.CreateSpace(userId, space)
}

func (s *SpaceService) UpdateSpace(user m.UserIdentity, spaceId int, space m.SpaceInput) error {
	if !s.repo.ManagerOwnsSpace(user.Id, spaceId) && user.Role != m.AdminRole {
		return ErrOwnershipViolation
	}

	return s.repo.UpdateSpace(spaceId, space)
}

func (s *SpaceService) DeleteSpace(user m.UserIdentity, spaceId int) error {
	if !s.repo.ManagerOwnsSpace(user.Id, spaceId) && user.Role != m.AdminRole {
		return ErrOwnershipViolation
	}

	return s.repo.DeleteSpace(spaceId)
}
