package postgres

import (
	"asset/internal/core/entity"
	"asset/internal/core/port/driven/dao"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID        uint64    `gorm:"primary_key;auto_increment"`
	Name      string    `gorm:"size:255;not null"`
	Symbol    string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt
}

type fileDAO struct {
	FileDAO dao.FileDAO
	db      *gorm.DB
}

func NewFileDAO(db *gorm.DB) dao.FileDAO {
	return &fileDAO{
		db: db,
	}
}

// Create implements dao.FileDAO.
func (f fileDAO) Create(entity.File) error {

	newFile := File{}

	query := f.db.Create(&newFile)
	if query.Error != nil {
		return query.Error
	}
	return nil

}

// Delete implements dao.FileDAO.
func (f fileDAO) Delete(id uuid.UUID) error {
	var file File
	query := f.db.First(&file, id)

	if query.Error != nil {
		return query.Error
	}

	query = f.db.Delete(&file)

	if query.Error != nil {
		return query.Error
	}

	return nil

}

// List implements dao.FileDAO.
func (f fileDAO) List(uuid.UUID) ([]entity.File, error) {
	panic("unimplemented")
}

func (f fileDAO) GetByID(uuid.UUID) (entity.File, error) {

	return entity.File{}, nil
}
