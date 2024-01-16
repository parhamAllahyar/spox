package driven

import (
	"board/internal/core/entity"

	"github.com/google/uuid"
)

type TagPortDriven interface {
	GetById(uuid.UUID) (entity.Tag, error)
	Index(int, int) ([]entity.Tag, error)
	Create(entity.Tag) error
	Delete(uuid.UUID) error
	Update(entity.Tag) error
	// TagWorkspaceID(uuid.UUID) (uuid.UUID, error)
}
