package mailer

type Mailer interface {
	Send(name string, address string, content string) error
}
