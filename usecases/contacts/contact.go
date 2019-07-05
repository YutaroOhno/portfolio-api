package contacts

import (
	"portfolio-api/usecases/mailer"
	"portfolio-api/usecases/logging"
	"portfolio-api/usecases/ports"
	"portfolio-api/usecases"
)

type ContactUsecase struct {
	Mailer  mailer.Mailer
	Logging logging.Logging
}

func (usecase *ContactUsecase) SendMail(input *ports.ContactInputPort) *usecases.UError {
	err := usecase.Mailer.Send(input.Name, input.MailAddress, input.Content)
	if uerr := usecases.GetUErrorByError(err); uerr != nil {
		usecase.Logging.Error(uerr)
		return uerr
	}

	return nil
}