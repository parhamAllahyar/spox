package main

import (
	"asset/config"
	"asset/internal/adapter/driven/postgres"
	"asset/internal/adapter/driven/s3"
	"asset/internal/adapter/driver/http"
	"asset/internal/core/usecase"
	"fmt"
)

func main() {

	fmt.Println("ANNNNNNNNNNNNNNNNNNNNNNNNN")

	config := config.LoadConfig()

	//S3
	S3 := s3.NewS3(config.S3)

	//Postgres
	postgresInstanse, err := postgres.Connection(config.Postgre)
	if err != nil {
		panic(err)
	}

	//DAO
	fileDAO := postgres.NewFileDAO(postgresInstanse)

	// Usecase
	fileUsecase := usecase.NewFileUsecase(fileDAO, S3)

	//Http
	application := http.NewApplication(fileUsecase)

	err = application.InitServer(config.HttpServer)

	if err != nil {
		panic(err)
	}

}
