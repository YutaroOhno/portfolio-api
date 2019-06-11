package entities

type Skill struct {
	ID    	 int    `gorm:"primary_key" json:"id"`
	Name 	 string `gorm:"not null" json:"name"`
	//levelは中間テーブルにある
	Level   int   `json:"level"`
}
