package rest

import (
	"github.com/gorilla/mux"
)

func (srv server) Router() *mux.Router {

	router := mux.NewRouter()

	workspace := router.PathPrefix("/workspace").Subrouter()
	workspace.HandleFunc("/create", srv.createWorkspace).Methods("POST")
	workspace.HandleFunc("/", srv.userWorkspacesList).Methods("GET")
	workspace.HandleFunc("/teammate/{id}", srv.showWorkspaceByID).Methods("GET")
	workspace.HandleFunc("/delete/{id}", srv.deleteWorkspace).Methods("POST")
	workspace.HandleFunc("/update/{id}", srv.updateWorkspace).Methods("POST")

	role := router.PathPrefix("/role").Subrouter()
	role.HandleFunc("/", srv.createRole).Methods("GET")
	role.HandleFunc("/{id}", srv.showRoleByID).Methods("GET")
	role.HandleFunc("/update/{id}", srv.updateRole).Methods("POST")
	role.HandleFunc("/update/{id}", srv.assignRoleToUser).Methods("POST")

	return router

}
