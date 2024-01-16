package dto

import "github.com/google/uuid"

type AddMeberToWorkspaceRequest struct {
	AccessToken    string
	UserId         uuid.UUID
	InvitationLink uuid.UUID
}

func (a AddMeberToWorkspaceRequest) AddMeberToWorkspaceRequestValidator() error {
	return nil
}

type WorkspaceMembersRequest struct {
	AccessToken string
	WorkspaceId uuid.UUID
}

type WorkspaceMembersResponse struct{}

func (w WorkspaceMembersRequest) WorkspaceMembersRequestValidator() error {
	return nil
}

type DeleteMembersRequest struct {
	AccessToken string
	WorkspaceId uuid.UUID
	UserId      uuid.UUID
}

func (w DeleteMembersRequest) DeleteMembersRequestValidator() error {
	return nil
}

type InvitRequest struct {
	AccessToken string
	WorkspaceId uuid.UUID
	UserId      uuid.UUID
	RoleID      uuid.UUID
}

type InvitResponse struct {
	AccessToken string
	WorkspaceId uuid.UUID
	UserId      uuid.UUID
}
