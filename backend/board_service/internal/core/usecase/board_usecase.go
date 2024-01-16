package usecase

import (
	"board/internal/core/dto"
	"board/internal/core/entity"
	"board/internal/core/port/driven"
	"board/utils/auth/authenticator"
	errMsg "board/utils/message/error"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type BoardUsecase interface {
	Boards(dto.BoardsRequest) ([]dto.BoardsResponse, error)
	Board(dto.BoardRequest) (dto.BoardResponse, error)
	Create(dto.CreateBoardRequest) error
	Delete(dto.DeleteBoardRequest) error
	Update(dto.UpdateBoardRequest) error
}

type boardUsecase struct {
	boardDao     driven.BoardPortDriven
	workspaceDao driven.WorkspacePortDriven
}

func NewBoardUsecase(boardDao driven.BoardPortDriven, workspaceDao driven.WorkspacePortDriven) BoardUsecase {
	return &boardUsecase{
		boardDao:     boardDao,
		workspaceDao: workspaceDao,
	}
}

func (b boardUsecase) Boards(data dto.BoardsRequest) ([]dto.BoardsResponse, error) {

	// decode user Token
	user, err := authenticator.DecodeToken(data.AccessToken)
	if err != nil {
		return nil, errors.New(errMsg.UserNotFound)
	}

	// check is user in workspace
	isMemeber, err := b.workspaceDao.CheckUserRoleInWorkspace(data.WorkspaceID, user.Id)

	if err != nil {
		return nil, errors.New(errMsg.Wrong)
	}
	if isMemeber != "admin" {
		return nil, errors.New(errMsg.NotAuthorized)
	}

	// get board list
	boards, err := b.boardDao.WorkspaceBoards(data.WorkspaceID)

	if err != nil {
		return nil, errors.New(errMsg.Wrong)
	}

	boardList := []dto.BoardsResponse{}
	for _, v := range boards {
		boardList = append(boardList, dto.BoardsResponse{
			Title:     v.Title,
			WallPaper: v.WallPaper,
		})
	}

	return boardList, nil

}

func (b boardUsecase) Board(data dto.BoardRequest) (dto.BoardResponse, error) {

	// decode user Token
	user, err := authenticator.DecodeToken(data.AccessToken)
	if err != nil {
		return dto.BoardResponse{}, errors.New(errMsg.UserNotFound)
	}

	// Check Board Is Exist
	board, err := b.boardDao.GetById(data.ID)
	if err != nil {
		return dto.BoardResponse{}, errors.New(errMsg.NotFound)
	}

	// check is user in workspace
	isMemeber, err := b.workspaceDao.CheckUserRoleInWorkspace(board.WorkspaceID, user.Id)

	if err != nil {
		return dto.BoardResponse{}, errors.New(errMsg.Wrong)
	}
	if isMemeber != "admin" {
		return dto.BoardResponse{}, errors.New(errMsg.NotAuthorized)
	}

	bordWithDetails, err := b.boardDao.BordWithDetails(board.ID)

	if err != nil {
		return dto.BoardResponse{}, errors.New(errMsg.Wrong)
	}

	fmt.Println(bordWithDetails)

	return dto.BoardResponse{}, errors.New(errMsg.UserNotFound)

}

func (b boardUsecase) Create(board dto.CreateBoardRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(board.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := b.workspaceDao.CheckUserRoleInWorkspace(board.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = b.boardDao.Create(entity.Board{
		ID:          uuid.New(),
		Title:       board.Title,
		WorkspaceID: board.WorkspaceID,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}

func (b boardUsecase) Delete(board dto.DeleteBoardRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(board.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// Check board is exist
	boardData, err := b.boardDao.GetById(board.ID)
	if err != nil {
		return errors.New(errMsg.NotFound)
	}

	// check is user has a permission in workspace
	userRole, err := b.workspaceDao.CheckUserRoleInWorkspace(boardData.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = b.boardDao.Delete(boardData.ID)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}

func (b boardUsecase) Update(board dto.UpdateBoardRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(board.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := b.workspaceDao.CheckUserRoleInWorkspace(board.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = b.boardDao.Update(entity.Board{
		Title: board.Title,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}
