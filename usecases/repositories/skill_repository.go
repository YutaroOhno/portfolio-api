package repositories

import (
	"github.com/jinzhu/gorm"
	"portfolio-api/entities"
)

type SkillRepository interface {
	Find(int, *gorm.DB) (entities.User, error)
}
