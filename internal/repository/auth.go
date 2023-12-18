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

func (r *AuthSQLite) UserIsBanned(userId int) (bool, error) {
	var isBanned bool

	err := r.db.NewSelect().
		Model((*model.User)(nil)).
		Column("isBanned").
		Where("id = ?", userId).
		Scan(context.Background(), &isBanned)
	if err != nil {
		return false, err
	}

	return isBanned, nil
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

func (r *AuthSQLite) EditUser(userId int, input model.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	if input.FullName != nil {
		setValues = append(setValues, fmt.Sprintf("fullName='$%d'", argId))
		args = append(args, *input.FullName)
		argId++
	}
	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email='$%d'", argId))
		args = append(args, *input.Email)
		argId++
	}
	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password='$%d'", argId))
		args = append(args, *input.Password)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = '$%d'",
		userTable, setQuery, argId)
	args = append(args, userId)
	_, err := r.db.Exec(query, args...)

	return err
}
