package dto

import "github.com/google/uuid"

type BoardsRequest struct {
	AccessToken string
	WorkspaceID uuid.UUID
}

type BoardsResponse struct {
	Title     string
	WallPaper string
}

type BoardRequest struct {
	AccessToken string
	ID          uuid.UUID
	WorkspaceID uuid.UUID
}

type BoardResponse struct {
	AccessToken string
	ID          uuid.UUID
}

type CreateBoardRequest struct {
	AccessToken string
	Title       string
	WorkspaceID uuid.UUID
}
type CreateBoardResponse struct {
	ID          uuid.UUID
	Title       string
	WorkspaceID uuid.UUID
}
type DeleteBoardRequest struct {
	ID          uuid.UUID
	AccessToken string
	WorkspaceID uuid.UUID
}
type UpdateBoardRequest struct {
	AccessToken string
	ID          uuid.UUID
	Title       string
	WorkspaceID uuid.UUID
}
type UpdateBoardResponse struct {
	ID          uuid.UUID
	Title       string
	WorkspaceId uuid.UUID
}
