package watchlater_service

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	watchlater_repository "ocra_server/repository/watch_later"
	"sync"
)

type WatchLaterServiceImpl struct {
	Repository watchlater_repository.WatchLaterRepository
}

func NewWatchLaterService(repo watchlater_repository.WatchLaterRepository) WatchLaterService {
	var doOnce sync.Once
	service := new(WatchLaterServiceImpl)

	doOnce.Do(func() {
		service = &WatchLaterServiceImpl{
			Repository: repo,
		}
	})

	return service
}

func (service *WatchLaterServiceImpl) GetAllWatchLaterRecords(userId string) ([]*joins_model.WatchLaterJoins, error) {
	return service.Repository.GetAllWatchLaterRecord(&entity.Watch_Laters{UserId: userId})
}

func (service *WatchLaterServiceImpl) CreateWatchLaterRecord(req *entity.Watch_Laters) error {
	if err := service.Repository.IsRecordAlreadyExist(req); err != nil {
		return err
	}

	return service.Repository.CreateWatchRecord(req)
}

func (service *WatchLaterServiceImpl) DeleteWatchLaterRecord(req *entity.Watch_Laters) error {
	return service.Repository.DeleteWatchRecord(req)
}
