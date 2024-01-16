package rest

import (
	"board/internal/adapter/driver/rest/resource"
	"board/internal/core/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO:Show all sub task by task ID
func (srv server) subTasks(w http.ResponseWriter, r *http.Request) {
	dto := dto.SubTasksRequest{}
	srv.SubTaskUsecase.SubTasks(dto)
}

func (srv server) showSubTaskByID(w http.ResponseWriter, r *http.Request) {

}

// TODO:delete sub task by id
func (srv server) deleteSubTask(w http.ResponseWriter, r *http.Request) {

	req := resource.DeleteSubTaskRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.SubTaskUsecase.Delete(dto.DeleteSubTaskRequest{})
	fmt.Println(err)
}

// TODO:update sub task by id
func (srv server) updateSubTask(w http.ResponseWriter, r *http.Request) {

	req := resource.DeleteSubTaskRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.SubTaskUsecase.Update(dto.UpdateSubTaskRequest{})
	fmt.Println(err)

}
