package watchlater_repository

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
)

type WatchLaterRepository interface {
	IsRecordAlreadyExist(req *entity.Watch_Laters) error
	CreateWatchRecord(req *entity.Watch_Laters) error
	GetAllWatchLaterRecord(req *entity.Watch_Laters) ([]*joins_model.WatchLaterJoins, error)
	DeleteWatchRecord(req *entity.Watch_Laters) error
}
