package usecase

import (
	"board/internal/core/dto"
	"board/internal/core/entity"
	"board/internal/core/port/driven"
	"board/utils/auth/authenticator"
	errMsg "board/utils/message/error"
	"errors"

	"github.com/google/uuid"
)

type SectionUsecase interface {
	Create(dto.CreateSectionRequest) error
	Delete(dto.DeleteSectionRequest) error
	Update(dto.UpdateSectionRequest) error
}

type sectionUsecase struct {
	sectionDao   driven.SectionPortDriven
	workspaceDao driven.WorkspacePortDriven
}

func NewSectionUsecase(sectionDao driven.SectionPortDriven, workspaceDao driven.WorkspacePortDriven) SectionUsecase {
	return &sectionUsecase{
		sectionDao:   sectionDao,
		workspaceDao: workspaceDao,
	}
}

func (s sectionUsecase) Create(section dto.CreateSectionRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(section.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := s.workspaceDao.CheckUserRoleInWorkspace(section.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = s.sectionDao.Create(entity.Section{
		ID:      uuid.New(),
		BoardID: section.BoardID,
		Title:   section.Title,
		Order:   section.Order,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}

func (s sectionUsecase) Delete(section dto.DeleteSectionRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(section.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := s.workspaceDao.CheckUserRoleInWorkspace(section.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = s.sectionDao.Delete(section.SectionID)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}

func (s sectionUsecase) Update(section dto.UpdateSectionRequest) error {
	// decode user Token
	user, err := authenticator.DecodeToken(section.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := s.workspaceDao.CheckUserRoleInWorkspace(section.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = s.sectionDao.Update(entity.Section{
		Title: section.Title,
		Order: section.Order,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}
