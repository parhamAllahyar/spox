package client

import (
	"net/http"
	"web/internal/rest/handler"
)

func (c ClientHandler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	handler.ForwardURL(w, r, "user/update", "user")
}
