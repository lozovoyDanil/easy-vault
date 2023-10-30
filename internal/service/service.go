package service

import (
	"main.go/internal/model"
	"main.go/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Subscription interface {
	SubByUserId(id int) (model.Subscription, error)
	CreateSub(model.Subscription) error
	UpdateSub(model.Subscription) error
}

type Space interface {
	UserSpaces(id int) ([]model.Space, error)
	SpaceById(userId, spaceId int) (model.Space, error)
	CreateSpace(userId int, space model.Space) (int, error)
	UpdateSpace(userId, spaceId int, space model.UpdateSpaceInput) error
	DeleteSpace(userId, spaceId int) error
}

type Group interface {
	SpaceGroups(userId, spaceId int) ([]model.StorageGroup, error)
	GroupById(userId, spaceId, groupId int) (model.StorageGroup, error)
	// CreateGroup()
	// UpdateGroup()
	// DeleteGroup()
}

type Unit interface {
}

type Service struct {
	Authorization
	Subscription
	Unit
	Group
	Space
}

func NewServices(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Subscription:  NewSubService(repo),
		// Unit:          NewUnitService(repo),
		Group: NewGroupService(repo),
		Space: NewSpaceService(repo),
	}
}
