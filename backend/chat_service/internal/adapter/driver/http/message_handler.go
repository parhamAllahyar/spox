package http

import (
	"chat/internal/adapter/driver/http/resource"
	"chat/internal/core/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/olahol/melody"
)

func (srv server) SendTextMessageHandler(s *melody.Session, msg []byte) {

	var message resource.SendTextmessageRequest
	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	err = srv.messageUsecase.CreateTextMessage(dto.CreateTextMessageRequest{
		Message: message.Message,
		Token:   s.Request.Header.Get("Authorization"),
		ChatID:  s.Request.URL.Query().Get("channel"),
	})

	srv.m.Broadcast(msg)

}

func (srv server) GetPastMessagesHandler(w http.ResponseWriter, r *http.Request) {

	id, _ := uuid.Parse(r.URL.Query().Get("channel"))
	page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	PageSize, _ := strconv.ParseInt(r.URL.Query().Get("page_size"), 10, 64)

	messages, err := srv.messageUsecase.PastMessages(dto.PastMessageRequest{
		ChatID:      id,
		AccessToken: r.Header.Get("Authorization"),
		Page:        page,
		PageSize:    PageSize,
	})

	var messageResponse []resource.LastMessageListResponse

	for _, v := range messages {
		messageResponse = append(messageResponse, resource.LastMessageListResponse{
			Message:  v.Message,
			SenderID: v.SenderID.String(),
			ID:       v.ID.String(),
		})
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(messageResponse)
	if err != nil {
		// TODO:
	}

}

func (srv server) SendFileHandler(w http.ResponseWriter, r *http.Request) {

}
