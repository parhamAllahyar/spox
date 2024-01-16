package cassandra

import (
	"chat/internal/core/entity"
	"chat/internal/core/port/driven/dao"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type messageDAO struct {
	MessageDAO dao.MessagePortDAO
	db         *gocql.Session
}

func NewMessageDAO(db *gocql.Session) dao.MessagePortDAO {
	return &messageDAO{
		db: db,
	}
}

func (m messageDAO) Create(entity.Message) error {

	// TODO: Change
	insertQuery := "INSERT INTO mytable (id, name, age) VALUES (?, ?, ?)"

	err := m.db.Query(insertQuery, 12, "John Doe", 30).Exec()

	return err

}
func (m messageDAO) Delete(uuid.UUID) error {

	// Delete query
	deleteQuery := "DELETE FROM mytable WHERE id = ?"

	// Define the ID of the record you want to delete
	id, err := gocql.ParseUUID("your-id-string-here") // Replace with the ID you want to delete

	if err != nil {
	}

	// TODO:Check is exist

	// Execute delete query
	err = m.db.Query(deleteQuery, id).Exec()

	return err

}

func (m messageDAO) Index(chatId uuid.UUID, pageSize int64, Perpage int64) ([]entity.Message, error) {
	return nil, nil
}

func (m messageDAO) Update(entity.Message) error {

	// Update query
	updateQuery := "UPDATE mytable SET name = ?, age = ? WHERE id = ?"

	// Define the ID of the record you want to update
	id, err := gocql.ParseUUID("your-id-string-here")

	// TODO:Check is exist

	// Execute update query
	err = m.db.Query(updateQuery, "Updated Name", 35, id).Exec()

	return err

}
