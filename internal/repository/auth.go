package repository

import (
	"github.com/jmoiron/sqlx"
	"main.go/internal/model"
)

type AuthSQLite struct {
	db *sqlx.DB
}

func NewAuthSQLite(db *sqlx.DB) *AuthSQLite {
	return &AuthSQLite{db: db}
}

func (r *AuthSQLite) CreateUser(user model.User) (int, error) {
	return 0, nil
}

func (r *AuthSQLite) GetUser(username, password string) (model.User, error) {
	return model.User{}, nil
}
