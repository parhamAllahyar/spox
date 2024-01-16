package driver

import (
	"workspace/internal/core/entity"
)

type RolePortDriver interface {
	Get(id uint) (entity.Workspace, error)
	Index() []entity.Workspace
}
