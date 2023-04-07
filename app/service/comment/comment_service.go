package comment_service

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/request"
)

type CommentService interface {
	GetVideoComments(req *request.CommentRequest) ([]*joins_model.CommentChannelJoin, error)
	CreateComment(req *entity.Comments) error
	UpdateComment(req *entity.Comments) error
	DeleteComment(req *entity.Comments) error
}
