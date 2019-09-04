package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
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

	var (
		connection string
		err        error
	)

	if os.Getenv("ENV") == "prod" {
		url := os.Getenv("DATABASE_URL")
		connection, err = pq.ParseURL(url)
		if err != nil {
			panic(err.Error())
		}
		connection += " sslmode=require"
	} else {
		connection = os.Getenv("DATABASE_URL")
	}

	db, err := gorm.Open("postgres", connection)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	return &DB{GormDB: db}
}
