package subscribe_service

import (
	"errors"
	"ocra_server/model/entity"
	"ocra_server/model/response"
	subscribe_repository "ocra_server/repository/subscribe"
	"sync"
)

type SubscribeServiceImpl struct {
	SubsRepo subscribe_repository.SubscribeRepository
}

func NewSubsService(repo subscribe_repository.SubscribeRepository) SubscribeService {
	var doOnce sync.Once
	service := new(SubscribeServiceImpl)

	doOnce.Do(func() {
		service = &SubscribeServiceImpl{
			SubsRepo: repo,
		}
	})

	return service
}

func (service *SubscribeServiceImpl) IsUserSubscribeThisChannel(req *entity.Subscribes) bool {
	return service.SubsRepo.IsUserSubscribeThisChannel(req)
}

func (service *SubscribeServiceImpl) CreateSubsRecord(req *entity.Subscribes) error {
	if service.SubsRepo.IsUserSubscribeThisChannel(req) {
		return errors.New(response.MessageUserAlreadySubscribe)
	}

	return service.SubsRepo.CreateSubsRecord(req)
}

func (service *SubscribeServiceImpl) DeleteSubsRecord(req *entity.Subscribes) error {
	return service.SubsRepo.DeleteSubsRecord(req)
}
