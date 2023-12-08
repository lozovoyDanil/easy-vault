package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/uptrace/bun"
	"main.go/internal/model"
)

type GroupSQLite struct {
	db *bun.DB
}

func NewGroupSQLite(db *bun.DB) *GroupSQLite {
	return &GroupSQLite{db: db}
}

func (r *GroupSQLite) GroupBelongsToUser(userId, groupId int) error {
	var count int

	count, err := r.db.NewSelect().
		Table(groupTable).
		Join(fmt.Sprintf("INNER JOIN %s s ON s.id=g.space_id", spaceTable)).
		Join(fmt.Sprintf("INNER JOIN %s us ON s.id = us.space_id", userSpacesTable)).
		Where("us.user_id = ? AND g.id = ?", userId, groupId).
		Count(context.Background())
	if count == 0 {
		return ErrOwnershipViolation
	}

	return err
}

func (r *GroupSQLite) SpaceGroups(spaceId int) ([]model.StorageGroup, error) {
	var groups []model.StorageGroup

	err := r.db.NewSelect().
		Model(&groups).
		Join(fmt.Sprintf("INNER JOIN %s s ON s.id=g.space_id", spaceTable)).
		Where("s.id = ?", spaceId).
		Scan(context.Background())

	return groups, err

}

func (r *GroupSQLite) GroupById(spaceId, groupId int) (model.StorageGroup, error) {
	var group model.StorageGroup

	err := r.db.NewSelect().
		Model(&group).
		Join(fmt.Sprintf("INNER JOIN %s s ON s.id=g.space_id", spaceTable)).
		Where("s.id = ? AND g.id = ?", spaceId, groupId).
		Scan(context.Background())

	return group, err
}

func (r *GroupSQLite) CreateGroup(userId, spaceId int, group model.StorageGroup) error {
	_, err := r.db.NewInsert().
		Model(&group).
		Exec(context.Background())

	return err
}

func (r *GroupSQLite) UpdateGroup(groupId int, input model.UpdateGroupInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.PricePer != nil {
		setValues = append(setValues, fmt.Sprintf("pricePer=$%d", argId))
		args = append(args, *input.PricePer)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s g SET %s WHERE g.id=$%d",
		groupTable, setQuery, argId)
	args = append(args, groupId)
	_, err := r.db.Exec(query, args...)

	return err

}

func (r *GroupSQLite) DeleteGroup(groupId int) error {
	_, err := r.db.NewDelete().
		Table(groupTable).
		Where("id = ?", groupId).
		Exec(context.Background())

	return err
}
