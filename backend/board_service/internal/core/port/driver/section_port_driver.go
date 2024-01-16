package driver

import (
	"board/internal/core/entity"
)

type SectionPortDriver interface {
	Get(id uint) (entity.Section, error)
	Index() []entity.Section
}
