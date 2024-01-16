package usecase

import (
	"chat/internal/core/dto"
	"chat/internal/core/port/driven/dao"

	"chat/utilities/auth"

	"github.com/google/uuid"
)

type ConversationUsecase interface {
	StartConversation(dto.StartConversation) error
	// StartConversation(dto.StartConversation)
}

type conversationUsecase struct {
	ConversationDAO dao.ConversationPortDAO
	UserDAO         dao.UserPortDAO
}

func NewConversationUsecase(ConversationDAO dao.ConversationPortDAO, UserDAO dao.UserPortDAO) ConversationUsecase {
	return &conversationUsecase{
		ConversationDAO: ConversationDAO,
		UserDAO:         UserDAO,
	}
}

func (c conversationUsecase) StartConversation(conversation dto.StartConversation) error {

	// Check user id
	userId, err := auth.IsUserToken(conversation.AccessToken)

	if err != nil {
		return err
	}

	// query and find a conversation id
	chatID, err := uuid.Parse(conversation.Channel)
	if err != nil {
		return err
	}

	err = c.ConversationDAO.CheckCoversationandUserID(userId, chatID)
	if err != nil {
		return err
	}

	// find a user id in converation id

	return nil
}

func (c conversationUsecase) Conversation(token string) ([]dto.ConversationListResponse, error) {

	userID, err := auth.IsUserToken(token)

	if err != nil {
		return nil, nil
	}

	conversations, err := c.ConversationDAO.UserConversations(userID)

	var ids []uuid.UUID

	for _, v := range conversations {
		if v.ID != userID {
			ids = append(ids, v.ID)
		}
	}

	users, err := c.UserDAO.UserList(ids)

	var conversationList []dto.ConversationListResponse

	for _, v := range users {

		conversationList = append(conversationList, dto.ConversationListResponse{
			// ID             :
			ParticipantFirstName: v.FirstName,
			ParticipantLastName:  v.LastName,
			ParticipantImage:     v.Avatar,
		})

	}

	return nil, nil
}
