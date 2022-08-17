package request

type VideoRequest struct {
}

type DetailVideoRequest struct {
	UserId    string `query:"user"`
	VideoId   string `query:"video"`
	ChannelId string `query:"channel"`
}
