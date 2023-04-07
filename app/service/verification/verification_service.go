package verification_service

type VerificationService interface {
	CreateVerifToken(email string) error
	ValidateVerifToken(email, token string) error
}
