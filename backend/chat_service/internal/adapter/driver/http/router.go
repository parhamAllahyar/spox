package http

import "github.com/gorilla/mux"

func (srv server) Router() *mux.Router {

	router := mux.NewRouter()

	// Chat
	message := router.PathPrefix("/message").Subrouter()

	// Past Message
	message.HandleFunc("/messages/{conversationID}", srv.GetPastMessagesHandler).Methods("GET")

	// Web Socket
	router.Handle("/ws", srv.websocketHandler())
	srv.m.HandleConnect(srv.StartConversationHandler)
	srv.m.HandleMessage(srv.SendTextMessageHandler)

	// Conversation
	conversation := router.PathPrefix("/conversation").Subrouter()

	// Join to Conversation
	conversation.HandleFunc("/join/{conversationID}", srv.GetPastMessagesHandler).Methods("GET")

	conversation.HandleFunc("/list", srv.UserConversationListHandler).Methods("GET")

	// Conversation Me
	conversation.HandleFunc("/user/me", srv.GetMeInConversationHandler).Methods("GET")

	// Conversation Participant
	conversation.HandleFunc("/user/{conversationID}", srv.GetUserInConversationHandler).Methods("GET")

	return router

}
