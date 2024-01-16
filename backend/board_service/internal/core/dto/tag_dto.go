package dto

import "github.com/google/uuid"

type CreateTagRequest struct {
	AccessToken string
	BoardID     uuid.UUID
	Title       string
	Pattern     string
	Order       uint
	WorkspaceID uuid.UUID
}
type CreateTagResponse struct {
	ID      uuid.UUID
	BoardID uuid.UUID
	Title   string
	Pattern string
	Order   uint
}
type DeleteTagRequest struct {
	AccessToken string
	ID          uuid.UUID
	WorkspaceID uuid.UUID
}
type UpdateTagRequest struct {
	AccessToken string
	ID          uuid.UUID
	BoardID     uuid.UUID
	Title       string
	Pattern     string
	Order       uint
	WorkspaceID uuid.UUID
}
type UpdateTagResponse struct {
	ID      uuid.UUID
	BoardID uuid.UUID
	Title   string
	Pattern string
	Order   uint
}
