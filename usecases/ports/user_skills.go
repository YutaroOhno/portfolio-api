package ports

import (
	"portfolio-api/entities"
)

type UserSkillsOutputPort struct {
	UserID int              `json:user_id`
	Skills []entities.Skill `gorm:"many2many:user_skills;"`
}
