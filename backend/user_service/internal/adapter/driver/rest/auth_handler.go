package rest

import (
	"encoding/json"
	"net/http"
	"user/internal/adapter/driver/rest/resource"
	"user/internal/core/dto"
)

func (srv server) signup(w http.ResponseWriter, r *http.Request) {

	var req resource.SignUpRequest
	json.NewDecoder(r.Body).Decode(&req)

	token, err := srv.AuthUsecase.SignUp(dto.SignUpRequest{
		Password: req.Password,
		Email:    req.Email,
	})

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

	response := resource.SignUpResponse{
		Token: token,
	}

	jsonData, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	return

}

func (srv server) signin(w http.ResponseWriter, r *http.Request) {

	var req resource.SignInRequest
	json.NewDecoder(r.Body).Decode(&req)

	token, err := srv.AuthUsecase.SignIn(dto.SignInRequest{
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		response := resource.MessageResponse{
			Message: err.Error(),
		}

		jsonData, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
		return

	}

	response := resource.SignInResponse{
		Token: token,
	}

	jsonData, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonData)
	return

}

func (srv server) forgetPasswordCode(w http.ResponseWriter, r *http.Request) {

}

func (srv server) changePassword(w http.ResponseWriter, r *http.Request) {

	responseData := map[string]interface{}{
		"key1": "value1",
		"key2": "nbnbnbnbnbn",
		"key3": true,
	}

	// Convert the map to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type as JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

	// req := resource.ChangePasswordRequest{}
	// json.NewDecoder(r.Body).Decode(&req)
	// err := srv.AuthUsecase.UpdatePassword(dto.UpdatePasswordRequest{})
	// fmt.Println(err)
}
