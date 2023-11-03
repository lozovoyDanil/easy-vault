package repository

import (
	"fmt"

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
	var id int

	query := fmt.Sprintf("INSERT INTO %s(fullName, email, password, role) VALUES($1, $2, $3, $4) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.FullName, user.Email, user.Password, user.Role)
	err := row.Scan(&id)

	return id, err
}

func (r *AuthSQLite) GetUser(username, password string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT id FROM %s u WHERE u.email=$1 AND u.password=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
