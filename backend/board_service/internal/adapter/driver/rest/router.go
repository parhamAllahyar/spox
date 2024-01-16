package rest

import (
	"github.com/gorilla/mux"
)

func (srv server) Router() *mux.Router {

	router := mux.NewRouter()

	// BOARD
	board := router.PathPrefix("/board").Subrouter()
	//all board of a workspace
	board.HandleFunc("/", srv.boards).Methods("GET")
	board.HandleFunc("/{id}", srv.showBoardByID).Methods("GET")
	board.HandleFunc("/delete/{id}", srv.deleteBoard).Methods("POST")
	board.HandleFunc("/update/{id}", srv.updateBoard).Methods("POST")

	// SECTION
	section := router.PathPrefix("/section").Subrouter()
	//all board of a board
	section.HandleFunc("/", srv.sections).Methods("GET")
	section.HandleFunc("/{id}", srv.showSectionByID).Methods("GET")
	section.HandleFunc("/delete/{id}", srv.deleteSection).Methods("POST")
	section.HandleFunc("/update/{id}", srv.updateSection).Methods("POST")

	// TASK
	task := router.PathPrefix("/task").Subrouter()
	task.HandleFunc("/", srv.tasks).Methods("GET")
	task.HandleFunc("/{id}", srv.showTaskByID).Methods("GET")
	task.HandleFunc("/delete/{id}", srv.deleteTask).Methods("POST")
	task.HandleFunc("/update/{id}", srv.updateTask).Methods("POST")

	// SUBTASK
	subTask := router.PathPrefix("/sub-task").Subrouter()
	subTask.HandleFunc("/", srv.subTasks).Methods("GET")
	subTask.HandleFunc("/{id}", srv.showSubTaskByID).Methods("GET")
	subTask.HandleFunc("/delete/{id}", srv.deleteSubTask).Methods("POST")
	subTask.HandleFunc("/update/{id}", srv.updateSubTask).Methods("POST")

	// TAG
	tag := router.PathPrefix("/tag").Subrouter()
	tag.HandleFunc("/", srv.tags).Methods("GET")
	tag.HandleFunc("/{id}", srv.showTagByID).Methods("GET")
	tag.HandleFunc("/delete/{id}", srv.deleteTag).Methods("POST")
	tag.HandleFunc("/update/{id}", srv.updateTag).Methods("POST")

	return router

}
