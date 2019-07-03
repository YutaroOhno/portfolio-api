package mailer

import (
  "gopkg.in/gomail.v2"
  "os"
)

type GoMailer struct {
	Message *gomail.Message
}

func NewMailer() *GoMailer {
	return &GoMailer{
		Message: gomail.NewMessage(),
	}
}

func (mailer *GoMailer) Send(name string, address string, content string) error {
	mailer.Message.SetHeader("From", address)
	mailer.Message.SetHeader("To", os.Getenv("TO_MAIL_ADDRESS"))
	mailer.Message.SetHeader("Subject", name)
	mailer.Message.SetBody("text/plain", content)

	var d *gomail.Dialer
	if os.Getenv("ENV") == "prod" {
		d = gomail.NewDialer("smtp.example.com", 587, os.Getenv("SENDGRID_USERNAME"), os.Getenv("SENDGRID_PASSWORD"))
	} else {
		d = gomail.NewDialer("mailhog", 1025, "", "")
	}

	if err := d.DialAndSend(mailer.Message); err != nil {
		return err
	}
	
	return nil
}