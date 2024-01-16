package dao

import (
	"user/internal/core/entity"

	"github.com/google/uuid"
)

type UserPortDriven interface {
	Create(entity.User) error
	Update(entity.User) error
	Delete(uuid.UUID) error
	GetByID(uuid.UUID) (entity.User, error)
	GetByEmail(string) (entity.User, error)
	Index(int, int) ([]entity.User, error)
	Search(string) ([]entity.User, error)
	UpdatePhone(uuid.UUID, string) error
	UpdateEmail(uuid.UUID, string) error
	UpdatePassword(uuid.UUID, string) error
	Count() (int64, error)
	SearchByEmail(string) (entity.User, error)
}
