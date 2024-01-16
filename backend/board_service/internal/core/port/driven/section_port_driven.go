package driven

import (
	"board/internal/core/entity"

	"github.com/google/uuid"
)

type SectionPortDriven interface {
	GetById(uuid.UUID) (entity.Section, error)
	Index(int, int) ([]entity.Section, error)
	Delete(uuid.UUID) error
	Update(entity.Section) error
	Create(entity.Section) error
	// SectionWorkspaceID(uuid.UUID) (uuid.UUID, error)
}
