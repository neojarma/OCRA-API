package entity

type Histories struct {
	HistoryId string `json:"historyId"`
	VideoId   string `json:"videoId" validate:"required"`
	ChannelId string `json:"channelId" validate:"required"`
	UserId    string `json:"userId" validate:"required"`
}
