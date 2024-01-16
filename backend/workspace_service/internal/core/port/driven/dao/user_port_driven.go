package dao

import (
	// "workspace/internal/core/entity"

	"github.com/google/uuid"
)

type UserPortDriven interface {
	IsUserExist(uuid.UUID) (bool, error)
}
