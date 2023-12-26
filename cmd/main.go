package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"main.go/internal/controller"
	"main.go/internal/model"
	"main.go/internal/repository"
	"main.go/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("an error occurred while initializing the config: %s", err.Error())
	}

	db, err := repository.NewSQLiteDB()
	if err != nil {
		logrus.Fatalf("an error occurred while opening db: %s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewServices(repo)
	handler := controller.NewHandler(services)

	srv := new(model.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("an error occurred while running the server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
