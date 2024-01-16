package entity

import (
	"time"

	"github.com/google/uuid"
)

type Workspace struct {
	ID        uuid.UUID
	Title     string
	CreatorId uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
