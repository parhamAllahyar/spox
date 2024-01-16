package driver

import (
	"board/internal/core/entity"
)

type TaskPortDriver interface {
	Get(id uint) (entity.Task, error)
	Index() ([]entity.Task, error)
	Create(entity.Task) (entity.Task, error)
	Update(entity.Task) (entity.Task, error)
}
