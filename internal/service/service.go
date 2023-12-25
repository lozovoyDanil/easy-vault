package service

import (
	"errors"

	m "main.go/internal/model"
	"main.go/internal/repository"
)

var (
	ErrOwnershipViolation = errors.New("access forbidden or obj does not exist")
	//ERROR: user already banned.
	ErrUserAlreadyBanned = errors.New("user already banned")
	//ERROR: cannot delete user.
	ErrCannotDeleteUser = errors.New("cannot delete user")
)

type Admin interface {
	AllUsers() ([]m.User, error)
	BanUser(id int) error
	DeleteUser(id int) error
}

type Authorization interface {
	CreateUser(user m.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (m.UserIdentity, error)
	UserInfo(userId int) (m.User, error)
	EditUser(userId int, input m.UpdateUserInput) error
}

type Subscription interface {
	SubByUserId(id int) (m.Subscription, error)
	CreateSub(m.Subscription) error
	UpdateSub(m.Subscription) error
}

type Space interface {
	AllSpaces(filter m.SpaceFilter) ([]m.Space, error)
	UserSpaces(id int) ([]m.Space, error)
	SpaceById(spaceId int) (m.Space, error)
	CreateSpace(userId int, space m.Space) (int, error)
	UpdateSpace(user m.UserIdentity, spaceId int, space m.SpaceInput) error
	DeleteSpace(user m.UserIdentity, spaceId int) error
}

type Group interface {
	SpaceGroups(spaceId int) ([]m.StorageGroup, error)
	GroupById(groupId int) (m.StorageGroup, error)
	CreateGroup(user m.UserIdentity, spaceId int, group m.StorageGroup) error
	UpdateGroup(user m.UserIdentity, groupId int, group m.GroupInput) error
	DeleteGroup(user m.UserIdentity, groupId int) error
}
type Unit interface {
	GroupUnits(user m.UserIdentity, groupId int) ([]m.StorageUnit, error)
	UnitById(userId, unitId int) (m.StorageUnit, error)
	CreateUnit(user m.UserIdentity, groupId int, unit m.StorageUnit) (int, error)
	UpdateUnit(user m.UserIdentity, unitId int, unit m.UnitInput) error
	DeleteUnit(user m.UserIdentity, unitId int) error
	ReservedUnits(userId int) ([]m.StorageUnit, error)
	UnitDetails(user m.UserIdentity, unitId int) (m.UnitDetails, error)
	ReserveUnit(userId, unitId int, reservInfo m.UnitInput) error
}

type Service struct {
	Admin
	Authorization
	// Subscription
	Unit
	Group
	Space
}

func NewServices(repo *repository.Repository) *Service {
	return &Service{
		Admin:         NewAdminService(repo),
		Authorization: NewAuthService(repo),
		// Subscription:  NewSubService(repo),
		Unit:  NewUnitService(repo),
		Group: NewGroupService(repo),
		Space: NewSpaceService(repo),
	}
}
