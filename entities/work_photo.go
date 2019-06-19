package entities

type WorkPhoto struct {
	WorkId int    	 `gorm:"primary_key" json:"id"`
	PhotoPath string `json:"photo_path"`
}