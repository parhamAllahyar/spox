package resource

type SignInRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type SignInResponse struct {
	Token string `json:"token"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type SignUpRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type SignUpResponse struct {
	Token string `json:"token"`
}

type ForgetPasswordCodeRequest struct{}
type ForgetPasswordCodeResponse struct{}

type ChangePasswordRequest struct{}
type ChangePasswordResponse struct{}
