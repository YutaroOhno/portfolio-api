package server

import (
	"portfolio-api/infrastructure/db"
	"portfolio-api/interfaces/api/profiles"
	"portfolio-api/interfaces/api/skills"
	"portfolio-api/interfaces/api/works"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"portfolio-api/usecases/logging"
)

func Run(db *db.DB, logging logging.Logging) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//とりあえずフロントからのアクセスを許可したいので、記述。
		AllowOrigins: []string{"http://localhost:8080"},
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

	router.Run()
}