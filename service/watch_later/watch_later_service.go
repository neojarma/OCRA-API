package watchlater_service

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
)

type WatchLaterService interface {
	GetAllWatchLaterRecords(userId string) ([]*joins_model.WatchLaterJoins, error)
	CreateWatchLaterRecord(req *entity.Watch_Laters) error
	DeleteWatchLaterRecord(req *entity.Watch_Laters) error
}
