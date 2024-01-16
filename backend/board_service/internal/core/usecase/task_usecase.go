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

type TaskUsecase interface {
	Tasks(dto.CreateTaskRequest) (dto.CreateTaskResponse, error)
	Create(dto.CreateTaskRequest) error
	Delete(dto.DeleteTaskRequest) error
	Update(dto.UpdateTaskRequest) error
	AssignTask(dto.AssignTaskRequest) error
}

type taskUsecase struct {
	taskDao      driven.TaskPortDriven
	workspaceDao driven.WorkspacePortDriven
}

func NewTaskUsecase(taskDao driven.TaskPortDriven, workspaceDao driven.WorkspacePortDriven) TaskUsecase {
	return &taskUsecase{
		taskDao:      taskDao,
		workspaceDao: workspaceDao,
	}
}

func (t *taskUsecase) Tasks(dto.CreateTaskRequest) (dto.CreateTaskResponse, error) {
	return dto.CreateTaskResponse{}, nil
}

func (t *taskUsecase) Create(task dto.CreateTaskRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(task.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := t.workspaceDao.CheckUserRoleInWorkspace(task.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = t.taskDao.Create(entity.Task{
		ID:        uuid.New(),
		Title:     task.Title,
		Order:     task.Order,
		SectionID: task.SectionID,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}

func (t *taskUsecase) Delete(task dto.DeleteTaskRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(task.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := t.workspaceDao.CheckUserRoleInWorkspace(task.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = t.taskDao.Delete(task.TaskID)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}

func (t *taskUsecase) Update(task dto.UpdateTaskRequest) error {

	// decode user Token
	user, err := authenticator.DecodeToken(task.AccessToken)
	if err != nil {
		return errors.New(errMsg.UserNotFound)
	}

	// check is user has a permission in workspace
	userRole, err := t.workspaceDao.CheckUserRoleInWorkspace(task.WorkspaceID, user.Id)
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = t.taskDao.Update(entity.Task{
		Title: task.Title,
		Order: task.Order,
	})
	if err != nil {
		return errors.New(errMsg.Wrong)
	}
	return nil

}

// TODO:
func (t *taskUsecase) AssignTask(task dto.AssignTaskRequest) error {

	return nil
}
