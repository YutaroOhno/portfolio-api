package repositories

import (
	"portfolio-api/entities"
	"github.com/jinzhu/gorm"
)

type ProfileRepository interface {
	Find(int, *gorm.DB) (entities.Profile, error)
}
