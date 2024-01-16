package usecase

import (
	"errors"
	"fmt"
	"workspace/internal/core/dto"
	"workspace/internal/core/entity"
	"workspace/internal/core/port/driven/dao"
	"workspace/utils/auth/authenticator"
	errMsg "workspace/utils/message/error"
)

type WorkspaceUsecase interface {
	WorkspacesById(dto.ShowWorkspaceByIDRequest) ([]dto.ShowWorkspaceByIDResponse, error)
	Create(dto.CreateWorkspaceRequest) error
	Update(dto.UpdateWorkspaceRequest) error
	Delete(dto.DeleteWorkspaceByIDRequest) error
	Index(token string) ([]dto.UserWorkspaceListResponse, error)
}

type workspaceUsecase struct {
	workspaceDao dao.WorkspacePortDriven
	roleUsecase  RoleUsecase
	roleDao      dao.RolePortDriven
}

func NewWorkspaceUsecase(workspaceDao dao.WorkspacePortDriven, roleUsecase RoleUsecase, roleDao dao.RolePortDriven) WorkspaceUsecase {
	return &workspaceUsecase{
		workspaceDao: workspaceDao,
		roleUsecase:  roleUsecase,
		roleDao:      roleDao,
	}
}

func (w *workspaceUsecase) WorkspacesById(dto.ShowWorkspaceByIDRequest) ([]dto.ShowWorkspaceByIDResponse, error) {
	return nil, nil
}

func (w *workspaceUsecase) Create(workspace dto.CreateWorkspaceRequest) error {

	user, err := authenticator.DecodeToken(workspace.AccessToken)

	if err != nil {
		return errors.New(errMsg.NotAuthorized)
	}

	//Validation
	err = workspace.CreateWorkspaceRequestValidator()
	if err != nil {
		return err
	}

	newWorkspace, err := w.workspaceDao.Create(entity.Workspace{
		Title:     workspace.Title,
		CreatorId: user.Id,
	})

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	// TODO:Add to Member
	role, err := w.roleDao.Get("admin")

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	err = w.workspaceDao.AssignUserRole(newWorkspace.CreatorId, newWorkspace.ID, role.ID)

	fmt.Println("err e", err)

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	return nil

}

func (w *workspaceUsecase) Update(workspace dto.UpdateWorkspaceRequest) error {

	//Validation
	err := workspace.UpdateWorkspaceRequestValidator()
	if err != nil {
		return err
	}

	//Check is admin
	user, err := authenticator.DecodeToken(workspace.AccessToken)
	if err != nil {
		return errors.New(errMsg.NotAuthorized)
	}

	// Check user role is admin
	userRole, err := w.roleUsecase.CheckUserRole(user.Id)

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = w.workspaceDao.Update(entity.Workspace{
		ID:    workspace.ID,
		Title: workspace.Title,
	})

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	return nil

}

func (w *workspaceUsecase) Delete(workspace dto.DeleteWorkspaceByIDRequest) error {

	//Check is admin
	user, err := authenticator.DecodeToken(workspace.AccessToken)
	if err != nil {
		return errors.New(errMsg.NotAuthorized)
	}

	// Check user role is admin
	userRole, err := w.roleUsecase.CheckUserRole(user.Id)

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	if userRole != "admin" {
		return errors.New(errMsg.NotAuthorized)
	}

	err = w.workspaceDao.Delete(workspace.ID)

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	return nil

}

func (w *workspaceUsecase) Index(token string) ([]dto.UserWorkspaceListResponse, error) {

	//Check is admin
	user, err := authenticator.DecodeToken(token)
	if err != nil {
		return nil, errors.New(errMsg.NotAuthorized)
	}

	workspaceList, err := w.workspaceDao.UserWorkspace(user.Id)

	if err != nil {
		return nil, errors.New(errMsg.Wrong)
	}

	workspaces := []dto.UserWorkspaceListResponse{}

	for _, v := range workspaceList {

		workspaces = append(workspaces, dto.UserWorkspaceListResponse{
			ID:    v.ID,
			Title: v.Title,
		})

	}

	return workspaces, nil

}
