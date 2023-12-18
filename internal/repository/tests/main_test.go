package repository_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"main.go/internal/model"
)

var testDB *bun.DB

func TestMain(m *testing.M) {
	err := SetEnv()
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.Close()

	os.Exit(m.Run())
}

func SetEnv() error {
	db, err := sql.Open("sqlite", "file::memory:?cache=shared")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	testDB = bun.NewDB(db, sqlitedialect.New())
	testDB.NewCreateTable().Model((*model.User)(nil)).Exec(context.Background())

	return nil
}
