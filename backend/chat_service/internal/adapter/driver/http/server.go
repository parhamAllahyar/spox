package http

import (
	"chat/configs"
	"chat/internal/core/usecase"
	"net/http"

	"github.com/olahol/melody"
)

type server struct {
	m                   *melody.Melody
	conversationUsecase usecase.ConversationUsecase
	messageUsecase      usecase.MessageUsecase
}

func NewServer(
	conversationUsecase usecase.ConversationUsecase,
	messageUsecase usecase.MessageUsecase,
) server {
	return server{
		conversationUsecase: conversationUsecase,
		messageUsecase:      messageUsecase,
		m:                   melody.New(),
	}
}

func (srv server) InitServer(config configs.HttpServer) error {

	addr := ":9091"

	//TODO: load from config
	return http.ListenAndServe(addr, srv.Router())

}
