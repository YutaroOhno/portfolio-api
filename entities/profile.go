package entities

type Profile struct {
	UserID       int    `gorm:"primary_key"`
	Nickname     string `gorm:"not null"`
	Introduction string `gorm:"not null"`
	Avatar       string
    Link 		 []Link   `gorm:"ForeignKey:UserId"`
}
