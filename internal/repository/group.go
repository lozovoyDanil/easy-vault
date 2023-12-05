package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"main.go/internal/model"
)

type GroupSQLite struct {
	db *sqlx.DB
}

func NewGroupSQLite(db *sqlx.DB) *GroupSQLite {
	return &GroupSQLite{db: db}
}

func (r *GroupSQLite) GroupBelongsToUser(userId, groupId int) error {
	var count int

	query := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s g
		INNER JOIN %s sg ON g.id=sg.group_id
		INNER JOIN %s us ON sg.space_id=us.space_id
		WHERE us.user_id=$1 AND g.id=$2`,
		groupTable, groupInSpaceTable, userSpacesTable)
	err := r.db.Get(&count, query, userId, groupId)
	if count == 0 {
		return ErrOwnershipViolation
	}

	return err
}

func (r *GroupSQLite) SpaceGroups(spaceId int) ([]model.StorageGroup, error) {
	var groups []model.StorageGroup

	query := fmt.Sprintf("SELECT g.id, g.name, g.size, g.numOfFree FROM '%s' g INNER JOIN %s sg ON g.id=sg.group_id INNER JOIN %s us ON sg.space_id=us.space_id WHERE us.space_id=$1",
		groupTable, groupInSpaceTable, userSpacesTable)
	err := r.db.Select(&groups, query, spaceId)

	return groups, err

}

func (r *GroupSQLite) GroupById(spaceId, groupId int) (model.StorageGroup, error) {
	var group model.StorageGroup

	query := fmt.Sprintf("SELECT g.id, g.name, g.size, g.numOfFree FROM '%s' g INNER JOIN %s sg ON g.id=sg.group_id INNER JOIN %s us ON sg.space_id=us.space_id WHERE us.space_id=$1 AND g.id=$2",
		groupTable, groupInSpaceTable, userSpacesTable)
	err := r.db.Select(&group, query, spaceId, groupId)

	return group, err
}

func (r *GroupSQLite) CreateGroup(userId, spaceId int, group model.StorageGroup) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s g(name, size, numOfFree) VALUES($1, $2, $3) RETURNING id", groupTable)
	row, err := tx.Exec(query, group.Name, group.Size, group.NumOfFree)
	if err != nil {
		tx.Rollback()
		return err
	}
	groupId, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s(space_id, group_id) SELECT space_id, $1 as group_id FROM User_Spaces WHERE space_id=$2 AND user_id=$3", groupInSpaceTable)
	_, err = tx.Exec(query, groupId, spaceId, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
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

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf("UPDATE %s g SET %s WHERE g.id=$%d",
		groupTable, setQuery, argId)
	args = append(args, groupId)
	_, err := r.db.Exec(query, args...)

	return err

}

func (r *GroupSQLite) DeleteGroup(groupId int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id=$1`,
		groupTable)
	_, err := r.db.Exec(query, groupId)

	return err
}
