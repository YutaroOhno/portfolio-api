package server

import (
	"portfolio-api/infrastructure/db"
	"portfolio-api/interfaces/api/profiles"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"portfolio-api/usecases/logging"
)

func Run(db *db.DB, logging logging.Logging) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//とりあえずフロントからのアクセスを許可したいので、記述。
		AllowOrigins: []string{"http://127.0.0.1:8080"},
		AllowMethods: []string{"GET", "PUT", "PATCH", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin"},
	}))

	// ユーザーのプロフィール情報を取得
	profileController := profiles.NewProfileController(db,logging)
	router.GET("/users/:user_id/profile", profileController.GetUserProfile)

	// ユーザーの作品情報を取得

	// ユーザーのスキル情報を取得

	router.Run()
}