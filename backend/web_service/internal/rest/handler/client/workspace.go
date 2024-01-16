package client

import (
	"fmt"
	"net/http"
	"web/internal/rest/handler"

	"github.com/gorilla/mux"
)

func CreateWorkspace(w http.ResponseWriter, r *http.Request) {

	fmt.Println("from bff")

	handler.ForwardURL(w, r, "workspace/create", "workspace")
}

func UpdateWorkspace(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("err")
	}

	fmt.Println("workspace/update" + id)

	handler.ForwardURL(w, r, "workspace/update"+id, "workspace")

}

func DeleteWorkspace(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("err")
	}

	handler.ForwardURL(w, r, "workspace/delete"+id, "workspace")
}

func WorkspaceList(w http.ResponseWriter, r *http.Request) {

	handler.ForwardURL(w, r, "workspace", "workspace")

}

func WorkspaceByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("err")
	}

	handler.ForwardURL(w, r, "workspace/update/"+id, "workspace")
}
