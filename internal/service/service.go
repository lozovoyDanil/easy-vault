package service

import (
	"errors"

	m "main.go/internal/model"
	"main.go/internal/repository"
)

var (
	//ERROR: user tries to access forbidden object.
	ErrOwnershipViolation = errors.New("access forbidden or object does not exist")
	//ERROR: user already banned.
	ErrUserAlreadyBanned = errors.New("user already banned")
	//ERROR: cannot delete user.
	ErrCannotDeleteUser = errors.New("cannot delete user")
	//ERROR: user tries to create new partnership, but he already has one.
	ErrPartnershipViolation = errors.New("user already has partnership")
	//ERROR: user tries to create new space, unit or calculate revenue,
	//but his partnership tier does not allow this action.
	ErrTierViolation = errors.New("your partnership tier does not allow this action")
	//ERROR: user has reached limit of spaces for his partnership tier.
	ErrSpacesLimitReached = errors.New("cannot create more spaces, try to upgrade your partnership tier")
	//ERROR: user has reached limit of units for his partnership tier.
	ErrUnitsLimitReached = errors.New("cannot create more units, try to upgrade your partnership tier")
)

// Partnership tiers.
const (
	FreeTier = iota
	ProTier
	EnterpriseTier
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

type Partnership interface {
	PartByUserId(id int) (m.Partnership, error)
	CreatePart(userId, tier int) error
	UpdatePart(userId, tier int) error
}

type Revenue interface {
	ManagerRevenue(userId int) (m.Revenue, error)
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
	CreateUnit(user m.UserIdentity, groupId int, unit m.StorageUnit) (int, error)
	UpdateUnit(user m.UserIdentity, unitId int, unit m.UnitInput) error
	DeleteUnit(user m.UserIdentity, unitId int) error
	ReservedUnits(userId int) ([]m.StorageUnit, error)
	UnitDetails(user m.UserIdentity, unitId int) (m.UnitDetails, error)
	ReserveUnit(userId, unitId int, reserveInfo m.UnitInput) error
}

type Service struct {
	Admin
	Authorization
	Partnership
	Revenue
	Space
	Group
	Unit
}

func NewServices(repo *repository.Repository) *Service {
	return &Service{
		Admin:         NewAdminService(repo),
		Authorization: NewAuthService(repo),
		Partnership:   NewPartService(repo),
		Revenue:       NewRevenueService(repo),
		Space:         NewSpaceService(repo),
		Group:         NewGroupService(repo),
		Unit:          NewUnitService(repo),
	}
}
