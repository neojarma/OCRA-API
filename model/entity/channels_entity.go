package entity

type Channels struct {
	ChannelId    string  `json:"channelId"`
	Name         string  `json:"channelName"`
	ProfileImage *string `json:"profileImage"`
	BannerImage  *string `json:"bannerImage,omitempty"`
	UserId       string  `json:"userId,omitempty"`
	CreatedAt    int64   `json:"createdAt,omitempty" gorm:"autoCreateTime:milli"`
	UpdatedAt    int64   `json:"updatedAt,omitempty" gorm:"autoUpdateTime:milli"`
	Subscriber   int64   `json:"subscriber,omitempty"`
}
