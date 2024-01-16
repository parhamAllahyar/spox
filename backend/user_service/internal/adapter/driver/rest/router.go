package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (srv server) Router() *mux.Router {

	router := mux.NewRouter()

	// h := http.HandlerFunc(NotFound)
	// router.NotFoundHandler = h

	router.HandleFunc("/users", srv.signin).Methods("GET")

	// AUTH
	auth := router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signin", srv.signin).Methods("POST", "OPTIONS")
	auth.HandleFunc("/signup", srv.signup).Methods("POST", "OPTIONS")
	auth.HandleFunc("/delete/{id}", srv.forgetPasswordCode).Methods("POST")

	// USER
	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/list", srv.users).Methods("GET")
	user.HandleFunc("/search", srv.search).Methods("POST")
	user.HandleFunc("/count", srv.userCount).Methods("GET")
	user.HandleFunc("/update", srv.userUpdate).Methods("POST")

	// corsMiddleware := func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		// Set headers to allow all origins
	// 		w.Header().Set("Access-Control-Allow-Origin", "*")
	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// 		if r.Method == "OPTIONS" {
	// 			// Handle preflight requests
	// 			w.WriteHeader(http.StatusOK)
	// 			return
	// 		}

	// 		// Continue processing request
	// 		next.ServeHTTP(w, r)
	// 	})
	// }

	// // Use CORS middleware with the router
	// router.Use(corsMiddleware)

	// user.Use(corsMiddleware)

	return router

}

func NotFound(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, this is a simple HTTP server user service!")

}
