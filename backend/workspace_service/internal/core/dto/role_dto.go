package dto

import "github.com/google/uuid"

type CreateRoleRequest struct {
	Title       string `json:"title"`
	AccessToken string
}

func (c CreateRoleRequest) CreateRoleRequestValidator() error {
	return nil
}

type CreateRoleResponse struct {
	ID    uuid.UUID
	Title string `json:"title"`
}

type UpdateRoleRequest struct {
	Title       string `json:"title"`
	AccessToken string
	ID          uuid.UUID
}

func (c UpdateRoleRequest) UpdateRoleRequestValidator() error {
	return nil
}

type UpdateRoleResponse struct {
	ID    uuid.UUID
	Title string `json:"title"`
}

type DeleteRoleRequest struct {
	AccessToken string
	ID          uuid.UUID
}

func (c DeleteRoleRequest) DeleteRoleRequestValidator() error {
	return nil
}

type ShowRoleByIDRequest struct {
	AccessToken string
	ID          uuid.UUID
}

func (c ShowRoleByIDRequest) ShowRoleByIDRequestValidator() error {
	return nil
}

type RoleListResponse struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type AssignRoleRequest struct {
	AccessToken string
	RoleID      uuid.UUID `json:"role_id"`
	UserID      uuid.UUID `json:"user_id"`
}

func (c AssignRoleRequest) AssignRoleRequestValidator() error {
	return nil
}
