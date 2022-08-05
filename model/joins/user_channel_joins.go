package joins_model

type UserChannelJoin struct {
	UserId       string  `json:"userId"`
	FullName     string  `json:"fullName"`
	ProfileImage *string `json:"profileImage"`
	Email        string  `json:"email"`
	CreatedAt    int64   `json:"createdAt"`
	ChannelId    *string `json:"channelId"`
	SessionId    *string `json:"sessionId,omitempty"`
}
