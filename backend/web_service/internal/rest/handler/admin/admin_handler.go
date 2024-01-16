package admin

import (
	config "web/configs"
)

type AdminHandler struct {
	Services config.ServiceAddr
}

func NewAdminHandler(Services config.ServiceAddr) AdminHandler {
	return AdminHandler{
		Services: Services,
	}
}
