package server

import (
	"portfolio-api/infrastructure/db"
	"portfolio-api/interfaces/api/profiles"
	"portfolio-api/interfaces/api/skills"
	"portfolio-api/interfaces/api/works"
	"portfolio-api/interfaces/api/contacts"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"portfolio-api/usecases/logging"
	"portfolio-api/usecases/mailer"
	"os"
)

func Run(db *db.DB, logging logging.Logging, mailer mailer.Mailer) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//とりあえずフロントからのアクセスを許可したいので、記述。
		AllowOrigins: []string{os.Getenv("CLIENT_ORIGIN")},
		AllowMethods: []string{"GET", "PUT", "PATCH", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin"},
	}))

	// ユーザーのプロフィール情報を取得
	profileController := profiles.NewProfileController(db,logging)
	router.GET("/users/:user_id/profile", profileController.GetUserProfile)

	// ユーザーのスキル情報を取得
	skillController := skills.NewUserSkillController(db,logging)
	router.GET("/users/:user_id/skills", skillController.GetUserSKills)

	// ユーザーの作品情報を取得
	workController := works.NewWorkController(db,logging)
	router.GET("/users/:user_id/works", workController.GetUserWorks)

	// コンタクト
	contactController := contacts.NewContactController(mailer, logging)
	router.POST("/users/:user_id/contact", contactController.SendMail)
	
	router.Run()
}