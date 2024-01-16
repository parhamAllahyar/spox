package usecase

import (
	"board/internal/core/dto"
	"board/internal/core/entity"
	"board/internal/core/port/driven"
	"board/utils/auth/authenticator"
	errMsg "board/utils/message/error"
	"errors"

	"github.com/google/uuid"
	// "fmt"
)

type TagUsecase interface {
	Create(dto.CreateTagRequest) error
	Delete(dto.DeleteTagRequest) error
	Update(dto.UpdateTagRequest) error
}

type tagUsecase struct {
	tagDao       driven.TagPortDriven
	workspaceDao driven.WorkspacePortDriven
}

func NewTagUsecase(tagDao driven.TagPortDriven, workspaceDao driven.WorkspacePortDriven) TagUsecase {
	return &tagUsecase{
		tagDao:       tagDao,
		workspaceDao: workspaceDao,
	}
}

func (t tagUsecase) Create(tag dto.CreateTagRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(tag.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := t.workspaceDao.CheckUserRoleInWorkspace(tag.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = t.tagDao.Create(entity.Tag{
		ID:      uuid.New(),
		Title:   tag.Title,
		Pattern: tag.Pattern,
		Order:   tag.Order,
		BoardID: tag.BoardID,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}
func (t tagUsecase) Delete(tag dto.DeleteTagRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(tag.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// Check board is exist
	workspaceID, err := t.tagDao.TagWorkspaceID(tag.ID)
	if err != nil {
		return errors.New(errMsg.NotFound)
	}

	// check is user has a permission in workspace
	userRole, err := t.workspaceDao.CheckUserRoleInWorkspace(workspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = t.tagDao.Delete(tag.ID)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}
func (t tagUsecase) Update(tag dto.UpdateTagRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(tag.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := t.workspaceDao.CheckUserRoleInWorkspace(tag.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = t.tagDao.Update(entity.Tag{
		Title:   tag.Title,
		Pattern: tag.Pattern,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}
