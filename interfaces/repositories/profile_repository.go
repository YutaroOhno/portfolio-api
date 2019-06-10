package repositories

import (
	"portfolio-api/entities"
	"github.com/jinzhu/gorm"
)

type ProfileRepository struct {
}

func (repo *ProfileRepository) Find(userID int, db *gorm.DB) (entities.Profile, error) {
	var profile entities.Profile
	db.First(&profile, "user_id = ?", userID)
	return profile, nil
}
