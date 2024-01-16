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

type SubTaskUsecase interface {
	SubTasks(dto.SubTasksRequest) ([]dto.SubTasksResponse, error)
	Create(dto.CreateSubTaskRequest) error
	Delete(dto.DeleteSubTaskRequest) error
	Update(dto.UpdateSubTaskRequest) error
}

type subTaskUsecase struct {
	subTaskDao   driven.SubTaskPortDriven
	workspaceDao driven.WorkspacePortDriven
}

func NewSubTaskUsecase(subTaskDao driven.SubTaskPortDriven, workspaceDao driven.WorkspacePortDriven) SubTaskUsecase {
	return &subTaskUsecase{
		subTaskDao:   subTaskDao,
		workspaceDao: workspaceDao,
	}
}

func (s subTaskUsecase) SubTasks(dto.SubTasksRequest) ([]dto.SubTasksResponse, error) {

	return nil, nil
}

func (s subTaskUsecase) Create(subTask dto.CreateSubTaskRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(subTask.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := s.workspaceDao.CheckUserRoleInWorkspace(subTask.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = s.subTaskDao.Create(entity.SubTask{
		ID:     uuid.New(),
		TaskID: subTask.TsakID,
		Title:  subTask.Title,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}

func (s subTaskUsecase) Delete(subTask dto.DeleteSubTaskRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(subTask.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := s.workspaceDao.CheckUserRoleInWorkspace(subTask.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = s.subTaskDao.Delete(subTask.SubTaskID)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}

func (s subTaskUsecase) Update(subTask dto.UpdateSubTaskRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(subTask.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := s.workspaceDao.CheckUserRoleInWorkspace(subTask.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = s.subTaskDao.Update(entity.SubTask{
		Title: subTask.Title,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}
