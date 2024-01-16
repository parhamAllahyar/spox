package driven

import (
	"board/internal/core/entity"

	"github.com/google/uuid"
)

type BoardPortDriven interface {
	Create(entity.Board) error
	Update(entity.Board) error
	GetById(uuid.UUID) (entity.Board, error)
	Index(int, int) ([]entity.Board, error)
	Delete(uuid.UUID) error
	WorkspaceBoards(uuid.UUID) ([]entity.Board, error)
	BordWithDetails(uuid.UUID) (entity.Board, error)
}
