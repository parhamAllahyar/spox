package client

import (
	config "web/configs"
)

type ClientHandler struct {
	Services config.ServiceAddr
}

func NewClientHandler(Services config.ServiceAddr) ClientHandler {
	return ClientHandler{
		Services: Services,
	}
}
