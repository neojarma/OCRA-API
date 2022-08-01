package response

type UserResponse struct {
	UserId           string  `json:"userId"`
	FullName         string  `json:"fullName"`
	UserProfileImage *string `json:"userProfileImage"`
	Email            string  `json:"email"`
	CreatedAt        string  `json:"createdAt"`
}
