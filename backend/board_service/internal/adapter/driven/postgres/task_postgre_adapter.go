package postgres

import (
	"board/internal/core/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
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

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepo {
	return &taskRepo{
		db: db,
	}
}

func (t taskRepo) GetById(id uuid.UUID) (entity.Task, error) {
	var task Task
	query := t.db.Where("id = ?", id).First(&task)

	if query.Error != nil {
		return entity.Task{}, query.Error
	}

	return entity.Task{}, nil
}

func (t taskRepo) Index(page int, pageSize int) ([]entity.Task, error) {

	var taskList []entity.Task

	var tasks []Task

	offset := (page - 1) * pageSize

	// Fetch all records from the 'products' table
	if err := t.db.Offset(offset).Limit(pageSize).Find(&tasks).Error; err != nil {
		panic("Failed to retrieve products!")
	}

	for _, v := range tasks {
		taskList = append(taskList, entity.Task{
			ID: v.ID,
		})
	}

	return taskList, nil

}

func (t taskRepo) Create(task entity.Task) error {
	newTask := Task{}

	// Create the user in the database
	query := t.db.Create(&newTask)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (t taskRepo) Update(data entity.Task) error {
	task := Task{}
	query := t.db.First(&task, "id = ?", task.ID)
	// TODO:
	if query.Error != nil {
		return query.Error
	}

	t.db.Model(&task).Updates(task)
	return nil
}

func (t taskRepo) Delete(id uuid.UUID) error {
	var task Task
	query := t.db.First(&task, id)

	if query.Error != nil {
		return query.Error
	}

	query = t.db.Delete(&task)

	if query.Error != nil {
		return query.Error
	}

	return nil
}
