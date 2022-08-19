package request

type UserRequest struct {
	UserId       string  `json:"userId" param:"id"`
	FullName     string  `json:"fullName" validate:"required" form:"full_name"`
	ProfileImage *string `json:"userProfileImage"`
	Email        string  `json:"email" validate:"required,email"`
	Password     string  `json:"password" validate:"required,min=12,passwd,containsany=!@#$%^&*()"`
}
