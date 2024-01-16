package http

import (
	"chat/internal/core/dto"

	"net/http"

	"github.com/olahol/melody"
)

func (srv server) GetUserInConversationHandler(w http.ResponseWriter, r *http.Request) {
}

func (srv server) GetMeInConversationHandler(w http.ResponseWriter, r *http.Request) {

}

func (srv server) UserConversationListHandler(w http.ResponseWriter, r *http.Request) {

}

func (srv *server) websocketHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.m.HandleRequest(w, r)
	})
}

func (srv server) StartConversationHandler(s *melody.Session) {

	errorMessage := "Failed to establish WebSocket connection"
	s.Write([]byte(errorMessage))
	s.Close()

	dto := dto.StartConversation{
		Channel:     s.Request.URL.Query().Get("channel"),
		Key:         s.Request.URL.Query().Get("key"),
		AccessToken: s.Request.Header.Get("Authorization"),
	}

	err := srv.conversationUsecase.StartConversation(dto)

	if err != nil {
		errorMessage := "Failed to establish WebSocket connection"
		s.Write([]byte(errorMessage))
		s.Close()
	}

}
