package postgres

import (
	"time"
	"workspace/internal/core/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Workspace struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title      string
	CreatorID  uuid.UUID
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Deleted_at gorm.DeletedAt
	Roles      []*Role `gorm:"many2many:role_workspaces;"`
}

type RoleWorkspace struct {
	UserID      uuid.UUID
	RoleID      uuid.UUID
	WorkspaceID uuid.UUID
}

type workspaceRepo struct {
	db *gorm.DB
}

func NewWorkspaceRepo(db *gorm.DB) *workspaceRepo {
	return &workspaceRepo{
		db: db,
	}
}

// Create workspace
func (w workspaceRepo) Create(workspace entity.Workspace) (entity.Workspace, error) {

	newWorkspace := Workspace{
		Title:     workspace.Title,
		CreatorID: workspace.CreatorId,
		// Other field assignments
	}

	query := w.db.Create(&newWorkspace).Find(&newWorkspace)

	if query.Error != nil {
		return entity.Workspace{}, query.Error
	}

	return entity.Workspace{
		ID:        newWorkspace.ID,
		Title:     newWorkspace.Title,
		CreatorId: newWorkspace.CreatorID,
	}, nil

}

// Workspace By ID

// Workspace list by user id (user is a member of workspace)

// Update Workspace
func (w workspaceRepo) Update(workspace entity.Workspace) error {
	data := Workspace{}
	query := w.db.Table("workspaces").Where("id = ?", workspace.ID).Updates(data)
	if query.Error != nil {
		return nil
	}
	return nil
}

// Delete Workspace
func (w workspaceRepo) Delete(id uuid.UUID) error {

	var workspace Workspace
	query := w.db.First(&workspace, id)

	if query.Error != nil {
		return query.Error
	}

	query = w.db.Delete(&workspace)

	if query.Error != nil {
		return query.Error
	}

	return nil

}

func (w workspaceRepo) Get(id uuid.UUID) (entity.Workspace, error) {

	return entity.Workspace{}, nil
}

func (w workspaceRepo) Index(int, int) ([]entity.Workspace, error) {
	return nil, nil
}

func (w workspaceRepo) AssignUserRole(userID uuid.UUID, workspaceID uuid.UUID, roleID uuid.UUID) error {

	newUserRolesWorkspace := RoleWorkspace{
		UserID:      userID,
		RoleID:      roleID,
		WorkspaceID: workspaceID,
	}

	return w.db.Create(&newUserRolesWorkspace).Error

}

func (w workspaceRepo) UserWorkspace(uuid.UUID) ([]entity.Workspace, error) {
	return nil, nil
}
