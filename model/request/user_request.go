package request

type UserRequest struct {
	UserId       string  `json:"userId"`
	FullName     string  `json:"fullName"`
	ProfileImage *string `json:"userProfileImage"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
}
