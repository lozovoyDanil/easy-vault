package service

import (
	m "main.go/internal/model"
	"main.go/internal/repository"
)

// Limits for partnership tiers.
var SpacesPlanLimits = map[int]int{
	//Free tier.
	0: 1,
	//Pro tier.
	1: 5,
	//Enterprise tier.
	2: 1024,
}

type SpaceService struct {
	repo repository.Space
	repository.Partnership
}

func NewSpaceService(repo *repository.Repository) *SpaceService {
	return &SpaceService{
		repo:        repo.Space,
		Partnership: repo.Partnership,
	}
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
	userPart, err := s.Partnership.PartByUserId(userId)
	if err != nil {
		return 0, err
	}
	spacesCount, err := s.repo.ManagerSpacesCount(userId)
	if err != nil {
		return 0, err
	}
	//Check if user has reached limit of spaces for his partnership tier.
	if spacesCount >= SpacesPlanLimits[userPart.Tier] {
		return 0, ErrSpacesLimitReached
	}

	return s.repo.CreateSpace(userId, space)
}

func (s *SpaceService) UpdateSpace(user m.UserIdentity, spaceId int, space m.SpaceInput) error {
	//Check if user is manager of this space.
	if !s.repo.ManagerOwnsSpace(user.Id, spaceId) && user.Role != m.AdminRole {
		return ErrOwnershipViolation
	}

	return s.repo.UpdateSpace(spaceId, space)
}

func (s *SpaceService) DeleteSpace(user m.UserIdentity, spaceId int) error {
	//Check if user is manager of this space.
	if !s.repo.ManagerOwnsSpace(user.Id, spaceId) && user.Role != m.AdminRole {
		return ErrOwnershipViolation
	}

	return s.repo.DeleteSpace(spaceId)
}
