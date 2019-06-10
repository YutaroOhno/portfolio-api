package repositories

import (
	"portfolio-api/entities"
	"github.com/jinzhu/gorm"
)

type SkillRepository struct {
}

func (repo *SkillRepository) Find(userID int, db *gorm.DB) (entities.User, error) {
	var user entities.User
	db.LogMode(true)
	db.Preload("Skills").First(&user)
	// 以下のような書き方もできるが、levelが欲しいので上記のようにしている。
	// db.First(&user, userID).Related(&user.Skills, "Skills")
	return user, nil
}
