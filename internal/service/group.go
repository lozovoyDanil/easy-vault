package service

import (
	"main.go/internal/model"
	"main.go/internal/repository"
)

type GroupService struct {
	repository.Group
	//* Необходима для проверки принадлежности Space пользователю
	repository.Space
}

func NewGroupService(repo *repository.Repository) *GroupService {
	return &GroupService{Group: repo.Group, Space: repo.Space}
}

func (s *GroupService) SpaceGroups(spaceId int) ([]model.StorageGroup, error) {
	return s.Group.SpaceGroups(spaceId)
}

func (s *GroupService) GroupById(spaceId, groupId int) (model.StorageGroup, error) {
	return s.Group.GroupById(spaceId, groupId)
}

func (s *GroupService) CreateGroup(userId, spaceId int, group model.StorageGroup) error {
	if err := s.Space.SpaceBelongsToUser(userId, spaceId); err != nil {
		return err
	}

	group.Size = 0
	group.NumOfFree = 0

	return s.Group.CreateGroup(userId, spaceId, group)
}

func (s *GroupService) UpdateGroup(userId, groupId int, input model.UpdateGroupInput) error {
	if err := s.Group.GroupBelongsToUser(userId, groupId); err != nil {
		return err
	}

	return s.Group.UpdateGroup(groupId, input)
}

func (s *GroupService) DeleteGroup(userId, groupId int) error {
	if err := s.Group.GroupBelongsToUser(userId, groupId); err != nil {
		return err
	}

	return s.Group.DeleteGroup(groupId)
}
