package profiles

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio-api/infrastructure/db"
	"portfolio-api/interfaces/api"
	"portfolio-api/interfaces/repositories"
	"portfolio-api/usecases/logging"
	"portfolio-api/usecases/profiles"
	"strconv"
)

type ProfileController struct {
	Usecase *profiles.ProfileUsecase
}

func NewProfileController(db *db.DB, logging logging.Logging) *ProfileController {
	return &ProfileController{
		Usecase: &profiles.ProfileUsecase{
			ProfileRepository: &repositories.ProfileRepository{},
			DB:                db,
			Logging:           logging,
		},
	}
}

func (controller *ProfileController) GetUserProfile(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	profile, uerr := controller.Usecase.GetUserProfile(userId)
	if uerr != nil {
		c.JSON(api.GetErrorResponse(uerr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profile": profile,
	})
}
