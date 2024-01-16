package driven

import (
	"board/internal/core/entity"

	"github.com/google/uuid"
)

type SubTaskPortDriven interface {
	GetById(uuid.UUID) (entity.SubTask, error)
	Index(int, int) ([]entity.SubTask, error)
	Create(entity.SubTask) error
	Update(entity.SubTask) error
	Delete(uuid.UUID) error
	// SubTaskWorkspaceID(uuid.UUID) (uuid.UUID, error)
}
