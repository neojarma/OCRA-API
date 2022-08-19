package response

type UserResponse struct {
	UserId           string  `json:"userId"`
	FullName         string  `json:"fullName"`
	UserProfileImage *string `json:"userProfileImage"`
	Email            string  `json:"email"`
	CreatedAt        int64   `json:"createdAt"`
	ChannelId        *string `json:"channelId"`
}
