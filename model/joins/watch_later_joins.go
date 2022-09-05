package joins_model

type WatchLaterJoins struct {
	WatchId   string `json:"watchLaterId"`
	VideoId   string `json:"videoId"`
	ChannelId string `json:"channelId"`
	Thumbnail string `json:"videoThumbnail"`
	Title     string `json:"videoTitle"`
	Name      string `json:"channelName"`
}
