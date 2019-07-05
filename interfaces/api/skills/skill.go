package skills

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portfolio-api/infrastructure/db"
	"portfolio-api/interfaces/api"
	"portfolio-api/interfaces/repositories"
	"portfolio-api/usecases/logging"
	"portfolio-api/usecases/skills"
	"strconv"
)

type SkillController struct {
	Usecase *skills.SkillUsecase
}

func NewUserSkillController(db *db.DB, logging logging.Logging) *SkillController {
	return &SkillController{
		Usecase: &skills.SkillUsecase{
			SkillRepository: &repositories.SkillRepository{},
			DB:              db,
			Logging:         logging,
		},
	}
}

func (controller *SkillController) GetUserSKills(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, uerr := controller.Usecase.GetUserSkills(userId)
	if err != nil {
		c.JSON(api.GetErrorResponse(uerr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": user.UserID,
		"skills":  user.Skills,
	})
}
