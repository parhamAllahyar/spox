package dao

import (
	"chat/internal/core/entity"

	"github.com/google/uuid"
)

type MessagePortDAO interface {
	Create(entity.Message) error
	Delete(uuid.UUID) error
	Index(uuid.UUID, int64, int64) ([]entity.Message, error)
	Update(entity.Message) error
}
