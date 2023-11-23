package service

import (
	"main.go/internal/model"
	"main.go/internal/repository"
)

type GroupService struct {
	repo repository.Group
}

func NewGroupService(repo repository.Group) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) SpaceGroups(spaceId int) ([]model.StorageGroup, error) {
	return s.repo.SpaceGroups(spaceId)
}

func (s *GroupService) GroupById(spaceId, groupId int) (model.StorageGroup, error) {
	return s.repo.GroupById(spaceId, groupId)
}

func (s *GroupService) CreateGroup(userId, spaceId int, group model.StorageGroup) error {
	group.Size = 0
	group.NumOfFree = 0

	return s.repo.CreateGroup(userId, spaceId, group)
}
