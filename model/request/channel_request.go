package request

type GetDetailChannelRequest struct {
	UserId    string `query:"user"`
	ChannelId string `query:"channel"`
	Page      string `query:"page"`
	Limit     string `query:"limit"`
	Exclude   string `query:"exclude"`
}
