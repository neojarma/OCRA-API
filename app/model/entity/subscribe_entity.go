package entity

type Subscribes struct {
	SubsId    int    `json:"subsId"`
	ChannelId string `json:"channelId"`
	UserId    string `json:"userId"`
}
