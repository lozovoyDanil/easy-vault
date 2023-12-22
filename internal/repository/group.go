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

func (r *GroupSQLite) ManagerOwnsGroup(managerId, groupId int) bool {
	var count int

	count, err := r.db.NewSelect().
		Model(&model.StorageGroup{}).
		Join(fmt.Sprintf("INNER JOIN %s s ON s.id=g.space_id", spaceTable)).
		Join(fmt.Sprintf("INNER JOIN %s us ON us.id=s.manager_id", userSpacesTable)).
		Where("g.id = ? AND us.user_id = ?", groupId, managerId).
		Count(context.Background())

	return count > 0 && err == nil
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

func (r *GroupSQLite) GroupById(groupId int) (model.StorageGroup, error) {
	var group model.StorageGroup

	err := r.db.NewSelect().
		Model(&group).
		Where("g.id = ?", groupId).
		Scan(context.Background())

	return group, err
}

func (r *GroupSQLite) CreateGroup(group model.StorageGroup) error {
	_, err := r.db.NewInsert().
		Model(&group).
		Exec(context.Background())

	return err
}

func (r *GroupSQLite) UpdateGroup(groupId int, input model.GroupInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)

	if input.Name != nil {
		setValues = append(setValues, "name=?")
		args = append(args, *input.Name)
	}

	if input.Price != nil {
		setValues = append(setValues, "price=?")
		args = append(args, *input.Price)
	}

	if input.PricePer != nil {
		setValues = append(setValues, "pricePer=?")
		args = append(args, *input.PricePer)
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s AS g SET %s WHERE id=?", groupTable, setQuery)
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
