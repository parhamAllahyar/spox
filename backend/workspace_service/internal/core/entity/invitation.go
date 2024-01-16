package entity

import (
	"time"

	"github.com/google/uuid"
)

type Invitation struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	WorkspaceID uuid.UUID
	RoleID      uuid.UUID
	ExpireDate  time.Time
}
