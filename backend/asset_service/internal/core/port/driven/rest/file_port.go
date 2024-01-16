package rest

type ConversationDAO interface {
	IsUserInConversation(string) (bool, error)
}
