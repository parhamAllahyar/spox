package usecase

import (
	"chat/internal/core/dto"
	"chat/internal/core/port/driven/dao"
	"fmt"
)

type MessageUsecase interface {
	CreateTextMessage(dto.CreateTextMessageRequest) error
	PastMessages(dto.PastMessageRequest) ([]dto.PastMessagesResponse, error)
	Update(message dto.UpdateMessageRequest)
	Delete(dto.DeleteMessageRequest)
}

type messageUsecase struct {
	DAO dao.MessagePortDAO
}

func NewMessageUsecase(DAO dao.MessagePortDAO) MessageUsecase {
	return messageUsecase{
		DAO: DAO,
	}
}

func (m messageUsecase) CreateTextMessage(message dto.CreateTextMessageRequest) error {

	return nil // m.DAO.Create()
}

func (m messageUsecase) PastMessages(data dto.PastMessageRequest) ([]dto.PastMessagesResponse, error) {

	// check
	messages, err := m.DAO.Index(data.ChatID, data.Page, data.PageSize)

	var response []dto.PastMessagesResponse
	for _, v := range messages {

		response = append(response, dto.PastMessagesResponse{
			ID:       v.ID,
			Message:  v.Message,
			SenderID: v.SenderID,
		})

	}

	fmt.Println(err)

	return nil, nil

	// m.DAO.Read()
}

func (m messageUsecase) Update(message dto.UpdateMessageRequest) {
	// m.DAO.Update()
}

func (m messageUsecase) Delete(message dto.DeleteMessageRequest) {
	// m.DAO.Delete()
}
