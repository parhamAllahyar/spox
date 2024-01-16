package entity

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Avatar    string
}
