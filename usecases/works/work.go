package works

import (
	"portfolio-api/entities"
	"portfolio-api/infrastructure/db"
	"portfolio-api/usecases/repositories"
	"portfolio-api/usecases/logging"
	"portfolio-api/usecases/ports"
	"portfolio-api/usecases"
)

type WorkUsecase struct {
	WorkRepository repositories.WorkRepository
	DB             *db.DB
	Logging 		logging.Logging
}

func (usecase  *WorkUsecase) GetUserWorks(userId int) (*ports.WorksOutputPort, *usecases.UError) {
	works, err := usecase.WorkRepository.Find(userId, usecase.DB.GormDB)
	if uerr := usecases.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
	}

	return createOutputPort(works), nil
}

func createOutputPort(works []entities.Work) *ports.WorksOutputPort {
	return &ports.WorksOutputPort{
		Works: works,
	}
}