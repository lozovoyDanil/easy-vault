package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/uptrace/bun"
	"main.go/internal/model"
)

type SpaceSQLite struct {
	db *bun.DB
}

func NewSpaceSQLite(db *bun.DB) *SpaceSQLite {
	return &SpaceSQLite{db: db}
}

func (r *SpaceSQLite) SpaceBelongsToUser(userId, spaceId int) (int, error) {
	var count int

	count, err := r.db.NewSelect().
		Table(userSpacesTable).
		Where("us.user_id = ? AND us.space_id = ?", userId, spaceId).
		Count(context.Background())
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *SpaceSQLite) AllSpaces() ([]model.Space, error) {
	var spaces []model.Space

	err := r.db.NewSelect().
		Model(&spaces).
		Scan(context.Background())

	return spaces, err
}

func (r *SpaceSQLite) UserSpaces(id int) ([]model.Space, error) {
	var spaces []model.Space

	err := r.db.NewSelect().
		Model(&spaces).
		Join(fmt.Sprintf("INNER JOIN %s us ON s.id=us.space_id", userSpacesTable)).
		Where("us.user_id = ?", id).
		Scan(context.Background())

	return spaces, err
}

func (r *SpaceSQLite) SpaceById(spaceId int) (model.Space, error) {
	var space model.Space

	err := r.db.NewSelect().
		Model(&space).
		Where("id = ?", spaceId).
		Scan(context.Background())

	return space, err
}

func (r *SpaceSQLite) CreateSpace(userId int, space model.Space) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	_, err = tx.NewInsert().
		Model(&space).
		Exec(context.Background())
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.NewInsert().
		Model(&model.UserSpace{
			UserId:  userId,
			SpaceId: space.Id,
		}).
		Exec(context.Background())
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return space.Id, tx.Commit()
}

func (r *SpaceSQLite) UpdateSpace(spaceId int, input model.UpdateSpaceInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Addr != nil {
		setValues = append(setValues, fmt.Sprintf("addr=$%d", argId))
		args = append(args, *input.Addr)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s s SET %s FROM %s us WHERE s.id = us.space_id AND us.space_id = $%d",
		spaceTable, setQuery, userSpacesTable, argId)
	args = append(args, spaceId)
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *SpaceSQLite) DeleteSpace(spaceId int) error {
	_, err := r.db.NewDelete().
		Table(spaceTable).
		Where("id = ?", spaceId).
		Exec(context.Background())

	return err
}
