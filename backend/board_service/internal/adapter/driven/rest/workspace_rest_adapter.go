package postgres

import (
	// "board/internal/core/entity"

	"github.com/google/uuid"
)

type workspaceAPI struct {
}

func NewWorkspaceAPI() *workspaceAPI {
	return &workspaceAPI{
		// db: db,
	}
}

func (w workspaceAPI) CheckUserRoleInWorkspace(uuid.UUID, uuid.UUID) (string, error) {

	return "", nil
}
