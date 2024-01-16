package main

import (
	"web/configs"

	"web/internal/rest"
	"web/internal/rest/handler/admin"
	"web/internal/rest/handler/client"
)

func main() {

	config := configs.LoadConfig()

	services := configs.ServicesAddres(config.Services)

	// Admin Handler
	adminHandler := admin.NewAdminHandler(services)
	clientHandler := client.NewClientHandler(services)

	httpserver := rest.NewServer(adminHandler, clientHandler)
	httpserver.InitServer(config.HTTPServer)

}
