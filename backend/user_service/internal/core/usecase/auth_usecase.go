package usecase

import (
	"errors"
	"fmt"
	"user/internal/core/dto"
	"user/internal/core/entity"
	"user/internal/core/port/driven/dao"
	"user/utils/auth/authenticator"
	"user/utils/auth/encryption"
	errMsg "user/utils/message/error"

	"github.com/google/uuid"
)

type AuthUsecase interface {
	SignUp(dto.SignUpRequest) (string, error)
	//ForgetPassword(customer dto.ForgetPassword) (dto.AuthResponse, error)
	SignIn(dto.SignInRequest) (string, error)
	UpdatePassword(dto.UpdatePasswordRequest) error
}

// interface
type authUsecase struct {
	userDao dao.UserPortDriven
}

func NewAuthUsecase(userDao dao.UserPortDriven) AuthUsecase {
	return &authUsecase{
		userDao: userDao,
	}
}

func (a *authUsecase) SignIn(user dto.SignInRequest) (string, error) {

	validationError := user.SignInRequestValidator()

	if validationError != nil {
		return "", validationError
	}

	userInfo, err := a.userDao.GetByEmail(user.Email)

	if err != nil {
		return "", errors.New(errMsg.UserNotFound)
	}

	isPasswordValid := encryption.ComparePasswords(userInfo.Password, user.Password)

	if !isPasswordValid {
		return "", errors.New(errMsg.SignInError)
	}

	token, err := authenticator.ExtractClaims(userInfo.ID)

	if err != nil {
		return "", errors.New(errMsg.Wrong)
	}

	return token, nil

}

func (a *authUsecase) SignUp(user dto.SignUpRequest) (string, error) {

	err := user.SignUpRequestValidator()

	if err != nil {
		return "", err
	}

	oldUser, err := a.userDao.GetByEmail(user.Email)

	if err != nil || (oldUser.Email != "") {
		return "", errors.New(errMsg.UserExist)
	}

	id := uuid.New()

	hashPassword, err := encryption.HashPassword(user.Password)
	if err != nil {
		fmt.Print(hashPassword)
	}

	newUserData := entity.User{
		ID:       id,
		Password: hashPassword,
		Email:    user.Email,
	}

	err = a.userDao.Create(newUserData)
	if err != nil {
		return "", err
	}

	token, err := authenticator.ExtractClaims(newUserData.ID)

	if err != nil {
		return "", errors.New(errMsg.Wrong)
	}

	return token, nil

}

func (a *authUsecase) ForgetPassword(user dto.ForgetPasswordRequest) (dto.User, error) {

	//TODO:

	return dto.User{}, nil
}

func (c *authUsecase) UpdatePassword(passwords dto.UpdatePasswordRequest) error {

	//TODO: check password
	tokenData, err := authenticator.DecodeToken(passwords.AccessToken)
	if err != nil {
		return err
	}

	err = passwords.UpdatePasswordValidatior()

	customer, err := c.userDao.GetByID(tokenData.Id)
	if err != nil {
		return err
	}

	isPasswordValid := encryption.ComparePasswords(customer.Password, passwords.OldPassword)

	if !isPasswordValid {
		return errors.New(errMsg.PasswordMisMatch)
	}

	password, err := encryption.HashPassword(passwords.NewPassword)

	if err != nil {
		return err
	}

	err = c.userDao.UpdatePassword(tokenData.Id, password)

	if err != nil {
		return err
	}

	return nil

}
