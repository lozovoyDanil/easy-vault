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

func (s *GroupService) SpaceGroups(userId, spaceId int) ([]model.StorageGroup, error) {
	return s.repo.SpaceGroups(userId, spaceId)
}

func (s *GroupService) GroupById(userId, spaceId, groupId int) (model.StorageGroup, error) {
	return s.repo.GroupById(userId, spaceId, groupId)
}
