package entity

import "github.com/google/uuid"

type File struct {
	ID     uuid.UUID
	Title  string
	UserID uuid.UUID
	Size   int64
	Type   string
	Bucket string
	FileID string
	ChatAsset
}

//TODO:
type ChatAsset struct {
	ChatID uuid.UUID
}
