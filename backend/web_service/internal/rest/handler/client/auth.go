package client

import (
	"net/http"
	"web/internal/rest/handler"
)

func (c ClientHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	handler.ForwardURL(w, r, "auth/signin", "user")
}

func (c ClientHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	handler.ForwardURL(w, r, "auth/signup", "user")
}
