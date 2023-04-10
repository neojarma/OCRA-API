package entity

type Watch_Laters struct {
	WatchId   int    `json:"watchLaterId"`
	VideoId   string `json:"videoId" validate:"required"`
	ChannelId string `json:"channelId" validate:"required"`
	UserId    string `json:"userId" validate:"required"`
}
