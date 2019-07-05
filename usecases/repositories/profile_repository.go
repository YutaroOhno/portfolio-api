package repositories

import (
	"github.com/jinzhu/gorm"
	"portfolio-api/entities"
)

type ProfileRepository interface {
	Find(int, *gorm.DB) (entities.Profile, error)
}
