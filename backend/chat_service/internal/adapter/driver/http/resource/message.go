package resource

type SendTextmessageRequest struct {
	Message string `json:message`
}

type LastMessageListResponse struct {
	Message  string
	SenderID string
	ID       string
}
