package entity

type Comments struct {
	CommentId string `json:"commentId"`
	VideoId   string `json:"videoId"`
	ChannelId string `json:"channelId"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"createdAt" gorm:"autoCreateTime:milli"`
}
