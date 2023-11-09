package repository

import (
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

func (r *GroupSQLite) SpaceGroups(userId, spaceId int) ([]model.StorageGroup, error) {
	var groups []model.StorageGroup

	query := fmt.Sprintf("SELECT g.id, g.name, g.size, g.numOfFree FROM '%s' g INNER JOIN %s sg ON g.id=sg.group_id INNER JOIN %s us ON sg.space_id=us.space_id WHERE us.user_id=$1 AND us.space_id=$2",
		groupTable, groupInSpaceTable, userSpacesTable)
	err := r.db.Select(&groups, query, userId, spaceId)

	return groups, err

}

func (r *GroupSQLite) GroupById(userId, spaceId, groupId int) (model.StorageGroup, error) {
	var group model.StorageGroup

	query := fmt.Sprintf("SELECT g.id, g.name, g.size, g.numOfFree FROM '%s' g INNER JOIN %s sg ON g.id=sg.group_id INNER JOIN %s us ON sg.space_id=us.space_id WHERE us.user_id=$1 AND us.space_id=$2 AND g.id=$3",
		groupTable, groupInSpaceTable, userSpacesTable)
	err := r.db.Select(&group, query, userId, spaceId, groupId)

	return group, err
}
