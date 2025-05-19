package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

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

	ctx, keyCollectorCancel := context.WithCancel(context.Background())
	go services.API.KeyCollector(ctx)

	server := new(mapsRequest.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Failed to initialize server: %s", err.Error())
		}
	}()

	logrus.Println("MapsRequestServer has started!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Println("MapsRequestServer has shutting down...")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Unable to shutdown server: %s", err.Error())
	}

	keyCollectorCancel()

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
