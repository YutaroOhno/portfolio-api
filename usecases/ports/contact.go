package ports

type ContactInputPort struct {
	Name 		string `json:"name"`
	MailAddress string `json:"mail_address"`
	Content 	string `json:"content"`
}
