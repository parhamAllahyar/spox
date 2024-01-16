package dao

import (
	"chat/internal/core/entity"

	"github.com/google/uuid"
)

type UserPortDAO interface {
	UserList(ids []uuid.UUID) ([]entity.User, error)
}
