package dao

import (
	"workspace/internal/core/entity"

	"github.com/google/uuid"
)

type RolePortDriven interface {
	Get(string) (entity.Role, error)
	Index(int, int) ([]entity.Role, error)
	Create(entity.Role) error
	Update(entity.Role) error
	Delete(uuid.UUID) error
}
