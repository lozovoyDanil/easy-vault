package service

import (
	m "main.go/internal/model"
	"main.go/internal/repository"
)

type GroupService struct {
	repository.Group
	//* Необходима для проверки принадлежности Space пользователю
	repository.Space
}

func NewGroupService(repo *repository.Repository) *GroupService {
	return &GroupService{
		Group: repo.Group,
		Space: repo.Space,
	}
}

func (s *GroupService) SpaceGroups(spaceId int) ([]m.StorageGroup, error) {
	return s.Group.SpaceGroups(spaceId)
}

func (s *GroupService) GroupById(groupId int) (m.StorageGroup, error) {
	return s.Group.GroupById(groupId)
}

func (s *GroupService) CreateGroup(user m.UserIdentity, spaceId int, group m.StorageGroup) error {
	count, err := s.Space.SpaceBelongsToUser(user.Id, spaceId)
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

func (s *GroupService) UpdateGroup(user m.UserIdentity, groupId int, input m.GroupInput) error {
	if !s.Group.ManagerOwnsGroup(user.Id, groupId) && user.Role != m.AdminRole {
		return ErrOwnershipViolation
	}

	return s.Group.UpdateGroup(groupId, input)
}

func (s *GroupService) DeleteGroup(user m.UserIdentity, groupId int) error {
	if !s.Group.ManagerOwnsGroup(user.Id, groupId) && user.Role != m.AdminRole {
		return ErrOwnershipViolation
	}

	return s.Group.DeleteGroup(groupId)
}
