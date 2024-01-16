package dto

import "github.com/google/uuid"

type CreateSubTaskRequest struct {
	WorkspaceID uuid.UUID
	AccessToken string
	TsakID      uuid.UUID
	Title       string
}
type CreateSubTaskResponse struct {
	TaskID uuid.UUID
	Title  string
	ID     uuid.UUID
}

type DeleteSubTaskRequest struct {
	AccessToken string
	SubTaskID   uuid.UUID
	WorkspaceID uuid.UUID
}

type UpdateSubTaskRequest struct {
	AccessToken string
	TaskID      uuid.UUID
	Title       string
	UserID  uuid.UUID
	WorkspaceID uuid.UUID
}

type UpdateSubTaskResponse struct {
	ID         uuid.UUID
	TaskID     uuid.UUID
	Title      string
	CustomerID uuid.UUID
}

type SubTasksRequest struct {
	AccessToken string
	TaskID      uuid.UUID
	WorkspaceID uuid.UUID
}

type SubTasksResponse struct {
}
