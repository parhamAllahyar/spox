package rest

import (
	"net/http"
	"workspace/configs"
	"workspace/internal/core/usecase"
)

type server struct {
	RoleUsecase      usecase.RoleUsecase
	WorkspaceUsecase usecase.WorkspaceUsecase
}

func NewServer(
	RoleUsecase usecase.RoleUsecase,
	WorkspaceUsecase usecase.WorkspaceUsecase) server {
	return server{
		RoleUsecase:      RoleUsecase,
		WorkspaceUsecase: WorkspaceUsecase,
	}
}

func (srv server) InitServer(config configs.HttpServer) error {

	//TODO: load from config
	return http.ListenAndServe(":82", srv.Router())

}
