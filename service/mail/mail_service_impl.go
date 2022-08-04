package mail_service

import (
	"bytes"
	"fmt"
	"html/template"
	"ocra_server/model/request"
	"os"
	"path"
	"sync"

	"gopkg.in/gomail.v2"
)

type MailServiceImpl struct {
	Dialer *gomail.Dialer
}

func NewMailService(dialer *gomail.Dialer) MailService {
	var doOnce sync.Once
	service := new(MailServiceImpl)

	doOnce.Do(func() {
		service = &MailServiceImpl{
			Dialer: dialer,
		}
	})

	return service
}

func (service *MailServiceImpl) CreateVerificationMail(request *request.MailRequest) error {
	body, err := service.getBodyForVerificationToken(request.To, request.Token)
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", request.From)
	mailer.SetHeader("To", request.To)
	mailer.SetHeader("Subject", request.Subject)
	mailer.SetHeader("Reply-To", request.From)
	mailer.SetBody("text/html", body)

	if err := service.Dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

func (service *MailServiceImpl) getBodyForVerificationToken(email, token string) (string, error) {

	data := map[string]string{
		"url": service.generateVerificationLink(email, token),
	}

	return service.parseTemplateFile(data)
}

func (service *MailServiceImpl) generateVerificationLink(email, token string) string {
	host := os.Getenv("HOST")
	schema := os.Getenv("SCHEMA")
	return fmt.Sprintf("%v://%v/api/v1/email-verification?email=%v&token=%v", schema, host, email, token)
}

func (service *MailServiceImpl) parseTemplateFile(data map[string]string) (string, error) {
	filepath := path.Join("service", "mail", "template_email.html")
	template, err := template.ParseFiles(filepath)
	if err != nil {
		return "", err
	}

	writer := new(bytes.Buffer)
	err = template.Execute(writer, data)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}
