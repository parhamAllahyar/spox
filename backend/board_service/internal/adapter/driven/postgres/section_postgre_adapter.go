package postgres

import (
	"board/internal/core/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Section struct {
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

type sectionRepo struct {
	db *gorm.DB
}

func NewSectionRepo(db *gorm.DB) *sectionRepo {
	return &sectionRepo{
		db: db,
	}
}

func (s sectionRepo) GetById(id uuid.UUID) (entity.Section, error) {

	var section Section
	query := s.db.Where("id = ?", id).First(&section)

	if query.Error != nil {
		return entity.Section{}, query.Error
	}

	return entity.Section{}, nil

}

func (s sectionRepo) Index(page int, pageSize int) ([]entity.Section, error) {

	var sectionList []entity.Section

	var sections []Section

	offset := (page - 1) * pageSize

	// Fetch all records from the 'products' table
	if err := s.db.Offset(offset).Limit(pageSize).Find(&sections).Error; err != nil {
		panic("Failed to retrieve products!")
	}

	for _, v := range sections {
		sectionList = append(sectionList, entity.Section{
			ID: v.ID,
		})
	}

	return sectionList, nil

}

func (s sectionRepo) Create(section entity.Section) error {
	newSection := Section{}

	// Create the user in the database
	query := s.db.Create(&newSection)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (s sectionRepo) Update(data entity.Section) error {

	subTask := SubTask{}
	query := s.db.First(&subTask, "id = ?", subTask.ID)
	// TODO:
	if query.Error != nil {
		return query.Error
	}

	s.db.Model(&subTask).Updates(subTask)
	return nil

}

func (s sectionRepo) Delete(id uuid.UUID) error {

	var section Section
	query := s.db.First(&section, id)

	if query.Error != nil {
		return query.Error
	}

	query = s.db.Delete(&section)

	if query.Error != nil {
		return query.Error
	}

	return nil

}
