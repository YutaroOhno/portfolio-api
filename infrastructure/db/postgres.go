package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type Postgre struct {
	gormDB *gorm.DB
}

func NewPostgres() *Postgre {
	return &Postgre{}
}

func (postgres *Postgre) Open() *DB {
	CONNECT := "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"
	db, err := gorm.Open("postgres", CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return &DB{GormDB: db}
}