package subscribe_service

import "ocra_server/model/entity"

type SubscribeService interface {
	IsUserSubscribeThisChannel(req *entity.Subscribes) bool
	CreateSubsRecord(req *entity.Subscribes) error
	DeleteSubsRecord(req *entity.Subscribes) error
}
