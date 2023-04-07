package verification_service

import (
	"ocra_server/helper"
	"ocra_server/model/entity"
	"ocra_server/model/request"
	user_repository "ocra_server/repository/user"
	verification_repository "ocra_server/repository/verification"
	mail_service "ocra_server/service/mail"
	"sync"
	"time"
)

type VerificationServiceImpl struct {
	Repo          verification_repository.VerificationRepository
	UserRepo      user_repository.UserRepository
	MailerService mail_service.MailService
}

func NewVerificationService(repo verification_repository.VerificationRepository, mailer mail_service.MailService, userRepo user_repository.UserRepository) VerificationService {
	var doOnce sync.Once
	service := new(VerificationServiceImpl)

	doOnce.Do(func() {
		service = &VerificationServiceImpl{
			Repo:          repo,
			UserRepo:      userRepo,
			MailerService: mailer,
		}
	})

	return service
}

func (service *VerificationServiceImpl) CreateVerifToken(email string) error {
	newToken := helper.GetRandomString(32)
	tokenExpired := time.Now().Add(time.Minute * 30).UnixMilli()
	domainVerifReq := &entity.Verifications{
		Email:     email,
		Token:     newToken,
		ExpiresAt: tokenExpired,
	}

	processErr := make(chan error)

	go func(processErr chan<- error) {
		if err := service.Repo.CreateVerificationToken(domainVerifReq); err != nil {
			processErr <- err
		}

		processErr <- nil
	}(processErr)

	mailReq := &request.MailRequest{
		From:    "Ocra Support <noreply-ocra@neojarma.com>",
		To:      email,
		Subject: "OCRA - Email Verification",
		Token:   newToken,
	}

	go func(processErr chan<- error) {
		if err := service.MailerService.CreateVerificationMail(mailReq); err != nil {
			processErr <- err
		}

		processErr <- nil
	}(processErr)

	return <-processErr
}

func (service *VerificationServiceImpl) ValidateVerifToken(email, token string) error {
	req := &entity.Verifications{
		Email: email,
		Token: token,
	}

	if err := service.Repo.VerifiedUserEmail(req); err != nil {
		return err
	}

	processErr := make(chan error)

	go func() {
		reqUser := &entity.Users{
			Email: email,
		}
		processErr <- service.UserRepo.UpdateUserEmailStatus(reqUser)
	}()

	go func() {
		req := &entity.Verifications{
			Email: email,
		}
		processErr <- service.Repo.DeleteUserVerification(req)
	}()

	return <-processErr
}
