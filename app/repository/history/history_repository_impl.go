package watchlater_repository

import (
	"errors"
	"log"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
	"sync"

	"gorm.io/gorm"
)

type HistoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	var doOnce sync.Once
	repo := new(HistoryRepositoryImpl)

	doOnce.Do(func() {
		repo = &HistoryRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repository *HistoryRepositoryImpl) IsRecordAlreadyExist(req *entity.Histories) error {
	isExist := repository.Db.Where("video_id = ? ", req.VideoId).Where("user_id = ?", req.UserId).Find(req).RowsAffected != 0

	if isExist {
		return errors.New(response.MessageFailedHistoryRecordExist)
	}

	return nil
}

func (repository *HistoryRepositoryImpl) CreateHistoryRecord(req *entity.Histories) error {
	if !repository.isChannelHasThisVideo(req.ChannelId, req.VideoId) {
		return errors.New(response.MessageFailedInsertWacthRecord)
	}

	return repository.Db.Omit("history_id").Create(req).Error
}

func (repository *HistoryRepositoryImpl) GetAllHistoryLaterRecord(req *entity.Histories) ([]*joins_model.HistoryJoins, error) {
	joinModel := make([]*joins_model.HistoryJoins, 0)

	err := repository.Db.Model(req).Select("histories.history_id", "videos.video_id", "videos.title", "videos.thumbnail", "videos.video", "videos.description", "videos.created_at", "videos.views_count", "channels.channel_id", "channels.name").Joins("JOIN videos ON histories.video_id = videos.video_id").Joins("JOIN channels ON histories.channel_id = channels.channel_id").Where("histories.user_id = ?", req.UserId).Find(&joinModel).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return joinModel, nil
}

func (repository *HistoryRepositoryImpl) isChannelHasThisVideo(channelId, videoId string) bool {
	videoModel := new(entity.Videos)
	return repository.Db.Where("video_id = ? ", videoId).Where("channel_id = ?", channelId).Find(videoModel).RowsAffected == 1
}

func (repository *HistoryRepositoryImpl) DeleteHistoryRecord(req *entity.Histories) error {
	return repository.Db.Where("history_id = ?", req.HistoryId).Delete(req).Error
}
