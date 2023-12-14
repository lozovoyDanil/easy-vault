package repository

import (
	"github.com/uptrace/bun"
	"main.go/internal/model"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)

	UserInfo(userId int) (model.User, error)
	EditUser(input model.UpdateUserInput) error
}

type Subscription interface {
	SubByUserId(id int) (model.Subscription, error)
	CreateSub(model.Subscription) error
	UpdateSub(model.Subscription) error
}

type Unit interface {
	UnitBelongsToUser(userId, unitId int) (int, error)

	GroupUnits(groupId int) ([]model.StorageUnit, error)
	UnitById(unitId int) (model.StorageUnit, error)
	CreateUnit(unit model.StorageUnit) (int, error)
	UpdateUnit(unitId int, input model.UpdateUnitInput) error
	DeleteUnit(unitId int) error

	ReservedUnits(userId int) ([]model.StorageUnit, error)
	UnitDetails(unitId int) (model.StorageUnit, error)
	ReserveUnit(userId, unitId int, reservInfo model.UpdateUnitInput) error
}

type Group interface {
	GroupBelongsToUser(userId, groupId int) (int, error)

	SpaceGroups(spaceId int) ([]model.StorageGroup, error)
	GroupById(groupId int) (model.StorageGroup, error)
	CreateGroup(group model.StorageGroup) error
	UpdateGroup(groupId int, input model.UpdateGroupInput) error
	DeleteGroup(groupId int) error
}

type Space interface {
	SpaceBelongsToUser(userId, spaceId int) (int, error)

	AllSpaces() ([]model.Space, error)
	SpaceById(spaceId int) (model.Space, error)

	UserSpaces(id int) ([]model.Space, error)
	CreateSpace(userId int, space model.Space) (int, error)
	UpdateSpace(spaceId int, input model.UpdateSpaceInput) error
	DeleteSpace(spaceId int) error
}

type Repository struct {
	Authorization
	Subscription
	Unit
	Group
	Space
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQLite(db),
		// Subscription:  NewSubRepository(db),
		// Unit:          NewUnitSQLite(db),
		Group: NewGroupSQLite(db),
		Space: NewSpaceSQLite(db),
	}
}
