package main

import (
	mapsRequest "github.com/SashaVolohov/mapsRequestServer"
	"github.com/SashaVolohov/mapsRequestServer/internal/handler"
	"github.com/SashaVolohov/mapsRequestServer/internal/repository"
	"github.com/SashaVolohov/mapsRequestServer/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Failed to read configuration: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	go services.API.KeyCollector()

	server := new(mapsRequest.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Failed to initialize server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
