package repository

import (
	"github.com/uptrace/bun"
	"main.go/internal/model"
)

type Admin interface {
	AllUsers() ([]model.User, error)
	BanUser(id int) error
	DeleteUser(id int) error
}

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)

	UserInfo(userId int) (model.User, error)
	EditUser(userId int, input model.UpdateUserInput) error
	UserIsBanned(userId int) (bool, error)
}

type Partnership interface {
	PartByUserId(id int) (model.Partnership, error)
	CreatePart(model.Partnership) error
	UpdatePart(model.Partnership) error
}

type Unit interface {
	UnitOwnerId(unitId int) (int, error)
	ManagerOwnsUnit(managerId, unitId int) bool
	ManagerUnitsCount(managerId int) (int, error)

	GroupUnits(groupId int) ([]model.StorageUnit, error)
	UnitById(unitId int) (model.StorageUnit, error)
	CreateUnit(unit model.StorageUnit) (int, error)
	UpdateUnit(unitId int, input model.UnitInput) error
	DeleteUnit(unitId int) error

	ReservedUnits(userId int) ([]model.StorageUnit, error)
	LogHistory(log model.UnitHistory) error
	UnitHistory(unitId int) ([]model.UnitHistory, error)
}

type Group interface {
	ManagerOwnsGroup(managerId, groupId int) bool

	SpaceGroups(spaceId int) ([]model.StorageGroup, error)
	GroupById(groupId int) (model.StorageGroup, error)
	CreateGroup(group model.StorageGroup) error
	UpdateGroup(groupId int, input model.GroupInput) error
	DeleteGroup(groupId int) error
}

type Space interface {
	ManagerOwnsSpace(userId, spaceId int) bool
	ManagerSpacesCount(userId int) (int, error)

	AllSpaces(filter model.SpaceFilter) ([]model.Space, error)
	SpaceById(spaceId int) (model.Space, error)

	UserSpaces(id int) ([]model.Space, error)
	CreateSpace(userId int, space model.Space) (int, error)
	UpdateSpace(spaceId int, input model.SpaceInput) error
	DeleteSpace(spaceId int) error
}

type Repository struct {
	Admin
	Authorization
	Partnership
	Unit
	Group
	Space
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{
		Admin:         NewAdminSQLite(db),
		Authorization: NewAuthSQLite(db),
		Partnership:   NewPartSQLite(db),
		Unit:          NewUnitSQLite(db),
		Group:         NewGroupSQLite(db),
		Space:         NewSpaceSQLite(db),
	}
}
