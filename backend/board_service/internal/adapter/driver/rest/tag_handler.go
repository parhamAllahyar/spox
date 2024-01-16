package rest

import (
	"board/internal/adapter/driver/rest/resource"
	"board/internal/core/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO:Show all tags of a board
func (srv server) tags(w http.ResponseWriter, r *http.Request) {
}

func (srv server) showTagByID(w http.ResponseWriter, r *http.Request) {

}

// TODO: delete tag by id
func (srv server) deleteTag(w http.ResponseWriter, r *http.Request) {

	req := resource.DeleteTagRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.TagUsecase.Delete(dto.DeleteTagRequest{})
	fmt.Println(err)

}

// TODO: update tag by id
func (srv server) updateTag(w http.ResponseWriter, r *http.Request) {
	req := resource.DeleteTagRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.TagUsecase.Update(dto.UpdateTagRequest{})
	fmt.Println(err)
}
