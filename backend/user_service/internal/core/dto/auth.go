package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignInRequest struct {
	Password string
	Email    string
}

func (s SignInRequest) SignInRequestValidator() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Email, validation.Required, validation.Length(5, 50), is.Email),
		validation.Field(&s.Password, validation.Required, validation.Length(8, 50)),
	)
}

type SignUpRequest struct {
	Password string
	Email    string
}

func (s SignUpRequest) SignUpRequestValidator() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Email, validation.Required, validation.Length(5, 50), is.Email),
		validation.Field(&s.Password, validation.Required, validation.Length(8, 50), validation.LengthRule{}),
	)
}

type AuthResponse struct {
	AccessToken string
}

type ForgetPasswordRequest struct{}

type UpdatePasswordRequest struct {
	AccessToken string
	OldPassword string
	NewPassword string
}

func (s UpdatePasswordRequest) UpdatePasswordValidatior() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.OldPassword, validation.Required, validation.Length(8, 50)),
		validation.Field(&s.NewPassword, validation.Required, validation.Length(8, 50)),
	)
}
