package cassandra

import (
	"chat/internal/core/entity"
	"chat/internal/core/port/driven/dao"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type conversationDAO struct {
	ConversationDAO dao.ConversationPortDAO
	db              *gocql.Session
}

func NewConversationDAO(db *gocql.Session) dao.ConversationPortDAO {
	return &conversationDAO{
		db: db,
	}
}

func (c conversationDAO) Create(entity.Conversation) error {

	// TODO: Change
	insertQuery := "INSERT INTO mytable (id, name, age) VALUES (?, ?, ?)"

	err := c.db.Query(insertQuery, 12, "John Doe", 30).Exec()

	return err

}
func (c conversationDAO) Delete(i uuid.UUID) error {

	// Delete query
	deleteQuery := "DELETE FROM mytable WHERE id = ?"

	// Define the ID of the record you want to delete
	id, err := gocql.ParseUUID("your-id-string-here") // Replace with the ID you want to delete

	if err != nil {
	}

	// TODO:Check is exist

	// Execute delete query
	err = c.db.Query(deleteQuery, id).Exec()

	return err

}

func (c conversationDAO) Index(pageSize int64, Perpage int64) ([]entity.Conversation, error) {
	return nil, nil
}

func (c conversationDAO) UserConversations(id uuid.UUID) ([]entity.Conversation, error) {
	return nil, nil
}

func (c conversationDAO) Update(entity.Conversation) error {

	// Update query
	updateQuery := "UPDATE mytable SET name = ?, age = ? WHERE id = ?"

	// Define the ID of the record you want to update
	id, err := gocql.ParseUUID("your-id-string-here")

	// TODO:Check is exist

	// Execute update query
	err = c.db.Query(updateQuery, "Updated Name", 35, id).Exec()

	return err

}

func (c conversationDAO) GetByID(uuid.UUID) (entity.Conversation, error) {

	return entity.Conversation{}, nil
}

func (c conversationDAO) CheckCoversationandUserID(uuid.UUID, uuid.UUID) error {

	return nil

}
