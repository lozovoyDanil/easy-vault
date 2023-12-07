package controller

import (
	"os"

	"github.com/jmoiron/sqlx"
	"main.go/internal/controller"
	"main.go/internal/repository"
	"main.go/internal/service"
	_ "modernc.org/sqlite"
)

type Environment struct {
	Handler *controller.Handler
	DB      *sqlx.DB
}

func InitEnv() (*Environment, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(db)
	services := service.NewServices(repo)
	handler := controller.NewHandler(services)

	return &Environment{Handler: handler, DB: db}, nil
}

func (env *Environment) Remove() error {
	err := env.DB.Close()
	if err != nil {
		return err
	}

	err = os.Remove("./test-db.sqlite")
	if err != nil {
		return err
	}

	return nil
}

func openDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite", "./test-db.sqlite")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// err = createTables(db)
	// if err != nil {
	// 	return nil, err
	// }

	// err = insertTestRows(db)
	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}

// func createTableUser(db *sqlx.DB) error {
// 	_, err := db.Exec(`
// 		CREATE TABLE IF NOT EXISTS User (
// 			id INTEGER PRIMARY KEY AUTOINCREMENT,
// 			fullName TEXT,
// 			email TEXT,
// 			password TEXT,
// 			role TEXT
// 		)
// 	`)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func createTables(db *sqlx.DB) error {
// 	_, err := db.Exec(`
// 		CREATE TABLE IF NOT EXISTS "Space" (
// 			"id"	INTEGER,
// 			"name"	TEXT NOT NULL,
// 			"addr"	TEXT,
// 			"numOfGroups"	INTEGER,
// 			"size"	INTEGER NOT NULL,
// 			"numOfFree"	TEXT,
// 			PRIMARY KEY("id" AUTOINCREMENT)
// 		);
// 	`)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func insertTestRows(db *sqlx.DB) error {
// 	_, err := db.Exec(`
// 		INSERT INTO User (fullName, email, password, role)
// 		VALUES
// 			('Daniil', 'daniil@gmail.com', '65757634353637356b64666a643435386468673433a94a8fe5ccb19ba61c4c0873d391e987982fbbd3', 'customer')
// 	`)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
