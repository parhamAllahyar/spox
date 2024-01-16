package usecase

import (
	"asset/internal/core/dto"
	"asset/internal/core/entity"
	"asset/internal/core/port/driven/dao"
	"asset/internal/core/port/driven/rest"
	"asset/internal/core/port/driven/storage"
	"asset/pkg/authenticator"

	"github.com/google/uuid"
)

type FileUsecase interface {
	Save(dto.SaveFileRequest) error
	Remove(dto.RemoveFileRequest) (dto.Response, error)
	List(dto.ListFileRequest) (dto.ListFileResponse, error)
	DownloadByID(dto.DownloadFileRequest) (dto.DownloadFileResponse, error)
	DownloadByChatID(dto.DownloadFileRequest) (dto.DownloadFileResponse, error)
}

type fileUsecase struct {
	DAO          dao.FileDAO
	Storage      storage.Storage
	Conversation rest.ConversationDAO
}

func NewFileUsecase(DAO dao.FileDAO, Storage storage.Storage) FileUsecase {
	return fileUsecase{
		DAO:     DAO,
		Storage: Storage,
	}
}

func (f fileUsecase) Save(file dto.SaveFileRequest) error {

	// Check user token
	jwt, err := authenticator.DecodeUserToken(file.AccessToken)
	if err != nil {
		//TODO:
	}
	// TODO:Validate file

	// Upload file
	err = f.Storage.ObjectUpload(file.File)

	//add in database
	err = f.DAO.Create(entity.File{
		ID:     uuid.New(),
		Title:  "test",
		UserID: jwt.ID,
		Size:   20,
		Type:   "",
	})

	//Send Response
	return nil
}

func (f fileUsecase) Remove(file dto.RemoveFileRequest) (dto.Response, error) {
	//Check Token
	authenticator.DecodeUserToken(file.AccessToken)
	//Send Response
	return dto.Response{}, nil
}

func (f fileUsecase) List(list dto.ListFileRequest) (dto.ListFileResponse, error) {
	//Check Token

	authenticator.DecodeUserToken(list.AccessToken)

	//Get list from database

	//Send Response
	return dto.ListFileResponse{}, nil
}

func (f fileUsecase) DownloadByID(file dto.DownloadFileRequest) (dto.DownloadFileResponse, error) {
	//Check Token
	userId, err := authenticator.DecodeUserToken(file.AccessToken)

	if err != nil {
		return dto.DownloadFileResponse{}, err
	}

	// Get file Data
	// check Chat Id is null
	fileObject, err := f.DAO.GetByID(file.ID)

	// Check file is accessable
	if userId.ID != fileObject.UserID {
		return dto.DownloadFileResponse{}, nil
	}

	//Create download link
	link, err := f.Storage.ObjectDownloadPerSign(fileObject.FileID)

	if err != nil {
		return dto.DownloadFileResponse{}, err
	}

	return dto.DownloadFileResponse{Link: link}, nil

}

func (f fileUsecase) DownloadByChatID(file dto.DownloadFileRequest) (dto.DownloadFileResponse, error) {

	//Check Token
	userId, err := authenticator.DecodeUserToken(file.AccessToken)

	// Check user is in the conversation
	isInTheConversation, err := f.Conversation.IsUserInConversation(userId.ID.String())

	if isInTheConversation == false || err != nil {
		return dto.DownloadFileResponse{}, nil
	}

	fileObject, err := f.DAO.GetByID(file.ID)

	//Create download link
	link, err := f.Storage.ObjectDownloadPerSign(fileObject.FileID)

	if err != nil {
		//TODO:  Log
	}

	//Send Response
	return dto.DownloadFileResponse{Link: link}, nil

}
