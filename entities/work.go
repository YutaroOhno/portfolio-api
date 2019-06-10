package entities

type Work struct {
	ID    	    int    `gorm:"primary_key" json:"id"`
	UserId	    int    `gorm:"not_null" json:"user_id"`
	Title       string `gorm:"not_null" json:"title"`
	SubTitle    string `gorm:"not_null" json:"sub_title"`
	Description string `json:"description"`
	URL 		string `json:"url"`
}