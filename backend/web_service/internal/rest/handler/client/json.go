package client

type SignUpRequest struct {
	Phone string `json:phone`
}

type SignInRequest struct {
	Phone    string `json:phone`
	Password string `json:password`
}

type ResponseMessage struct {
	Message string `json:message`
}
