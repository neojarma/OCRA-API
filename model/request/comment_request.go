package request

type CommentRequest struct {
	Page    string `query:"page"`
	Limit   string `query:"limit"`
	VideoId string `query:"video"`
	UserId  string `query:"user"`
}
