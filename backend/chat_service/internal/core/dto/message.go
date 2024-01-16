package dto

import "github.com/google/uuid"

type PastMessagesResponse struct {
	ID       uuid.UUID
	Message  string
	SenderID uuid.UUID
}

type PastMessageRequest struct {
	AccessToken string
	Page        int64
	PageSize    int64
	ChatID      uuid.UUID
}

type CreateTextMessageRequest struct {
	Message string
	Token   string
	ChatID  string
}

type UpdateMessageRequest struct {
	MessageID uuid.UUID
	Message   string
	Token     string
}

type DeleteMessageRequest struct {
	MessageID uuid.UUID
	Token     string
}

type CreateFileMessageRequest struct {
	Message string
	Token   string
	ChatID  string
}

type StartConversation struct {
	Channel     string
	Key         string
	AccessToken string
}
