package dao

import (
	"chat/internal/core/entity"

	"github.com/google/uuid"
)

type ConversationPortDAO interface {
	Create(entity.Conversation) error
	Delete(uuid.UUID) error
	Index(int64, int64) ([]entity.Conversation, error)
	Update(entity.Conversation) error
	GetByID(uuid.UUID) (entity.Conversation, error)
	CheckCoversationandUserID(uuid.UUID, uuid.UUID) error
	UserConversations(id uuid.UUID) ([]entity.Conversation, error)
}
