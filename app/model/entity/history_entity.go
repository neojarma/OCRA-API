package entity

type Histories struct {
	HistoryId int    `json:"historyId"`
	VideoId   string `json:"videoId" validate:"required"`
	ChannelId string `json:"channelId" validate:"required"`
	UserId    string `json:"userId" validate:"required"`
}
