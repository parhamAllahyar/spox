package resource

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SignInResponse struct {
	Token string `json:"token"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SignUpResponse struct {
	Token string `json:"token"`
}

type ForgetPasswordCodeRequest struct{}
type ForgetPasswordCodeResponse struct{}

type ChangePasswordRequest struct{}
type ChangePasswordResponse struct{}
