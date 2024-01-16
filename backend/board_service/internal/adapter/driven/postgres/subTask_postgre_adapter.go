package postgres

import (
	"board/internal/core/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubTask struct {
	gorm.Model
	ID uuid.UUID `gorm:"primary_key;auto_increment"`
	// OrderType  string  `gorm:"size:255;not null"`
	// Price      float32 `gorm:"size:255;not null"`
	// Quantity   float32 `gorm:"size:255;not null"`
	// CurrencyID uint64
	// Currency   Currency  `gorm:"foreignKey:CurrencyID"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Deleted_at gorm.DeletedAt
}

type subTaskRepo struct {
	db *gorm.DB
}

func NewSubTaskdRepo(db *gorm.DB) *subTaskRepo {
	return &subTaskRepo{
		db: db,
	}
}

func (s subTaskRepo) GetById(id uuid.UUID) (entity.SubTask, error) {
	var subTask SubTask
	query := s.db.Where("id = ?", id).First(&subTask)

	if query.Error != nil {
		return entity.SubTask{}, query.Error
	}

	return entity.SubTask{}, nil
}

func (s subTaskRepo) Index(page int, pageSize int) ([]entity.SubTask, error) {

	var subtaskList []entity.SubTask

	var subtasks []Tag

	offset := (page - 1) * pageSize

	// Fetch all records from the 'products' table
	if err := s.db.Offset(offset).Limit(pageSize).Find(&subtasks).Error; err != nil {
		panic("Failed to retrieve products!")
	}

	for _, v := range subtasks {
		subtaskList = append(subtaskList, entity.SubTask{
			ID: v.ID,
		})
	}

	return subtaskList, nil

}

func (s subTaskRepo) Create(subTask entity.SubTask) error {

	newSubtask := SubTask{}

	// Create the user in the database
	query := s.db.Create(&newSubtask)
	if query.Error != nil {
		return query.Error
	}
	return nil

}

func (s subTaskRepo) Update(data entity.SubTask) error {
	subTask := SubTask{}
	query := s.db.First(&subTask, "id = ?", subTask.ID)
	// TODO:
	if query.Error != nil {
		return query.Error
	}

	s.db.Model(&subTask).Updates(subTask)
	return nil
}

func (s subTaskRepo) Delete(id uuid.UUID) error {

	var subtask SubTask
	query := s.db.First(&subtask, id)

	if query.Error != nil {
		return query.Error
	}

	query = s.db.Delete(&subtask)

	if query.Error != nil {
		return query.Error
	}

	return nil

}
