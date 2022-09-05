package watchlater_service

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	history_repository "ocra_server/repository/history"
	"sync"
)

type HistoryServiceImpl struct {
	Repository history_repository.HistoryRepository
}

func NewHistoryService(repo history_repository.HistoryRepository) HistoryService {
	var doOnce sync.Once
	service := new(HistoryServiceImpl)

	doOnce.Do(func() {
		service = &HistoryServiceImpl{
			Repository: repo,
		}
	})

	return service
}

func (service *HistoryServiceImpl) GetAllHistoryRecords(userId string) ([]*joins_model.HistoryJoins, error) {
	return service.Repository.GetAllHistoryLaterRecord(&entity.Histories{UserId: userId})
}

func (service *HistoryServiceImpl) CreateHistoryRecord(req *entity.Histories) error {
	if err := service.Repository.IsRecordAlreadyExist(req); err != nil {
		return err
	}

	return service.Repository.CreateHistoryRecord(req)
}

func (service *HistoryServiceImpl) DeleteHistoryRecord(req *entity.Histories) error {
	return service.Repository.DeleteHistoryRecord(req)
}
