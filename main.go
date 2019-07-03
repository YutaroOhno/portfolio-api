package main

import (
	"portfolio-api/infrastructure/db"
	"portfolio-api/infrastructure/server"
	"portfolio-api/infrastructure/logging"
	"portfolio-api/infrastructure/mailer"
	ulogging "portfolio-api/usecases/logging"
	umailer "portfolio-api/usecases/mailer"
)

func main() {
	var newDB db.Database
	newDB = db.NewPostgres()
	db := newDB.Open()
	defer db.Close()

	// 環境によってlogginを分ける想定。現状logrusのみ。
	var ulogging ulogging.Logging
	ulogging = logging.NewLogrusLogging()

	var umailer umailer.Mailer
	umailer = mailer.NewMailer()
	server.Run(db, ulogging, umailer)
}