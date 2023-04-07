package subscribe_repository

import "ocra_server/model/entity"

type SubscribeRepository interface {
	IsUserSubscribeThisChannel(req *entity.Subscribes) bool
	CreateSubsRecord(req *entity.Subscribes) error
	DeleteSubsRecord(req *entity.Subscribes) error
}
