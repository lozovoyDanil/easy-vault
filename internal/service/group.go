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

func (s *GroupService) GroupById(groupId int) (model.StorageGroup, error) {
	return s.Group.GroupById(groupId)
}

func (s *GroupService) CreateGroup(userId, spaceId int, group model.StorageGroup) error {
	count, err := s.Space.SpaceBelongsToUser(userId, spaceId)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrOwnershipViolation
	}

	group.Size = 0
	group.NumOfFree = 0

	return s.Group.CreateGroup(group)
}

func (s *GroupService) UpdateGroup(userId, groupId int, input model.UpdateGroupInput) error {
	count, err := s.Group.GroupBelongsToUser(userId, groupId)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrOwnershipViolation
	}

	return s.Group.UpdateGroup(groupId, input)
}

func (s *GroupService) DeleteGroup(userId, groupId int) error {
	count, err := s.Group.GroupBelongsToUser(userId, groupId)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrOwnershipViolation
	}

	return s.Group.DeleteGroup(groupId)
}
