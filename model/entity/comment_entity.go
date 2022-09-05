package entity

type Comments struct {
	CommentId string `json:"commentId"`
	VideoId   string `json:"videoId"`
	ChannelId string `json:"channelId"`
	Comment   string `json:"comment"`
	CreatedAt int64  `json:"createdAt" gorm:"autoCreateTime:milli"`
}
