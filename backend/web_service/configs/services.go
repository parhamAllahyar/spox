package configs

import "fmt"

type ServiceAddr struct {
	User      string
	Workspace string
	Board     string
}

func ServicesAddres(services Services) ServiceAddr {
	return ServiceAddr{
		User:      fmt.Sprintf("%s%s%s%s%s", "http://", services.UserHost, ":", services.UserPort, "/"),
		Workspace: fmt.Sprintf("%s%s", services.WorkspaceHost, services.WorkspacePORT),
		Board:     fmt.Sprintf("%s%s", services.BoardHost, services.BoardPort),
	}
}
