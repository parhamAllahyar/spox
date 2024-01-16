package dto

import (
	"io"

	"github.com/google/uuid"
)

type SaveFileRequest struct {
	ID          uuid.UUID
	AccessToken string
	File        io.Reader
}

type RemoveFileRequest struct {
	ID          uuid.UUID
	AccessToken string
}

type DownloadFileRequest struct {
	ID          uuid.UUID
	AccessToken string
}

type ListFileRequest struct {
	ID          uuid.UUID
	AccessToken string
}

type ListFileResponse struct {
	ID          uuid.UUID
	AccessToken string
}

type DownloadFileResponse struct {
	Link string
}

type Response struct {
	Response string
}
