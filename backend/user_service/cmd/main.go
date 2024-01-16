package main

import (
	"fmt"
	"user/configs"
	"user/internal/adapter/driven/postgres"
	"user/internal/adapter/driver/rest"
	"user/internal/core/usecase"
	"user/utils/dev"
)

func main() {

	// CONFIG
	config := configs.LoadConfig()

	if config.ENV.ENV == "dev" {
		dev.Migration(config.Postgre)
	}

	// DAO
	dao, err := postgres.Connection(config.Postgre)

	if err != nil {
		fmt.Println(err)
	}

	userDao := postgres.NewUserRepo(dao)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(userDao)
	userUsecase := usecase.NewUserUsecase(userDao)

	// START SERVER
	server := rest.NewServer(authUsecase, userUsecase)
	server.InitServer(config.HttpServer)

}
