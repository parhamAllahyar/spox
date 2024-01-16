package dto

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
}

type UpdateUserRequest struct {
	AccessToken string
	FirstName   string
	LastName    string
}

func (u UpdateUserRequest) UpdateUserRequestValidatior() error {

	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Length(5, 50)),
		validation.Field(&u.LastName, validation.Length(8, 50)),
	)

}

type UserListRequest struct {
	AccessToken string
	Page        int
	PageSize    int
}

func (u UserListRequest) UserListRequestValidatior() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Page, validation.Max(70)),
		validation.Field(&u.PageSize, validation.Max(70)),
	)
}

type SearchRequest struct {
	AccessToken string
	Item        string
}

func (s SearchRequest) UserListRequestValidatior() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Item, validation.Max(70)),
	)
}

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
