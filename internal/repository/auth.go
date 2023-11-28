package repository

import (
	"fmt"
	"strings"

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

func (r *AuthSQLite) UserInfo(userId int) (model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT fullName, email, password FROM %s u WHERE id = $1", userTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *AuthSQLite) EditUser(input model.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	if input.FullName != nil {
		setValues = append(setValues, fmt.Sprintf("fullName=$%d", argId))
		args = append(args, *input.FullName)
		argId++
	}
	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}
	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, *input.Password)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s u SET %s WHERE u.id = $%d",
		userTable, setQuery, argId)
	args = append(args, input.Id)
	_, err := r.db.Exec(query, args...)

	return err

}
