package main

import (
	"Resrik/internal/server"
	"Resrik/pkg/handler"
	"Resrik/pkg/repository"
	"Resrik/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initing cfg: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to load env: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error with database %s", err.Error())
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handlers := handler.NewHandler(serv)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouts()); err != nil {
		log.Fatalf("error while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
