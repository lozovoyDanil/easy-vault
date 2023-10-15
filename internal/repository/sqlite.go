package repository

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

const (
	userTable         = "User"
	userUnitsTable    = "User_Unit"
	unitTable         = "Unit"
	unitInGroupTable  = "Unit_Group"
	groupTable        = "Group"
	groupInSpaceTable = "Group_Space"
	spaceTable        = "Space"
)

func NewSQLiteDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite", "./easy-vault.sqlite")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
