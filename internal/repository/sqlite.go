package repository

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

const (
	userTable         = "User"
	userUnitsTable    = "User_Units"
	unitTable         = "Unit"
	unitInGroupTable  = "Group_Units"
	groupTable        = "Group"
	groupInSpaceTable = "Space_Groups"
	spaceTable        = "Space"
	userSpacesTable   = "User_Spaces"
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
