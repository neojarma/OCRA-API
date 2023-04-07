package comment_service

import (
	"ocra_server/helper"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/request"
	comment_repository "ocra_server/repository/comment"
	"sync"
)

type CommentServiceImpl struct {
	CommentRepository comment_repository.CommentRepository
}

func NewCommentService(repo comment_repository.CommentRepository) CommentService {
	var doOnce sync.Once
	service := new(CommentServiceImpl)

	doOnce.Do(func() {
		service = &CommentServiceImpl{
			CommentRepository: repo,
		}
	})

	return service
}

func (service *CommentServiceImpl) GetVideoComments(req *request.CommentRequest) ([]*joins_model.CommentChannelJoin, error) {
	offset, limit, _ := helper.ParseOffsetValue(req.Page, req.Limit)
	return service.CommentRepository.GetCommentByVideoId(req.VideoId, offset, limit)
}

func (service *CommentServiceImpl) CreateComment(req *entity.Comments) error {
	return service.CommentRepository.CreateComment(req)
}

func (service *CommentServiceImpl) UpdateComment(req *entity.Comments) error {
	return service.CommentRepository.UpdateComment(req)
}

func (service *CommentServiceImpl) DeleteComment(req *entity.Comments) error {
	return service.CommentRepository.DeleteComment(req)
}
