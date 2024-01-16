package rest

import (
	"net/http"
	"user/configs"
	"user/internal/core/usecase"
)

type server struct {
	AuthUsecase usecase.AuthUsecase
	UserUsecase usecase.UserUsecase
}

func NewServer(
	AuthUsecase usecase.AuthUsecase,
	UserUsecase usecase.UserUsecase,
) server {
	return server{
		AuthUsecase: AuthUsecase,
		UserUsecase: UserUsecase,
	}
}

func (srv server) InitServer(config configs.HttpServer) error {

	// headersOk := handlers.AllowedHeaders([]string{"Content-Type, Authorization"})
	// originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000/login"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// //TODO: load from config
	// fmt.Println("BFF server is Start")
	// // log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))

	// return http.ListenAndServe(":8085", handlers.CORS(originsOk, headersOk, methodsOk)(srv.Router()))

	return http.ListenAndServe(":8085", srv.Router())

}
