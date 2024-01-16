package main

import (
	"workspace/configs"
	"workspace/internal/adapter/driven/postgres"
	"workspace/internal/adapter/driver/rest"
	"workspace/internal/core/usecase"
	"workspace/utils/dev"
)

func main() {

	// CONFIG
	config := configs.LoadConfig()

	// if config.ENV.ENV == "dev" {
	dev.Migration(config.Postgre)
	// }

	// Database Connection
	dao, err := postgres.Connection(config.Postgre)
	if err != nil {
		panic(err)
	}

	// DAO
	roleDao := postgres.NewRoleRepo(dao)
	workspaceDao := postgres.NewWorkspaceRepo(dao)

	// USECASE
	roleUsecase := usecase.NewRoleUsecase(roleDao)
	workspaceUsecase := usecase.NewWorkspaceUsecase(workspaceDao, roleUsecase, roleDao)

	// START SERVER
	server := rest.NewServer(roleUsecase, workspaceUsecase)
	server.InitServer(config.HttpServer)

}
