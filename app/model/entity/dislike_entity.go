package entity

type Dislikes struct {
	DislikeId int    `json:"dislikeId"`
	UserId    string `json:"userId"`
	VideoId   string `json:"videoId"`
}
