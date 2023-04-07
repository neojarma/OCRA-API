package entity

type Dislikes struct {
	DislikeId string `json:"dislikeId"`
	UserId    string `json:"userId"`
	VideoId   string `json:"videoId"`
}
