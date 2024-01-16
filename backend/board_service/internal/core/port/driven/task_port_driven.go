package driven

import (
	"board/internal/core/entity"

	"github.com/google/uuid"
)

type TaskPortDriven interface {
	GetById(uuid.UUID) (entity.Task, error)
	Index(int, int) ([]entity.Task, error)
	Create(entity.Task) error
	Update(entity.Task) error
	Delete(uuid.UUID) error
	// TaskWorkspaceID(uuid.UUID) (uuid.UUID, error)
}
