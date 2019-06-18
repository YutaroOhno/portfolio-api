package profiles

import (
	"portfolio-api/entities"
	"portfolio-api/infrastructure/db"
	"portfolio-api/usecases/ports"
	"portfolio-api/usecases/repositories"
	"portfolio-api/usecases"
	"portfolio-api/usecases/logging"
	"os"
)

type ProfileUsecase struct {
	ProfileRepository repositories.ProfileRepository
	DB             *db.DB
	Logging 		logging.Logging
}

func (usecase *ProfileUsecase) GetUserProfile(userId int) (*ports.ProfileOutputPort, *usecases.UError) {
	profile, err := usecase.ProfileRepository.Find(userId, usecase.DB.GormDB)
	profile.Avatar = os.Getenv("BUCKET_PATH") + profile.Avatar
	if uerr := usecases.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return nil, uerr
	}

	return createOutputPort(profile), nil
}

func createOutputPort(profile entities.Profile) *ports.ProfileOutputPort {
	return &ports.ProfileOutputPort{
		UserID:    	   profile.UserID,
		Nickname: 	   profile.Nickname,
		Introduction:  profile.Introduction,
		Avatar: 	   profile.Avatar,
	}
}