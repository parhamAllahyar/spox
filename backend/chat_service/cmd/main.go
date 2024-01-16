package main

import (
	"chat/configs"
	"chat/internal/adapter/driven/cassandra"
	"chat/internal/adapter/driver/http"
	"chat/internal/core/usecase"
	"fmt"
)

func main() {

	config := configs.LoadConfig()

	//DAO
	dao, err := cassandra.CassandraConnection(config.CassandaraDBConfig)

	// conversationDAO := cassandra.
	messageDAO := cassandra.NewMessageDAO(dao)

	conversationDAO := cassandra.NewConversationDAO(dao)

	conversationUsecase := usecase.NewConversationUsecase(conversationDAO)
	messageUsecase := usecase.NewMessageUsecase(messageDAO)

	server := http.NewServer(conversationUsecase, messageUsecase)
	err = server.InitServer(config.HttpServer)
	if err != nil {

		fmt.Println(err)

	}

}
