package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/uptrace/bun"
	"main.go/internal/model"
)

type AuthSQLite struct {
	db *bun.DB
}

func NewAuthSQLite(db *bun.DB) *AuthSQLite {
	return &AuthSQLite{db: db}
}

func (r *AuthSQLite) CreateUser(user model.User) (int, error) {
	_, err := r.db.NewInsert().
		Model(&user).
		Exec(context.Background())

	return user.Id, err
}

func (r *AuthSQLite) GetUser(username, password string) (model.User, error) {
	var user model.User

	err := r.db.NewSelect().
		Model(&user).
		Where("email = ? AND password = ?", username, password).
		Scan(context.Background())

	return user, err
}

func (r *AuthSQLite) UserInfo(userId int) (model.User, error) {
	var user model.User

	err := r.db.NewSelect().
		Model(&user).
		Where("id = ?", userId).
		Scan(context.Background())

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
