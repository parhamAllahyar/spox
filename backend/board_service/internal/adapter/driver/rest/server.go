package rest

import (
	"board/configs"
	"board/internal/core/usecase"
	"net/http"
)

type server struct {
	BoardUsecase   usecase.BoardUsecase
	SectionUsecase usecase.SectionUsecase
	SubTaskUsecase usecase.SubTaskUsecase
	TagUsecase     usecase.TagUsecase
	TaskUsecase    usecase.TaskUsecase
}

func NewServer(
	BoardUsecase usecase.BoardUsecase,
	SectionUsecase usecase.SectionUsecase,
	SubTaskUsecase usecase.SubTaskUsecase,
	TagUsecase usecase.TagUsecase,
	TaskUsecase usecase.TaskUsecase,
) server {
	return server{
		BoardUsecase:   BoardUsecase,
		SectionUsecase: SectionUsecase,
		SubTaskUsecase: SubTaskUsecase,
		TagUsecase:     TagUsecase,
		TaskUsecase:    TaskUsecase,
	}
}

func (srv server) InitServer(config configs.HttpServer) error {

	addr := config.Port

	//TODO: load from config
	return http.ListenAndServe(":82", srv.Router())

}
