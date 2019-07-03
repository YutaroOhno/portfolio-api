package repositories

import (
	"portfolio-api/entities"
	"github.com/jinzhu/gorm"
)

type SkillRepository struct {
}

func (repo *SkillRepository) Find(userID int, db *gorm.DB) (entities.User, error) {
	var user entities.User
	user.ID = userID
	db.Preload("Skills").First(&user)

	return user, nil
}
