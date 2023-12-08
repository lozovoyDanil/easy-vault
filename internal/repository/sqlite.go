package repository

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
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

func NewSQLiteDB() (*bun.DB, error) {
	db, err := sql.Open("sqlite", "./easy-vault.sqlite")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	bunDB := bun.NewDB(db, sqlitedialect.New())

	return bunDB, nil
}
