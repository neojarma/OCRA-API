package verification_repository

import (
	"errors"
	"ocra_server/model/entity"
	"ocra_server/model/response"
	"sync"
	"time"

	"gorm.io/gorm"
)

type VerificationRepositoryImpl struct {
	Db *gorm.DB
}

func NewVerificationRepository(db *gorm.DB) VerificationRepository {
	var doOnce sync.Once
	repo := new(VerificationRepositoryImpl)

	doOnce.Do(func() {
		repo = &VerificationRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repo *VerificationRepositoryImpl) VerifiedUserEmail(request *entity.Verifications) error {
	now := time.Now().UnixMilli()
	result := repo.Db.Where("token = ?", request.Token).First(request)

	isAlreadyVerified := result.RowsAffected == 0
	if isAlreadyVerified {
		return errors.New(response.MessageUserIsAlreadyVerified)
	}

	needToResendCode := request.ExpiresAt < now
	if needToResendCode {
		return errors.New(response.MessageFailedVerifyEmail)
	}

	return nil
}

func (repo *VerificationRepositoryImpl) CreateVerificationToken(request *entity.Verifications) error {
	if err := repo.Db.Create(request).Error; err != nil {
		return err
	}

	return nil
}

func (repo *VerificationRepositoryImpl) DeleteUserVerification(request *entity.Verifications) error {
	if err := repo.Db.Where("email = ?", request.Email).Delete(request).Error; err != nil {
		return err
	}

	return nil
}
