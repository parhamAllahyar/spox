package entity

import "github.com/google/uuid"

type Conversation struct {
	ID      uuid.UUID
	UsersId []uuid.UUID
}
