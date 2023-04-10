package comment_repository

import (
	"errors"
	"ocra_server/helper"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
	"sync"

	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	Db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	var doOnce sync.Once
	repo := new(CommentRepositoryImpl)

	doOnce.Do(func() {
		repo = &CommentRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repository *CommentRepositoryImpl) GetCommentByVideoId(videoId string, offset, limit int) ([]*joins_model.CommentChannelJoin, error) {
	commentModel := new(entity.Comments)
	joinModel := make([]*joins_model.CommentChannelJoin, 0)

	paginationFunc := helper.GetPaginationFunc(repository.Db, offset, limit)
	err := repository.Db.Model(commentModel).Select("comments.comment_id", "comments.video_id", "comments.channel_id", "comments.comment", "comments.created_at", "channels.channel_id", "channels.name", "channels.profile_image").Joins("JOIN channels on comments.channel_id = channels.channel_id").Scopes(paginationFunc).Where("comments.video_id = ?", videoId).Find(&joinModel).Error

	if err != nil {
		return nil, err
	}

	return joinModel, nil
}

func (repository *CommentRepositoryImpl) CreateComment(req *entity.Comments) error {
	err := repository.Db.Omit("comment_id").Create(req).Error

	if err != nil {
		if helper.IsInvalidForeignKey(err) {
			return errors.New(response.MessageInvalidChannelId)
		}
	}

	return err
}

func (repository *CommentRepositoryImpl) UpdateComment(req *entity.Comments) error {
	err := repository.Db.Where("channel_id = ? ", req.ChannelId).Updates(req).Error

	if err != nil {
		if helper.IsInvalidForeignKey(err) {
			return errors.New(response.MessageInvalidChannelId)
		}
	}

	return err
}

func (repository *CommentRepositoryImpl) DeleteComment(req *entity.Comments) error {
	return repository.Db.Where("comment_id = ? ", req.CommentId).Delete(req).Error
}
