package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user/internal/adapter/driver/rest/resource"
	"user/internal/core/dto"
	"user/utils/auth/authenticator"
)

func (srv server) userCount(w http.ResponseWriter, r *http.Request) {

}

func (srv server) search(w http.ResponseWriter, r *http.Request) {
	req := resource.ChangePasswordRequest{}
	json.NewDecoder(r.Body).Decode(&req)
	err := srv.AuthUsecase.UpdatePassword(dto.UpdatePasswordRequest{})
	fmt.Println(err)
}

func (srv server) searchByEmail(w http.ResponseWriter, r *http.Request) {

}

func (srv server) users(w http.ResponseWriter, r *http.Request) {

}

func (srv server) userUpdate(w http.ResponseWriter, r *http.Request) {

	req := resource.UpdateUserInfoRequest{}
	json.NewDecoder(r.Body).Decode(&req)

	Token, err := authenticator.GetToken(r)

	if err != nil {
		fmt.Println(err)
	}

	err = srv.UserUsecase.Update(dto.UpdateUserRequest{
		AccessToken: Token,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
	})

	fmt.Println(err)

}
