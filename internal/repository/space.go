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

func (r *SpaceSQLite) ManagerOwnsSpace(userId, spaceId int) bool {
	var count int

	count, err := r.db.NewSelect().
		Table(userSpacesTable).
		Where("us.user_id = ? AND us.space_id = ?", userId, spaceId).
		Count(context.Background())

	return count > 0 && err == nil
}

func (r *SpaceSQLite) AllSpaces(filter model.SpaceFilter) ([]model.Space, error) {
	var spaces []model.Space
	filterValues := make([]string, 0)
	args := make([]any, 0)

	if filter.Name != nil {
		filterValues = append(filterValues, "name LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", *filter.Name))
	}
	if filter.Addr != nil {
		filterValues = append(filterValues, "addr LIKE ?")
		args = append(args, fmt.Sprintf("%%%s%%", *filter.Addr))
	}
	if filter.MinSize != nil {
		filterValues = append(filterValues, "size >= ?")
		args = append(args, *filter.MinSize)
	}
	if filter.MaxSize != nil {
		filterValues = append(filterValues, "size <= ?")
		args = append(args, *filter.MaxSize)
	}

	if len(filterValues) == 1 {
		filterValues = append([]string{""}, filterValues...)
	}
	filterQuery := strings.Join(filterValues, " AND ")
	query := fmt.Sprintf("SELECT * FROM %s WHERE 1=1%s", spaceTable, filterQuery)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var space model.Space
		err := rows.Scan(
			&space.Id, &space.Name, &space.Addr,
			&space.Size, &space.NumOfGroups, &space.NumOfFree,
		)
		if err != nil {
			return nil, err
		}

		spaces = append(spaces, space)
	}

	return spaces, nil
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

func (r *SpaceSQLite) UpdateSpace(spaceId int, input model.SpaceInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)

	if input.Name != nil {
		setValues = append(setValues, "name=?")
		args = append(args, *input.Name)
	}
	if input.Addr != nil {
		setValues = append(setValues, "addr=?")
		args = append(args, *input.Addr)
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s AS s SET %s FROM %s us WHERE s.id = us.space_id AND us.space_id = ?",
		spaceTable, setQuery, userSpacesTable)
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
