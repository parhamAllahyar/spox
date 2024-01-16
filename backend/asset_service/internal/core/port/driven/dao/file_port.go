package dao

import (
	"asset/internal/core/entity"

	"github.com/google/uuid"
)

type FileDAO interface {
	Create(entity.File) error
	Delete(uuid.UUID) error
	List(uuid.UUID) ([]entity.File, error)
	GetByID(uuid.UUID) (entity.File, error)
}
