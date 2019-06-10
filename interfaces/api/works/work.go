package works

import (
	"portfolio-api/infrastructure/db"
	"portfolio-api/usecases/logging"
	"portfolio-api/usecases/works"
	"portfolio-api/interfaces/repositories"
	"github.com/gin-gonic/gin"
	"portfolio-api/interfaces/api"
	"net/http"
	"strconv"
)

type WorkController struct {
	Usecase *works.WorkUsecase
}

func NewWorkController(db *db.DB, logging logging.Logging) *WorkController {
	return &WorkController{
		Usecase: &works.WorkUsecase {
			WorkRepository: &repositories.WorkRepository{},
			DB:             db,
			Logging: logging,
		},
	}
}

func (controller *WorkController) GetUserWorks(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))	
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	works, uerr := controller.Usecase.GetUserWorks(userId)
	if err != nil {
		c.JSON(api.GetErrorResponse(uerr))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		// これは構造体見直したい
		"works": works.Works,
	})
}
