package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"main.go/internal/model"
)

var (
	ErrOwnershipViolation = errors.New("access forbiden or obj does not exist")
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
	UnitBelongsToUser(userId, unitId int) error

	GroupUnits(groupId int) ([]model.StorageUnit, error)
	UnitById(unitId int) (model.StorageUnit, error)
	CreateUnit(groupId int, unit model.StorageUnit) (int, error)
	DeleteUnit(unitId int) error
	UpdateUnit(unitId int, input model.UpdateUnitInput) error

	ReservedUnits(userId int) ([]model.StorageUnit, error)
}

type Group interface {
	GroupBelongsToUser(userId, groupId int) error

	SpaceGroups(spaceId int) ([]model.StorageGroup, error)
	GroupById(spaceId, groupId int) (model.StorageGroup, error)
	CreateGroup(userId, spaceId int, group model.StorageGroup) error
	UpdateGroup(groupId int, input model.UpdateGroupInput) error
	DeleteGroup(groupId int) error
}

type Space interface {
	SpaceBelongsToUser(userId, spaceId int) error

	AllSpaces() ([]model.Space, error)
	SpaceById(spaceId int) (model.Space, error)

	UserSpaces(id int) ([]model.Space, error)
	CreateSpace(userId int, space model.Space) (int, error)
	UpdateSpace(userId, spaceId int, input model.UpdateSpaceInput) error
	DeleteSpace(userId, spaceId int) error
}

type Repository struct {
	Authorization
	Subscription
	Unit
	Group
	Space
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQLite(db),
		Subscription:  NewSubRepository(db),
		// Unit:          NewUnitSQLite(db),
		Group: NewGroupSQLite(db),
		Space: NewSpaceSQLite(db),
	}
}
