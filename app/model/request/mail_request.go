package request

type MailRequest struct {
	From    string
	To      string
	Subject string
	Token   string
}
