package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"board/internal/adapter/driver/rest/resource"
	"board/internal/core/dto"
)

func (srv server) sections(w http.ResponseWriter, r *http.Request) {

}

func (srv server) showSectionByID(w http.ResponseWriter, r *http.Request) {

	// req := resource.ShowSectionByIDRequest{}
	// json.NewDecoder(r.Body).Decode(&req)
	// token, err := srv.SectionUsecase.
	// fmt.Println(token, err)

}

// TODO: Delete a section By ID
func (srv server) deleteSection(w http.ResponseWriter, r *http.Request) {

	req := resource.DeleteSectionRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.SectionUsecase.Delete(dto.DeleteSectionRequest{})
	fmt.Println(err)

}

// TODO: Update a section By ID
func (srv server) updateSection(w http.ResponseWriter, r *http.Request) {

	req := resource.ShowSectionByIDRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.SectionUsecase.Update(dto.UpdateSectionRequest{})
	fmt.Println(err)

}
