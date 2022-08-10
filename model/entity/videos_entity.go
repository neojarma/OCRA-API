package entity

type Videos struct {
	VideoId    string `json:"videoId"`
	ChannelId  string `json:"channelId,omitempty"`
	Thumbnail  string `json:"videothumbnail"`
	Video      string `json:"videoUrl"`
	Title      string `json:"videoTitle"`
	CreatedAt  int64  `json:"createdAt" gorm:"autoCreateTime:milli"`
	ViewsCount int64  `json:"viewsCount"`
}

type Video struct {
	VideoId       string `json:"videoId"`
	ChannelId     string `json:"channelId,omitempty"`
	Thumbnail     string `json:"videothumbnail"`
	Video         string `json:"videoUrl"`
	Title         string `json:"videoTitle"`
	Description   string `json:"videoDesc"`
	Tags          string `json:"videoTags"`
	CreatedAt     int64  `json:"createdAt" gorm:"autoCreateTime:milli"`
	ViewsCount    int64  `json:"viewsCount"`
	LikesCount    int64  `json:"likesCount"`
	DislikesCount int64  `json:"dislikesCount"`
}
