package controller

import (
	"github.com/jmoiron/sqlx"
	"main.go/internal/controller"
	"main.go/internal/repository"
	"main.go/internal/service"
	_ "modernc.org/sqlite"
)

func InitEnv() (*controller.Handler, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(db)
	services := service.NewServices(repo)
	handler := controller.NewHandler(services)

	return handler, nil
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

	return db, nil
}
