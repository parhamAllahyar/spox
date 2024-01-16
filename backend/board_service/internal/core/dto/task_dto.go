package dto

// CreateTask(token, title string, id string) (dto.CreateTaskResponse, error)
// DeleteTask(token string, id uuid.UUID) error
// UpdateTask(dto.UpdateTaskRequest) (dto.UpdateTaskResponse, error)

import (
	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	AccessToken string
	Title       string
	SectionID   uuid.UUID
	Order       uint
	WorkspaceID uuid.UUID
}

type CreateTaskResponse struct {
	ID        uuid.UUID
	Title     string
	SectionID uuid.UUID
	Order     uint
}

type DeleteTaskRequest struct {
	AccessToken string
	TaskID      uuid.UUID
	WorkspaceID uuid.UUID
}

type UpdateTaskRequest struct {
	AccessToken string
	ID          uuid.UUID
	Title       string
	SectionID   uuid.UUID
	Order       uint
	WorkspaceID uuid.UUID
}
type UpdateTaskResponse struct {
	ID        uuid.UUID
	Title     string
	SectionID uuid.UUID
	Order     uint
}

type AssignTaskRequest struct {
	AccessToken string
	TaskID      uuid.UUID
	UserID  uuid.UUID
	WorkspaceID uuid.UUID
}
