package repositories

import (
	"portfolio-api/entities"
	"github.com/jinzhu/gorm"
)

type WorkRepository struct {
}

func (repo *WorkRepository) Find(userId int, db *gorm.DB) ([]entities.Work, error) {
	var works []entities.Work
	db.Where("user_id = ?", userId).Find(&works)
	return works, nil
}
