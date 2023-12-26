package repository

import (
	"context"

	"github.com/uptrace/bun"
	"main.go/internal/model"
)

type PartSQLite struct {
	db *bun.DB
}

func NewPartSQLite(db *bun.DB) *PartSQLite {
	return &PartSQLite{db: db}
}

func (r *PartSQLite) PartByUserId(id int) (model.Partnership, error) {
	var part model.Partnership

	err := r.db.NewSelect().
		Model(&part).
		Where("user_id = ?", id).
		Scan(context.Background())

	return part, err
}

func (r *PartSQLite) CreatePart(input model.Partnership) error {
	_, err := r.db.NewInsert().
		Model(&input).
		Ignore().On("CONFLICT (user_id) DO NOTHING").
		Exec(context.Background())

	return err
}

func (r *PartSQLite) UpdatePart(input model.Partnership) error {
	_, err := r.db.NewUpdate().
		Model(&input).
		Where("user_id = ?", input.UserId).
		Exec(context.Background())

	return err
}
