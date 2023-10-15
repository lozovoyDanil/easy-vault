package repository

import (
	"github.com/jmoiron/sqlx"
	"main.go/internal/model"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Unit interface {
}

type Group interface {
}

type Space interface {
}

type Repository struct {
	Authorization
	Unit
	Group
	Space
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQLite(db),
		// Unit:          NewUnitSQLite(db),
		// Group:         NewGroupSQLite(db),
		// Space:         NewSpaceSQLite(db),
	}
}
