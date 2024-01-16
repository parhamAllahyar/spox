package dao

import (
	"workspace/internal/core/entity"

	"github.com/google/uuid"
)

type InvitationPortDriven interface {
	Create(entity.Invitation) error
	GetByID(uuid.UUID) (entity.Invitation, error)
}
