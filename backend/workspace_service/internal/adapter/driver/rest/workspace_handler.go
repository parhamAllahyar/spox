package rest

import (
	"encoding/json"
	"net/http"
	"workspace/internal/adapter/driver/rest/resource"
	"workspace/internal/core/dto"
	"workspace/utils/auth/authenticator"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// TODO:List of all boards
func (srv server) userWorkspacesList(w http.ResponseWriter, r *http.Request) {

	// workspaces, err := srv.WorkspaceUsecase.UserWorkspaces(r.Header.Get("x-access-token"))
	// if err != nil {
	// 	// TODO:LOG
	// }

	// err = json.NewEncoder(w).Encode(&workspaces)
	// if err != nil {
	// 	//
	// }

	// if err != nil {
	// 	// TODO:LOG
	// }

	// TODO:
	// dto := dto.Wor{
	// 	AccessToken: r.Header.Get("x-access-token"),
	// 	WorkspaceID: uuid.New(),
	// }

	// srv.BoardUsecase.Boards(dto)

}

// TODO:Show Board with all data such as task section tag
func (srv server) showWorkspaceByID(w http.ResponseWriter, r *http.Request) {

	dto := dto.ShowWorkspaceByIDRequest{
		AccessToken: r.Header.Get("x-access-token"),
		ID:          uuid.MustParse(mux.Vars(r)["id"]),
	}

	workspace, err := srv.WorkspaceUsecase.WorkspacesById(dto)
	if err != nil {
		// TODO:LOG
	}

	err = json.NewEncoder(w).Encode(&workspace)
	if err != nil {
		//
	}

}

func (srv server) deleteWorkspace(w http.ResponseWriter, r *http.Request) {

	dto := dto.DeleteWorkspaceByIDRequest{
		AccessToken: r.Header.Get("x-access-token"),
		ID:          uuid.MustParse(mux.Vars(r)["id"]),
	}

	err := srv.WorkspaceUsecase.Delete(dto)

	if err != nil {
		// TODO:LOG
	}

}

func (srv server) updateWorkspace(w http.ResponseWriter, r *http.Request) {

	var dto dto.UpdateWorkspaceRequest
	err := json.NewDecoder(r.Body).Decode(&dto)
	dto.AccessToken = r.Header.Get("x-access-token")
	dto.ID = uuid.MustParse(mux.Vars(r)["id"])

	if err != nil {
		// TODO:LOG
	}

	err = srv.WorkspaceUsecase.Update(dto)

	if err != nil {
		// TODO:LOG
	}

}

func (srv server) createWorkspace(w http.ResponseWriter, r *http.Request) {

	var dto dto.CreateWorkspaceRequest
	err := json.NewDecoder(r.Body).Decode(&dto)

	dto.AccessToken, err = authenticator.GetToken(r)

	if err != nil {
		response := resource.MessageResponse{
			Message: err.Error(),
		}

		jsonData, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
		return
	}

	err = srv.WorkspaceUsecase.Create(dto)

	if err != nil {
		response := resource.MessageResponse{
			Message: err.Error(),
		}

		jsonData, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
		return
	}

	response := resource.MessageResponse{
		Message: "created",
	}
	jsonData, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	return

}
