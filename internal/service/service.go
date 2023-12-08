package service

import (
	"main.go/internal/model"
	"main.go/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (model.UserIdentity, error)

	UserInfo(userId int) (model.User, error)
	EditUser(userId int, input model.UpdateUserInput) error
}

type Subscription interface {
	SubByUserId(id int) (model.Subscription, error)
	CreateSub(model.Subscription) error
	UpdateSub(model.Subscription) error
}

type Space interface {
	AllSpaces() ([]model.Space, error)

	UserSpaces(id int) ([]model.Space, error)
	SpaceById(spaceId int) (model.Space, error)
	CreateSpace(userId int, space model.Space) (int, error)
	UpdateSpace(userId, spaceId int, space model.UpdateSpaceInput) error
	DeleteSpace(userId, spaceId int) error
}

type Group interface {
	SpaceGroups(spaceId int) ([]model.StorageGroup, error)
	GroupById(groupId int) (model.StorageGroup, error)
	CreateGroup(userId, spaceId int, group model.StorageGroup) error
	UpdateGroup(userId, groupId int, group model.UpdateGroupInput) error
	DeleteGroup(userId, groupId int) error
}

type Unit interface {
	GroupUnits(userId, groupId int) ([]model.StorageUnit, error)
	UnitById(userId, unitId int) (model.StorageUnit, error)
	CreateUnit(userId, groupId int, unit model.StorageUnit) (int, error)
	UpdateUnit(userId, unitId int, unit model.UpdateUnitInput) error
	DeleteUnit(userId, unitId int) error

	ReservedUnits(userId int) ([]model.StorageUnit, error)
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
