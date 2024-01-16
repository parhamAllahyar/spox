package driver

import (
	"workspace/internal/core/entity"
)

type WorkspacePortDriver interface {
	Get(id uint) (entity.Workspace, error)
	Index() []entity.Workspace
}
