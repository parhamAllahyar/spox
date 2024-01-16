package usecase

import (
	"errors"

	"user/internal/core/dto"
	"user/internal/core/entity"
	"user/internal/core/port/driven/dao"
	"user/utils/auth/authenticator"
	errMsg "user/utils/message/error"
)

type UserUsecase interface {
	Count(string) (int64, error)
	Users(dto.UserListRequest) ([]dto.User, error)
	Search(dto.SearchRequest) ([]dto.User, error)
	Update(dto.UpdateUserRequest) error
	SearchByEmail(dto.SearchRequest) (dto.User, error)
}

// interface
type userUsecase struct {
	userDao dao.UserPortDriven
}

func NewUserUsecase(userDao dao.UserPortDriven) UserUsecase {
	return &userUsecase{
		userDao: userDao,
	}
}

func (u *userUsecase) Count(string) (int64, error) {
	count, err := u.userDao.Count()
	return count, err
}

func (u *userUsecase) Users(data dto.UserListRequest) ([]dto.User, error) {

	isAdmin := authenticator.IsAdminToken(data.AccessToken)

	if isAdmin == false {
		return nil, errors.New(errMsg.Wrong)
	}

	users, err := u.userDao.Index(data.Page, data.PageSize)

	if err != nil {
		return nil, errors.New(errMsg.Wrong)
	}

	userList := []dto.User{}
	for _, v := range users {
		userList = append(userList, dto.User{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
			Phone:     v.Phone,
		})
	}

	return userList, nil
}

func (u *userUsecase) Search(searchItem dto.SearchRequest) ([]dto.User, error) {

	isAdmin := authenticator.IsAdminToken(searchItem.AccessToken)

	if isAdmin == false {
		return nil, errors.New(errMsg.Wrong)
	}

	users, err := u.userDao.Search(searchItem.Item)

	if err != nil {
		return nil, errors.New(errMsg.Wrong)
	}

	userList := []dto.User{}
	for _, v := range users {
		userList = append(userList, dto.User{
			ID:        v.ID,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Email:     v.Email,
			Phone:     v.Phone,
		})
	}

	return userList, nil

}

func (u *userUsecase) Update(user dto.UpdateUserRequest) error {

	err := user.UpdateUserRequestValidatior()

	if err != nil {
		return err
	}

	token, err := authenticator.DecodeToken(user.AccessToken)

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	err = u.userDao.Update(entity.User{
		ID:        token.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})

	if err != nil {
		return errors.New(errMsg.Wrong)
	}

	return nil

}

func (u *userUsecase) SearchByEmail(searchItem dto.SearchRequest) (dto.User, error) {

	_, err := authenticator.DecodeToken(searchItem.AccessToken)

	if err != nil {
		return dto.User{}, errors.New(errMsg.Wrong)
	}

	user, err := u.userDao.SearchByEmail(searchItem.Item)

	if err != nil {
		return dto.User{}, errors.New(errMsg.Wrong)
	}

	return dto.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil

}
