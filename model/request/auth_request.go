package request

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=12,passwd,containsany=!@#$%^&*()"`
}
