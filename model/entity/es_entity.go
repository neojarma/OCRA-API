package entity

type ElasticsearchVideo struct {
	VideoId    string `json:"videoId"`
	VideoTitle string `json:"videoTitle"`
	VideoDesc  string `json:"videoDesc"`
	VideoTags  string `json:"videoTags"`
	CreatedAt  int64  `json:"createdAt"`
}
