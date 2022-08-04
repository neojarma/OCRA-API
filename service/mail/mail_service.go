package mail_service

import "ocra_server/model/request"

type MailService interface {
	CreateVerificationMail(request *request.MailRequest) error
}
