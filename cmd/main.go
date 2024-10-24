package main

import (
	buildingBase "github.com/kahuri1/building_base"
	"github.com/kahuri1/building_base/model"
	"github.com/kahuri1/building_base/pkg/handler"
	"github.com/kahuri1/building_base/pkg/repository"
	"github.com/kahuri1/building_base/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title building_base API
// @version 1.0
// @description API server building_base

// @host localhost:8000
// @basePath /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	db, err := repository.NewDB(model.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Errorf("Failed initialization db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.Newhandler(service)
	srv := new(buildingBase.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
