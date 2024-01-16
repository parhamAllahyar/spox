package http

import (
	"asset/config"
	"asset/internal/core/usecase"
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
)

type application struct {
	FileUsecase usecase.FileUsecase
}

func NewApplication(FileUsecase usecase.FileUsecase) application {

	return application{
		FileUsecase: FileUsecase,
	}

}

func (app application) InitServer(config config.HttpServer) error {

	//TODO: load from config

	addr := config.Port

	fmt.Println(addr)

	return http.ListenAndServe(":82", app.Router())

}
