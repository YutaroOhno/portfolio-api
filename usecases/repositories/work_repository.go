package repositories

import (
	"github.com/jinzhu/gorm"
	"portfolio-api/entities"
)

type WorkRepository interface {
	Find(int, *gorm.DB) ([]entities.Work, error)
}
