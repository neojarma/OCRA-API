package entity

type Watch_Laters struct {
	WatchId   string
	VideoId   string `json:"videoId" validate:"required"`
	ChannelId string `json:"channelId" validate:"required"`
	UserId    string `json:"userId" validate:"required"`
}
