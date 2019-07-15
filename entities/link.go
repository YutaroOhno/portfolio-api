package entities

type Link struct {
	ID        int    `gorm:"primary_key" json:"id"`
	UserId    string `gorm:"not null" json:"user_id"`
	Title	  string `json:"title"`
	URL 	  string `json:"url"`
}
