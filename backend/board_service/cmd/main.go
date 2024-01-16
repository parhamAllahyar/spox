package main

import (
	"board/configs"
	"board/internal/adapter/driven/postgres"
	restDriven "board/internal/adapter/driven/rest"
	"board/internal/adapter/driver/rest"
	"board/internal/core/usecase"
	"board/utils/dev"
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
		// TODO:
	}

	// DAO
	boardDao := postgres.NewBoardRepo(dao)
	sectionDao := postgres.NewSectionRepo(dao)
	taskDao := postgres.NewTaskRepo(dao)
	subTaskDao := postgres.NewSubTaskdRepo(dao)
	tagDao := postgres.NewTagRepo(dao)

	workspaceDao := restDriven.NewWorkspaceAPI()

	// USECASE
	boardUsecase := usecase.NewBoardUsecase(boardDao, workspaceDao)
	sectionUsecase := usecase.NewSectionUsecase(sectionDao, workspaceDao)
	taskUsecase := usecase.NewTaskUsecase(taskDao, workspaceDao)
	subTasUsecase := usecase.NewSubTaskUsecase(subTaskDao, workspaceDao)
	tagUsecase := usecase.NewTagUsecase(tagDao, workspaceDao)

	// START SERVER
	server := rest.NewServer(boardUsecase, sectionUsecase, subTasUsecase, tagUsecase, taskUsecase)
	server.InitServer(config.HttpServer)

}
