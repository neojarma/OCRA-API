package request

type UserRequest struct {
	UserId           string `json:"userId"`
	FullName         string `json:"fullName"`
	UserProfileImage string `json:"userProfileImage"`
	Email            string `json:"email"`
	Password         string `json:"password"`
}
