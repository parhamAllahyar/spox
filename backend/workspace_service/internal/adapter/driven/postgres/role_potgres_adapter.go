package postgres

import (
	"time"
	"workspace/internal/core/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primary_key;auto_increment"`
	Title      string
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Deleted_at gorm.DeletedAt
	Workspaces []*Workspace `gorm:"many2many:role_workspaces;"`
}

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *roleRepo {
	return &roleRepo{
		db: db,
	}
}

// Create workspace
func (r *roleRepo) Create(role entity.Role) error {
	result := r.db.Table("roles").Create(&Role{})
	if result.Error != nil {
		// TODO: Log and custome Error
		return result.Error
	}
	return nil
}

func (r *roleRepo) Update(role entity.Role) error {
	data := Role{}
	query := r.db.Table("roles").Where("id = ?", role.ID).Updates(data)
	if query.Error != nil {
		return nil
	}
	return nil
}

func (r *roleRepo) Delete(id uuid.UUID) error {

	var role Role
	query := r.db.First(&role, id)

	if query.Error != nil {
		return query.Error
	}

	query = r.db.Delete(&role)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r roleRepo) Get(title string) (entity.Role, error) {

	var role entity.Role
	if err := r.db.Where("title = ?", title).First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}

func (r roleRepo) Index(int, int) ([]entity.Role, error) {
	return nil, nil
}
