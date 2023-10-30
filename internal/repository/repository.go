package repository

import (
	"github.com/jmoiron/sqlx"
	"main.go/internal/model"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Subscription interface {
	SubByUserId(id int) (model.Subscription, error)
	CreateSub(model.Subscription) error
	UpdateSub(model.Subscription) error
}

type Unit interface {
}

type Group interface {
	SpaceGroups(userId, spaceId int) ([]model.StorageGroup, error)
	GroupById(userId, spaceId, groupId int) (model.StorageGroup, error)
}

type Space interface {
	UserSpaces(id int) ([]model.Space, error)
	SpaceById(userId, spaceId int) (model.Space, error)
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
