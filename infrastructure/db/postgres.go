package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/lib/pq"
	"os"
	"fmt"
)

type Postgres struct {
	gormDB *gorm.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (postgres *Postgres) Open() *DB {
	// 本番環境（heroku）を考慮
	// var (
	// 	connection string
	// 	err        error
	// )
	// if os.Getenv("ENV") == "prod" {
	// 	url := os.Getenv("DATABASE_URL")
	// 	connection, err = pq.ParseURL(url)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	connection += " sslmode=require"
	// } else {
	// 	connection = os.Getenv("DATABASE_URL")
	// }

	connections := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", os.Getenv("MASTER_USER"), os.Getenv("CLOUDSQL_DBNAME"), os.Getenv("MASTER_PASSWORD")) //Build connection string
	db, err := gorm.Open("postgres", connections)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	return &DB{GormDB: db}
}
