package entity

import "github.com/google/uuid"

type Role struct {
	ID    uuid.UUID
	Title string
}
