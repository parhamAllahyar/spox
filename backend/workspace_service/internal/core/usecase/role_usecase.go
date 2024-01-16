package usecase

import (
	"errors"
	"workspace/internal/core/dto"
	"workspace/internal/core/entity"
	"workspace/internal/core/port/driven/dao"
	"workspace/utils/auth/authenticator"
	errMsg "workspace/utils/message/error"

	"github.com/google/uuid"
)

type RoleUsecase interface {
	Create(dto.CreateRoleRequest) error
	Update(dto.UpdateRoleRequest) error
	Delete(dto.DeleteRoleRequest) error
	AssignRoleToUser(dto.AssignRoleRequest) error
	RoleList() ([]dto.RoleListResponse, error)
	CheckUserRole(uuid.UUID) (string, error)
}

type roleUsecase struct {
	roleDao dao.RolePortDriven
	// logger  logging.Logging
}

func NewRoleUsecase(roleDao dao.RolePortDriven) RoleUsecase {
	return &roleUsecase{
		roleDao: roleDao,
		// logger:  logger,
	}
}

func (r *roleUsecase) Create(role dto.CreateRoleRequest) error {

	//Check is admin
	_, err := authenticator.IsAdminToken(role.AccessToken)
	if err != nil {
		return errors.New(errMsg.NotAuthorized)
	}

	//Validation
	err = role.CreateRoleRequestValidator()
	if err != nil {
		return err
	}

	// Create
	err = r.roleDao.Create(entity.Role{
		ID:    uuid.New(),
		Title: role.Title,
	})

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	return nil

}

func (r *roleUsecase) Update(role dto.UpdateRoleRequest) error {

	//Check is admin
	_, err := authenticator.IsAdminToken(role.AccessToken)
	if err != nil {
		return errors.New(errMsg.NotAuthorized)
	}

	//Validation
	err = role.UpdateRoleRequestValidator()
	if err != nil {
		return err
	}

	// Create
	err = r.roleDao.Update(entity.Role{
		ID:    uuid.New(),
		Title: role.Title,
	})

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	return nil

}

func (r *roleUsecase) Delete(role dto.DeleteRoleRequest) error {
	// Check is admin
	_, err := authenticator.IsAdminToken(role.AccessToken)
	if err != nil {
		return errors.New(errMsg.NotAuthorized)
	}

	// Delete
	err = r.roleDao.Delete(role.ID)

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	return nil

}

func (r *roleUsecase) AssignRoleToUser(dto.AssignRoleRequest) error {

	return nil
}

func (r *roleUsecase) RoleList() ([]dto.RoleListResponse, error) {

	// Delete
	roles, err := r.roleDao.Index(12, 12)
	if err != nil {
		return nil, errors.New(errMsg.Wrong)
	}
	roleList := []dto.RoleListResponse{}
	for _, v := range roles {
		roleList = append(roleList, dto.RoleListResponse{
			ID:    v.ID,
			Title: v.Title,
		})
	}

	return roleList, nil

}

func (r *roleUsecase) CheckUserRole(uuid.UUID) (string, error) {
	return "", nil
}
