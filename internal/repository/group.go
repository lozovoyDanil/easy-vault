package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"main.go/internal/model"
)

type GroupSQLite struct {
	db *sqlx.DB
}

func NewGroupSQLite(db *sqlx.DB) *GroupSQLite {
	return &GroupSQLite{db: db}
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

	query := fmt.Sprintf("INSERT INTO %s g(name) VALUES($1, $2, $3) RETURNING id", groupTable)
	row := tx.QueryRow(query, group.Name, group.Size, group.NumOfFree)
	var groupId int
	if err := row.Scan(&groupId); err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s(space_id, group_id) SELECT space_id, $1 as group_id FROM User_Spaces WHERE space_id=$2 AND user_id=$3", groupInSpaceTable)
	res, err := tx.Exec(query, groupId, spaceId, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	// If res.RowsAffected() returns 0, this means that eather
	// space does not exist or user does not own it.
	if r, _ := res.RowsAffected(); r == 0 {
		tx.Rollback()
		return errors.New("access forbiden or object does not exist")
	}

	return tx.Commit()
}
