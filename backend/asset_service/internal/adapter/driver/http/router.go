package http

import (
	"github.com/gorilla/mux"

)

func (app application) Router() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/create", app.CreateOrder).Methods("POST")
	return router
}
