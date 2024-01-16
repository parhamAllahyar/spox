package rest

import (
	// "workspace/internal/core/dto"
	"encoding/json"
	"net/http"
	"workspace/internal/core/dto"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	// "github.com/google/uuid"
)

// TODO:List of all boards
func (srv server) roles(w http.ResponseWriter, r *http.Request) {

	// // TODO:
	// dto := dto.BoardsRequest{
	// 	AccessToken: "",
	// 	WorkspaceID: uuid.New(),
	// }

	// srv.BoardUsecase.Boards(dto)

}

// TODO:Show Board with all data such as task section tag
func (srv server) showRoleByID(w http.ResponseWriter, r *http.Request) {

	permission, err := srv.RoleUsecase.RoleList()
	if err != nil {
		// TODO:LOG
	}

	err = json.NewEncoder(w).Encode(&permission)
	if err != nil {
		//
	}

}

func (srv server) deleteRole(w http.ResponseWriter, r *http.Request) {

	dto := dto.DeleteRoleRequest{
		AccessToken: r.Header.Get("x-access-token"),
		ID:          uuid.MustParse(mux.Vars(r)["id"]),
	}

	err := srv.RoleUsecase.Delete(dto)

	if err != nil {
		// TODO:LOG
	}

}

func (srv server) updateRole(w http.ResponseWriter, r *http.Request) {

	var dto dto.UpdateRoleRequest
	err := json.NewDecoder(r.Body).Decode(&dto)
	dto.AccessToken = r.Header.Get("x-access-token")
	dto.ID = uuid.MustParse(mux.Vars(r)["id"])

	if err != nil {
		// TODO:LOG
	}

	err = srv.RoleUsecase.Update(dto)

	if err != nil {
		// TODO:LOG
	}

}

func (srv server) createRole(w http.ResponseWriter, r *http.Request) {

	var dto dto.CreateRoleRequest
	err := json.NewDecoder(r.Body).Decode(&dto)
	dto.AccessToken = r.Header.Get("x-access-token")

	if err != nil {
		// TODO:LOG
	}

	err = srv.RoleUsecase.Create(dto)

	if err != nil {
		// TODO:LOG
	}

}

func (srv server) assignRoleToUser(w http.ResponseWriter, r *http.Request) {

	var dto dto.CreateRoleRequest
	err := json.NewDecoder(r.Body).Decode(&dto)
	dto.AccessToken = r.Header.Get("x-access-token")

	if err != nil {
		// TODO:LOG
	}

	err = srv.RoleUsecase.Create(dto)

	if err != nil {
		// TODO:LOG
	}

}
