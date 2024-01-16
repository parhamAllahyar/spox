package dto

import "github.com/google/uuid"

type ConversationListResponse struct {
	ID                   uuid.UUID
	ParticipantFirstName string
	ParticipantLastName  string
	ParticipantImage     string
	LastMessage          string
}
