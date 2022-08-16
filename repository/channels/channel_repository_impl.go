package channel_repository

import (
	"errors"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
	"strings"
	"sync"

	"gorm.io/gorm"
)

type ChannelRepositoryImpl struct {
	Db *gorm.DB
}

func NewChannelRepository(db *gorm.DB) ChannelRepository {
	var doOnce sync.Once
	repo := new(ChannelRepositoryImpl)

	doOnce.Do(func() {
		repo = &ChannelRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repository *ChannelRepositoryImpl) DetailChannel(channelId, excludeVideo string, offset, limit int) ([]*joins_model.VideoChannelJoin, error) {
	videoModel := new(entity.Videos)
	joinModel := make([]*joins_model.VideoChannelJoin, 0)

	paginationFunc := func(d *gorm.DB) *gorm.DB {
		return d.Offset(offset).Limit(limit)
	}

	err := repository.Db.Model(videoModel).Select("videos.video_id", "videos.channel_id", "videos.thumbnail", "videos.video", "videos.title", "videos.created_at", "videos.views_count", "channels.channel_id", "channels.name", "channels.profile_image", "channels.banner_image", "channels.created_at", "channels.subscriber").Joins("JOIN channels on videos.channel_id = channels.channel_id").Scopes(paginationFunc).Where("channels.channel_id = ? ", channelId).Where("videos.video_id != ?", excludeVideo).Find(&joinModel).Error
	if err != nil {
		return nil, err
	}

	return joinModel, nil
}

func (repository *ChannelRepositoryImpl) IsUserSubscribeThisChannel(userId, channelId string) bool {
	model := new(entity.Subscribes)
	result := repository.Db.Where("user_id = ?", userId).Where("channel_id = ?", channelId).Find(model)
	return result.RowsAffected == 1
}

func (repository *ChannelRepositoryImpl) GetOnlyChannelData(channelId string) (*entity.Channel, error) {
	model := new(entity.Channel)
	result := repository.Db.Where("channel_id = ?", channelId).Find(model)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New(response.MessageNoChannelWithID)
	}

	return model, nil
}

func (repository *ChannelRepositoryImpl) CreateChannel(req *entity.Channel) error {
	err := repository.Db.Create(req).Error

	if err != nil {
		errStr := strings.Split(err.Error(), " ")
		isForeignKeyErr := errStr[1] == "1452:"
		if isForeignKeyErr {
			return errors.New(response.MessageInvalidUserId)
		}
	}

	return err
}

func (repository *ChannelRepositoryImpl) IsUserHasChannel(userId string) error {
	model := new(entity.Channels)
	result := repository.Db.Where("user_id = ?", userId).Find(model)
	if result.RowsAffected != 0 {
		return errors.New(response.MessageUserAlreadyHasAChannel)
	}

	return result.Error
}

func (repository *ChannelRepositoryImpl) UpdateChannel(req *entity.Channels) error {
	result := repository.Db.Model(req).Where("channel_id = ?", req.ChannelId).Updates(req)
	if result.RowsAffected == 0 {
		return errors.New(response.MessageNoChannelWithID)
	}

	return result.Error
}
