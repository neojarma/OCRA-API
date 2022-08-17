package subscribe_repository

import (
	"ocra_server/model/entity"
	"sync"

	"gorm.io/gorm"
)

type SubscribeRepositoryImpl struct {
	Db *gorm.DB
}

func NewSubsRepository(db *gorm.DB) SubscribeRepository {
	var doOnce sync.Once
	repo := new(SubscribeRepositoryImpl)

	doOnce.Do(func() {
		repo = &SubscribeRepositoryImpl{
			Db: db,
		}
	})

	return repo
}

func (repository *SubscribeRepositoryImpl) IsUserSubscribeThisChannel(req *entity.Subscribes) bool {
	newModel := new(entity.Subscribes)
	return repository.Db.Where("channel_id = ?", req.ChannelId).Where("user_id = ?", req.UserId).Find(newModel).RowsAffected == 1
}

func (repository *SubscribeRepositoryImpl) CreateSubsRecord(req *entity.Subscribes) error {
	return repository.Db.Create(req).Error
}

func (repository *SubscribeRepositoryImpl) DeleteSubsRecord(req *entity.Subscribes) error {
	return repository.Db.Where("channel_id = ?", req.ChannelId).Where("user_id = ?", req.UserId).Delete(req).Error
}
