package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"os"
	"fmt"
)

type Postgre struct {
	gormDB *gorm.DB
}

func NewPostgres() *Postgre {
	return &Postgre{}
}

func (postgres *Postgre) Open() *DB {
	// herokuの場合
	connection := ""
	if os.Getenv("ENV") == "prod" {
		url := os.Getenv("DATABASE_URL")
		connection, err := pq.ParseURL(url)
		if err != nil {
			fmt.Println("エラー")
		}
		connection += " sslmode=require"
	} else  {
		connection = os.Getenv("DATABASE_URL")
	}

	db, err := gorm.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return &DB{GormDB: db}
}