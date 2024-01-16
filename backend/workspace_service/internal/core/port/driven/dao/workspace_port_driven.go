package dao

import (
	"workspace/internal/core/entity"

	"github.com/google/uuid"
)

type WorkspacePortDriven interface {
	Get(id uuid.UUID) (entity.Workspace, error)
	Index(int, int) ([]entity.Workspace, error)
	UserWorkspace(uuid.UUID) ([]entity.Workspace, error)
	Create(workspace entity.Workspace) (entity.Workspace, error)
	Delete(uuid.UUID) error
	Update(workspace entity.Workspace) error
	AssignUserRole(userID uuid.UUID, WorkspaceID uuid.UUID, RoleID uuid.UUID) error
}
