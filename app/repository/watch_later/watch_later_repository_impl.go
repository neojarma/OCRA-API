package watchlater_repository

import (
	"errors"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
	"sync"

	"gorm.io/gorm"
)

type WatchLaterRepositoryImpl struct {
	Db *gorm.DB
}

func NewWatchLateRepository(db *gorm.DB) WatchLaterRepository {
	var doOnce sync.Once
	repo := new(WatchLaterRepositoryImpl)

	doOnce.Do(func() {
		repo = &WatchLaterRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repository *WatchLaterRepositoryImpl) IsRecordAlreadyExist(req *entity.Watch_Laters) error {
	isExist := repository.Db.Where("video_id = ? ", req.VideoId).Where("user_id = ?", req.UserId).Find(req).RowsAffected != 0

	if isExist {
		return errors.New(response.MessageFailedWatchLaterRecordExist)
	}

	return nil
}

func (repository *WatchLaterRepositoryImpl) CreateWatchRecord(req *entity.Watch_Laters) error {
	if !repository.isChannelHasThisVideo(req.ChannelId, req.VideoId) {
		return errors.New(response.MessageFailedInsertWacthRecord)
	}

	return repository.Db.Omit("watch_id").Create(req).Error
}

func (repository *WatchLaterRepositoryImpl) GetAllWatchLaterRecord(req *entity.Watch_Laters) ([]*joins_model.WatchLaterJoins, error) {
	joinModel := make([]*joins_model.WatchLaterJoins, 0)

	err := repository.Db.Model(req).Select("watch_laters.watch_id", "watch_laters.video_id", "watch_laters.channel_id", "videos.thumbnail", "videos.title", "channels.name").Joins("JOIN videos ON watch_laters.video_id = videos.video_id").Joins("JOIN channels ON watch_laters.channel_id = channels.channel_id").Where("watch_laters.user_id = ?", req.UserId).Find(&joinModel).Error

	if err != nil {
		return nil, err
	}

	return joinModel, nil
}

func (repository *WatchLaterRepositoryImpl) isChannelHasThisVideo(channelId, videoId string) bool {
	videoModel := new(entity.Videos)
	return repository.Db.Where("video_id = ? ", videoId).Where("channel_id = ?", channelId).Find(videoModel).RowsAffected == 1
}

func (repository *WatchLaterRepositoryImpl) DeleteWatchRecord(req *entity.Watch_Laters) error {
	return repository.Db.Where("watch_id = ?", req.WatchId).Delete(req).Error
}
