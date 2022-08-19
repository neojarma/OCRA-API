package request

type UserRequest struct {
	UserId       string  `json:"userId"`
	FullName     string  `json:"fullName" validate:"required"`
	ProfileImage *string `json:"userProfileImage"`
	Email        string  `json:"email" validate:"required,email"`
	Password     string  `json:"password" validate:"required,min=12,passwd,containsany=!@#$%^&*()"`
}
