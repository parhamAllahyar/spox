package usecase

import (
	// "errors"
	"workspace/internal/core/dto"
	"workspace/internal/core/entity"
	"workspace/internal/core/port/driven/dao"
	"workspace/utils/auth/authenticator"
	// errMsg "workspace/utils/message/error"
	// "github.com/google/uuid"
)

type MembershipUsecase interface {
	AddUserTOWorkspace(dto.AddMeberToWorkspaceRequest) error
	WorkspaceMembers(member dto.WorkspaceMembersRequest) ([]dto.WorkspaceMembersResponse, error)
	Invitation(member dto.InvitRequest) error
	// IsUserWorkspaceMember() (bool , error)
}

type membershipUsecase struct {
	workspaceDao dao.WorkspacePortDriven
	userDao      dao.UserPortDriven
	invitation   dao.InvitationPortDriven
	roleUsecase  RoleUsecase
}

func NewMembershipUsecase(workspaceDao dao.WorkspacePortDriven, userDao dao.UserPortDriven, roleUsecase RoleUsecase, invitationDao dao.InvitationPortDriven) MembershipUsecase {
	return &membershipUsecase{
		workspaceDao: workspaceDao,
		invitation:   invitationDao,
		userDao:      userDao,
		roleUsecase:  roleUsecase,
	}
}

func (m *membershipUsecase) AddUserTOWorkspace(member dto.AddMeberToWorkspaceRequest) error {

	// Check token

	token, err := authenticator.DecodeToken(member.AccessToken)

	if err != nil {
		return err
	}

	// get invit link by id
	invitation, err := m.invitation.GetByID(member.InvitationLink)

	if err != nil {
		return err
	}

	if invitation.UserID != token.Id {
		return nil
	}

	// add user to workspace
	err = m.workspaceDao.AssignUserRole(invitation.UserID, invitation.WorkspaceID, invitation.RoleID)

	return nil
	// //Check token
	// user, err := authenticator.DecodeToken(member.AccessToken)
	// if err != nil {
	// 	return errors.New(errMsg.NotAuthorized)
	// }

	// // Check user role is admin
	// userRole, err := m.roleUsecase.CheckUserRole(user.Id)

	// if err != nil {
	// 	return errors.New(errMsg.Wrong)
	// }

	// if userRole != "admin" {
	// 	return errors.New(errMsg.NotAuthorized)
	// }

	// // Check is user exist
	// isExist, err := m.userDao.IsUserExist(member.UserId)

	// if err != nil {
	// 	return errors.New(errMsg.Wrong)
	// }

	// if isExist != true {
	// 	return errors.New(errMsg.NotAuthorized)
	// }

	// // m.memberDao.AddMeber(entity.Member{
	// 	WorkspaceID: member.WorkspaceId,
	// 	UserID:      member.UserId,
	// 	Role:        member.Role,
	// })

}

// Invitation
func (m *membershipUsecase) Invitation(member dto.InvitRequest) error {

	// get user id by token
	token, err := authenticator.DecodeToken(member.AccessToken)

	if err != nil {
		return err
	}

	// Check user has a role
	role, err := m.roleUsecase.CheckUserRole(token.Id)

	if err != nil {
		return err
	}

	if role != "admin" {
		return nil
	}

	// create invitation

	// TODO:
	err = m.invitation.Create(entity.Invitation{
		UserID:      token.Id,
		WorkspaceID: member.WorkspaceId,
		RoleID:      member.RoleID,
	})

	// TODO: Send notification

	return nil
}

func (m *membershipUsecase) WorkspaceMembers(member dto.WorkspaceMembersRequest) ([]dto.WorkspaceMembersResponse, error) {

	return nil, nil

}
