package videos_repository

import (
	"errors"
	"ocra_server/helper"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
	"sync"

	"gorm.io/gorm"
)

type VideosRepositoryImpl struct {
	Db *gorm.DB
}

func NewVideosRepository(db *gorm.DB) VideosRepository {
	var doOnce sync.Once
	repo := new(VideosRepositoryImpl)

	doOnce.Do(func() {
		repo = &VideosRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repository *VideosRepositoryImpl) GetAllVideos(offset, limit int) ([]*joins_model.HomeVideoJoin, error) {

	videoModel := new(entity.Videos)
	joinModel := make([]*joins_model.HomeVideoJoin, 0)

	paginationFunc := helper.GetPaginationFunc(repository.Db, offset, limit)

	err := repository.Db.Model(videoModel).Select("videos.video_id", "videos.channel_id", "videos.thumbnail", "videos.video", "videos.title", "videos.created_at", "videos.views_count", "channels.channel_id", "channels.name", "channels.profile_image").Joins("JOIN channels on videos.channel_id = channels.channel_id").Scopes(paginationFunc).Find(&joinModel).Error
	if err != nil {
		return nil, err
	}

	return joinModel, nil
}

func (repository *VideosRepositoryImpl) GetAllSubscribedVideos(offset, limit int, userId string) ([]*joins_model.HomeVideoJoin, error) {
	videoModel := new(entity.Videos)
	joinModel := make([]*joins_model.HomeVideoJoin, 0)

	paginationFunc := helper.GetPaginationFunc(repository.Db, offset, limit)

	err := repository.Db.Model(videoModel).Select("videos.video_id", "videos.channel_id", "videos.thumbnail", "videos.video", "videos.title", "videos.created_at", "videos.views_count", "channels.channel_id", "channels.name", "channels.profile_image").Joins("JOIN channels on videos.channel_id = channels.channel_id").Joins("JOIN subscribes ON subscribes.channel_id = videos.channel_id").Where("subscribes.user_id = ? ", userId).Scopes(paginationFunc).Find(&joinModel).Error
	if err != nil {
		return nil, err
	}

	return joinModel, nil
}

func (repository *VideosRepositoryImpl) CountTotalRows() int64 {
	videoModel := make([]entity.Videos, 0)

	result := repository.Db.Find(&videoModel)

	return result.RowsAffected
}

func (repository *VideosRepositoryImpl) GetDetailVideos(videoId string) (*joins_model.DetailVideoJoin, error) {
	videoModel := new(entity.Videos)
	joinModel := new(joins_model.DetailVideoJoin)

	result := repository.Db.Model(videoModel).Select("videos.video_id", "videos.channel_id", "videos.thumbnail", "videos.video", "videos.title", "videos.description", "videos.tags", "videos.likes_count", "videos.dislikes_count", "videos.created_at", "videos.views_count", "channels.name", "channels.channel_id", "channels.profile_image", "channels.subscriber").Joins("JOIN channels on videos.channel_id = channels.channel_id").Where("videos.video_id = ?", videoId).Find(joinModel)

	if result.RowsAffected == 0 {
		return nil, errors.New(response.MessageNoVideo)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return joinModel, nil
}

func (repository *VideosRepositoryImpl) CreateVideo(req *entity.Video) error {
	if err := repository.Db.Create(req).Error; err != nil {
		return err
	}

	return nil
}

func (repository *VideosRepositoryImpl) UpdateVideo(req *entity.Video) error {
	return nil
}

func (repository *VideosRepositoryImpl) IncrementViewsCount(videoId string) error {
	model := new(entity.Videos)
	return repository.Db.Model(model).Where("video_id = ? ", videoId).UpdateColumn("views_count", gorm.Expr("views_count + ?", 1)).Error
}
