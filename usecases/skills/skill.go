package skills

import (
	"portfolio-api/entities"
	"portfolio-api/infrastructure/db"
	"portfolio-api/usecases/repositories"
	"portfolio-api/usecases/logging"
	"portfolio-api/usecases/ports"
	"portfolio-api/usecases"
	"fmt"
)

type SkillUsecase struct {
	SkillRepository repositories.SkillRepository
	DB             *db.DB
	Logging 		logging.Logging
}

func (usecase *SkillUsecase) GetUserSkills(userId int) (*ports.UserSkillsOutputPort, *usecases.UError) {
	userWithSkills, err := usecase.SkillRepository.Find(userId, usecase.DB.GormDB)
	fmt.Println(userWithSkills)
	if uerr := usecases.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
	}

	return createOutputPort(userWithSkills), nil
}

func createOutputPort(userWithSkills entities.User) *ports.UserSkillsOutputPort {
	return &ports.UserSkillsOutputPort{
		UserID:    	userWithSkills.ID,
		Skills:     userWithSkills.Skills,
	}
}