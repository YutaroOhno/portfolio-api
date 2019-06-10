package entities

type User struct {
	ID    	 int      `gorm:"primary_key"`
	Skills   []Skill  `gorm:"many2many:user_skills;"`	
}
