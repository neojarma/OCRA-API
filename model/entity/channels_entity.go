package entity

type Channels struct {
	ChannelId    string  `form:"channel_id" json:"channelId"`
	UserId       string  `form:"user_id" json:"userId,omitempty"`
	Name         string  `form:"channel_name" json:"channelName"`
	ProfileImage *string `form:"profileImage" json:"channelProfile"`
}

type Channel struct {
	ChannelId    string  `json:"channelId"`
	Name         string  `json:"channelName"`
	ProfileImage *string `json:"profileImage"`
	BannerImage  *string `json:"bannerImage"`
	UserId       string  `json:"userId"`
	CreatedAt    int64   `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt    int64   `json:"updatedAt" gorm:"autoUpdateTime:milli"`
	Subscriber   int64   `json:"subscriber"`
}
