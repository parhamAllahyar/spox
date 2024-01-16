package rest

import (
	"web/internal/rest/handler/admin"
	"web/internal/rest/handler/client"

	// "github.com/rs/cors"
	"github.com/gorilla/mux"
)

func (srv Server) Router() *mux.Router {

	router := mux.NewRouter()

	user := router.PathPrefix("").Subrouter()
	//----------------------------- USER AUTH -----------------------------
	user.HandleFunc("/signin", srv.ClientHandler.SignIn).Methods("POST", "OPTIONS")
	user.HandleFunc("/signup", srv.ClientHandler.SignUp).Methods("POST", "OPTIONS")
	user.HandleFunc("/user/update", srv.ClientHandler.UserUpdate).Methods("POST")

	//----------------------------- USER WORKSPACE -------------------------
	userWorkspace := router.PathPrefix("/workspace").Subrouter()
	userWorkspace.HandleFunc("/create", client.CreateWorkspace).Methods("POST")
	userWorkspace.HandleFunc("/delete/{id}", client.DeleteWorkspace).Methods("POST")
	userWorkspace.HandleFunc("/update/{id}", client.UpdateWorkspace).Methods("POST")
	userWorkspace.HandleFunc("/", client.WorkspaceList).Methods("GET")
	userWorkspace.HandleFunc("/{id}", client.WorkspaceByID).Methods("GET")

	//----------------------------- USER BOARD -----------------------------
	userBoard := router.PathPrefix("/board").Subrouter()
	userBoard.HandleFunc("/{id}", client.BoardByID).Methods("GET")
	userBoard.HandleFunc("/", client.BoardList).Methods("GET")
	userBoard.HandleFunc("/create", client.CreateBoard).Methods("POST")
	userBoard.HandleFunc("/update/{id}", client.UpdateBoard).Methods("POST")
	userBoard.HandleFunc("/delete/{id}", client.DeleteBoard).Methods("POST")

	//----------------------------- USER SECTION -----------------------------
	userSection := router.PathPrefix("/section").Subrouter()
	userSection.HandleFunc("/delete/{id}", client.DeleteSection).Methods("POST")
	userSection.HandleFunc("/update/{id}", client.UpdateSection).Methods("POST")
	userSection.HandleFunc("/create", client.CreateSection).Methods("POST")

	//----------------------------- USER TASK --------------------------------
	userTask := router.PathPrefix("/task").Subrouter()
	userTask.HandleFunc("/{id}", client.TaskByID).Methods("GET")
	userTask.HandleFunc("/", client.TaskList).Methods("GET")
	userTask.HandleFunc("/create", client.CreateTask).Methods("POST")
	userTask.HandleFunc("/update/{id}", client.UpdateTask).Methods("POST")
	userTask.HandleFunc("/delete/{id}", client.DeleteTask).Methods("POST")

	//----------------------------- ADMIN AUTH -----------------------------
	administrator := router.PathPrefix("/admin").Subrouter()
	administrator.HandleFunc("/signin", srv.AdminHandler.SignIn).Methods("POST")

	//----------------------------- ADMIN STATISTICS -----------------------------
	administratorStatistics := router.PathPrefix("/admin/statistics").Subrouter()
	administratorStatistics.HandleFunc("/overview", admin.Overview).Methods("POST")
	administratorStatistics.HandleFunc("/payments", admin.Payment).Methods("POST")
	administratorStatistics.HandleFunc("/users", admin.User).Methods("POST")
	administratorStatistics.HandleFunc("/boards", admin.Board).Methods("POST")
	administratorStatistics.HandleFunc("/storage", admin.Storage).Methods("POST")
	administratorStatistics.HandleFunc("/communication", admin.Communication).Methods("POST")

	//----------------------------- ADMIN USERS -----------------------------
	administratorUser := router.PathPrefix("/admin/user").Subrouter()
	administratorUser.HandleFunc("/", admin.User).Methods("GET")
	administratorUser.HandleFunc("/{id}", admin.UserByID).Methods("GET")

	//----------------------------- ADMIN WORKSPACES -------------------------
	administratorWorkspace := router.PathPrefix("/workspace").Subrouter()
	administratorWorkspace.HandleFunc("/", admin.Workspace).Methods("GET")
	administratorWorkspace.HandleFunc("/{id}", admin.WorkspaceByID).Methods("GET")

	//----------------------------- ADMIN Payments -------------------------==
	administratorPayment := router.PathPrefix("/payment").Subrouter()
	administratorPayment.HandleFunc("/", admin.Payment).Methods("POST")

	return router

}
