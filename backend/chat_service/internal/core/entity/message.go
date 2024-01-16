package entity

import "github.com/google/uuid"

type Message struct {
	ID        uuid.UUID
	SenderID  uuid.UUID
	Message   string
	ReciverID uuid.UUID
}
