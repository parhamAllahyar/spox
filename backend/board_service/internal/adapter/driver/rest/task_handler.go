package rest

import (
	"board/internal/adapter/driver/rest/resource"
	"board/internal/core/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

func (srv server) tasks(w http.ResponseWriter, r *http.Request) {

	// srv.TaskUsecase.

}

func (srv server) showTaskByID(w http.ResponseWriter, r *http.Request) {

}

// TODO: delete a task by id
func (srv server) deleteTask(w http.ResponseWriter, r *http.Request) {

	req := resource.DeleteSectionRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.TaskUsecase.Update(dto.UpdateTaskRequest{})
	fmt.Println(err)

}

// TODO: Update a task by id
func (srv server) updateTask(w http.ResponseWriter, r *http.Request) {
	req := resource.DeleteTaskRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.TaskUsecase.Update(dto.UpdateTaskRequest{})
	fmt.Println(err)
}
