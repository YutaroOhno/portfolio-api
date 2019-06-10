package repositories

import (
	"portfolio-api/entities"
	"github.com/jinzhu/gorm"
)

type WorkRepository interface {
	Find(int, *gorm.DB) ([]entities.Work, error)
}
