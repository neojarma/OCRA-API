package comment_repository

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
)

type CommentRepository interface {
	GetCommentByVideoId(videoId string, offset, limit int) ([]*joins_model.CommentChannelJoin, error)
	CreateComment(req *entity.Comments) error
	UpdateComment(req *entity.Comments) error
	DeleteComment(req *entity.Comments) error
}
