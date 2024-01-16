package driver

import (
	"board/internal/core/entity"
)

type BoardPortDriver interface {
	Get(id uint) (entity.Board, error)
	Index() []entity.Board
}
