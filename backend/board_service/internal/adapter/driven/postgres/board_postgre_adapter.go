package postgres

import (
	"board/internal/core/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	WorkspaceID uuid.UUID
	Title       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type boardRepo struct {
	db *gorm.DB
}

func NewBoardRepo(db *gorm.DB) *boardRepo {
	return &boardRepo{
		db: db,
	}
}

func (b boardRepo) Create(board entity.Board) error {
	newUser := Board{}

	// Create the user in the database
	query := b.db.Create(&newUser)
	if query.Error != nil {
		return query.Error
	}
	return nil

}

func (b boardRepo) Update(data entity.Board) error {

	board := Board{}
	query := b.db.First(&board, "id = ?", board.ID)
	// TODO:
	if query.Error != nil {
		return query.Error
	}

	b.db.Model(&board).Updates(board)
	return nil

}

func (b boardRepo) GetById(id uuid.UUID) (entity.Board, error) {

	var board Board
	query := b.db.Where("id = ?", id).First(&board)

	if query.Error != nil {
		return entity.Board{}, query.Error
	}

	return entity.Board{}, nil

}
func (b boardRepo) Index(page int, pageSize int) ([]entity.Board, error) {
	var boardList []entity.Board

	var boards []Tag

	offset := (page - 1) * pageSize

	// Fetch all records from the 'products' table
	if err := b.db.Offset(offset).Limit(pageSize).Find(&boards).Error; err != nil {
		panic("Failed to retrieve products!")
	}

	for _, v := range boards {
		boardList = append(boardList, entity.Board{
			ID: v.ID,
		})
	}

	return boardList, nil

}

func (b boardRepo) Delete(id uuid.UUID) error {

	var board Board
	query := b.db.First(&board, id)

	if query.Error != nil {
		return query.Error
	}

	query = b.db.Delete(&board)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (b boardRepo) BordWithDetails(uuid.UUID) (entity.Board, error) {

	return entity.Board{}, nil
}

func (b boardRepo) WorkspaceBoards(uuid.UUID) ([]entity.Board, error) {

	return nil, nil
}
