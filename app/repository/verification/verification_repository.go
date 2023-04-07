package verification_repository

import (
	"ocra_server/model/entity"
)

type VerificationRepository interface {
	CreateVerificationToken(request *entity.Verifications) error
	VerifiedUserEmail(request *entity.Verifications) error
	DeleteUserVerification(request *entity.Verifications) error
}
