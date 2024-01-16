package validator

import (
	"errors"
	"regexp"
	errMsg "user/utils/message/error"
)

func EmailValidator(email string) error {

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New(errMsg.EmailFormat)
	}
	return nil
}

func RequiredValidator(fields []any) error {
	for _, field := range fields {
		if field.(string) == "" || field.(int) == 0 {
			return errors.New(errMsg.Required)
		} else {
			return errors.New(errMsg.TypeUnsupport)
		}
	}
	return nil
}

func FirstNameValidator(firstName string) error {
	if len(firstName) < 2 || len(firstName) > 50 {
		return errors.New(errMsg.FirstNameError)
	}
	return nil
}

func LastNameValidator(lastName string) error {
	if len(lastName) < 2 || len(lastName) > 50 {
		return errors.New(errMsg.LastNameError)
	}
	return nil
}

func PhoneValidator(phone string) error {
	phoneRegex := regexp.MustCompile(`^\+\d{1,2}\d{3,14}$`)
	if !phoneRegex.MatchString(phone) {
		return errors.New(errMsg.PhoneError)
	}
	return nil
}
