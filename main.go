package main

import (
	"portfolio-api/infrastructure/db"
	"portfolio-api/infrastructure/server"
	"portfolio-api/infrastructure/logging"
	ulogging "portfolio-api/usecases/logging"
)

func main() {
	var newDB db.Database
	newDB = db.NewMysql()
	db := newDB.Open()
	defer db.Close()


	// 環境によってlogginを分ける想定。現状logrusのみ。
	var ulogging ulogging.Logging
	ulogging = logging.NewLogrusLogging()

	server.Run(db, ulogging)
}