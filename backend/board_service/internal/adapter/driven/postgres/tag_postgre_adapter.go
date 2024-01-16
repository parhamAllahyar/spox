package postgres

import (
	"board/internal/core/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
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

type tagRepo struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) tagRepo {
	return tagRepo{
		db: db,
	}
}

func (t tagRepo) GetById(id uuid.UUID) (entity.Tag, error) {
	var tag Tag
	query := t.db.Where("id = ?", id).First(&tag)

	if query.Error != nil {
		return entity.Tag{}, query.Error
	}

	return entity.Tag{}, nil
}

func (t tagRepo) Index(page int, pageSize int) ([]entity.Tag, error) {

	var tagList []entity.Tag

	var tags []Tag

	offset := (page - 1) * pageSize

	// Fetch all records from the 'products' table
	if err := t.db.Offset(offset).Limit(pageSize).Find(&tags).Error; err != nil {
		panic("Failed to retrieve products!")
	}

	for _, v := range tags {
		tagList = append(tagList, entity.Tag{
			ID: v.ID,
		})
	}

	return tagList, nil

}

func (t tagRepo) Create(tag entity.Tag) error {
	newTag := Tag{}

	// Create the user in the database
	query := t.db.Create(&newTag)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (t tagRepo) Update(data entity.Tag) error {

	tag := Tag{}
	query := t.db.First(&tag, "id = ?", tag.ID)
	// TODO:
	if query.Error != nil {
		return query.Error
	}

	t.db.Model(&tag).Updates(tag)
	return nil

}

func (t tagRepo) Delete(id uuid.UUID) error {

	var tag Tag
	query := t.db.First(&tag, id)

	if query.Error != nil {
		return query.Error
	}

	query = t.db.Delete(&tag)

	if query.Error != nil {
		return query.Error
	}

	return nil

}
