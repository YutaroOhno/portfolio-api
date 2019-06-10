package repositories

import (
	"portfolio-api/entities"
	"github.com/jinzhu/gorm"
)

type SkillRepository interface {
	Find(int, *gorm.DB) (entities.User, error)
}
