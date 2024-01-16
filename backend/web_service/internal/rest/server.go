package rest

import (
	"fmt"
	"net/http"
	"web/configs"
	"web/internal/rest/handler/admin"
	"web/internal/rest/handler/client"

	"github.com/rs/cors"
)

type Server struct {
	AdminHandler  admin.AdminHandler
	ClientHandler client.ClientHandler
}

func NewServer(AdminHandler admin.AdminHandler, ClientHandler client.ClientHandler) *Server {
	return &Server{
		AdminHandler:  AdminHandler,
		ClientHandler: ClientHandler,
	}
}

func (srv Server) InitServer(config configs.HttpServer) error {

	fmt.Println(config.Port)

	port := fmt.Sprintf("%s%s", ":", config.Port)

	// headersOk := handlers.AllowedHeaders([]string{"Content-Type, Authorization"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// return http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(srv.Router()))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000","http://localhost:3001"},
		AllowCredentials: true,
	})

	handler := c.Handler(srv.Router())

	return http.ListenAndServe(port, handler)
}
