package repository

import (
	"context"

	"github.com/uptrace/bun"
	"main.go/internal/model"
)

type AdminSQLite struct {
	db *bun.DB
}

func NewAdminSQLite(db *bun.DB) *AdminSQLite {
	return &AdminSQLite{db: db}
}

func (r *AdminSQLite) AllUsers() ([]model.User, error) {
	var users []model.User

	err := r.db.NewSelect().
		Model(&users).
		Column("fullName", "email", "role").
		Where("u.role != ?", "admin").
		Scan(context.Background())

	return users, err
}

func (r *AdminSQLite) BanUser(id int) error {
	_, err := r.db.NewUpdate().
		Table(userTable).
		Column("isBanned").
		Set("isBanned = ?", true).
		Where("id = ?", id).
		Exec(context.Background())

	return err
}

func (r *AdminSQLite) DeleteUser(id int) error {
	_, err := r.db.NewDelete().
		Table(userTable).
		Where("id = ?", id).
		Exec(context.Background())

	return err
}
