package rest

import (
	"board/internal/adapter/driver/rest/resource"
	"board/internal/core/dto"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// TODO:List of all boards
func (srv server) boards(w http.ResponseWriter, r *http.Request) {

	// TODO:
	dto := dto.BoardsRequest{
		AccessToken: "",
		WorkspaceID: uuid.New(),
	}

	srv.BoardUsecase.Boards(dto)

}

// TODO:Show Board with all data such as task section tag
func (srv server) showBoardByID(w http.ResponseWriter, r *http.Request) {

	dto := dto.BoardRequest{
		AccessToken: "",
		ID:          uuid.New(),
	}
	srv.BoardUsecase.Board(dto)

}

func (srv server) deleteBoard(w http.ResponseWriter, r *http.Request) {

	dto := dto.DeleteBoardRequest{
		AccessToken: "",
		ID:          uuid.New(),
	}
	srv.BoardUsecase.Delete(dto)

}

func (srv server) updateBoard(w http.ResponseWriter, r *http.Request) {

	req := resource.DeleteBoardRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.BoardUsecase.Delete(dto.DeleteBoardRequest{})
	fmt.Println(err)

}

func (srv server) createBoard(w http.ResponseWriter, r *http.Request) {

	req := resource.DeleteBoardRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.BoardUsecase.Update(dto.UpdateBoardRequest{})
	fmt.Println(err)

}
