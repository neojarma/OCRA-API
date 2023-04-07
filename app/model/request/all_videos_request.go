package request

type AllVideosRequest struct {
	Page         string `query:"page"`
	Size         string `query:"size"`
	IsSubscribed string `query:"subscribed"`
	UserId       string
}
