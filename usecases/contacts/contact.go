package contacts

import (
	"portfolio-api/usecases/mailer"
	"portfolio-api/usecases/logging"
	"portfolio-api/usecases/ports"
)

type ContactUsecase struct {
	Mailer  mailer.Mailer
	Logging logging.Logging
}

func (usecase *ContactUsecase) SendMail(input *ports.ContactInputPort) error {
	err := usecase.Mailer.Send(input.Name, input.MailAddress, input.Content)
	if err != nil {
		return err
	}

	return nil
}