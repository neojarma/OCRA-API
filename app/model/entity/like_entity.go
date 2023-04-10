package entity

type Likes struct {
	LikeId  int    `json:"likeId"`
	UserId  string `json:"userId"`
	VideoId string `json:"videoId"`
}
