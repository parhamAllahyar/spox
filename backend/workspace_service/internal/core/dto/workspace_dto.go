package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type CreateWorkspaceRequest struct {
	Title       string `json:"title"`
	AccessToken string
}

func (c CreateWorkspaceRequest) CreateWorkspaceRequestValidator() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required, validation.Length(5, 50)),
	)
}

type CreateWorkspaceResponse struct {
	ID    uuid.UUID
	Title string `json:"title"`
}

type UpdateWorkspaceRequest struct {
	Title       string `json:"titile"`
	AccessToken string
	ID          uuid.UUID
}

func (c UpdateWorkspaceRequest) UpdateWorkspaceRequestValidator() error {
	return nil
}

type UpdateWorkspaceResponse struct {
	ID    uuid.UUID
	Title string `json:"title"`
}

type WorkspacesResponse struct {
	ID    uuid.UUID
	Title string `json:"title"`
}

type ShowWorkspaceByIDRequest struct {
	AccessToken string
	ID          uuid.UUID
}

func (c ShowWorkspaceByIDRequest) ShowWorkspaceByIDRequestValidator() error {
	return nil
}

type ShowWorkspaceByIDResponse struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type DeleteWorkspaceByIDRequest struct {
	AccessToken string
	ID          uuid.UUID
}

func (c DeleteWorkspaceByIDRequest) DeleteWorkspaceByIDRequestValidator() error {
	return nil
}

type UserWorkspaceListResponse struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}
