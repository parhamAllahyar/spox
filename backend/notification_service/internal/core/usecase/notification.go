package usecase

import "notification/internal/core/port/driven"

// // "board/internal/core/dto"

// "notification/internal/core/port/driven"
// "board/utils/auth/authenticator"
// errMsg "board/utils/message/error"
// "errors"
// "fmt"

// "github.com/google/uuid"

type NotificationUsecase interface {
	SMS()
	Email()
}

type notificationUsecase struct {
	EmailProvider driven.Email
	SMSProvider   driven.SMS
}

func NewBoardUsecase(EmailProvider driven.Email, SMSProvider driven.SMS) NotificationUsecase {
	return &notificationUsecase{
		EmailProvider: EmailProvider,
		SMSProvider:   SMSProvider,
	}
}

func (n notificationUsecase) SMS() {

	// Send By priprity

}
func (n notificationUsecase) Email() {

}
