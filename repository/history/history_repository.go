package watchlater_repository

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
)

type HistoryRepository interface {
	IsRecordAlreadyExist(req *entity.Histories) error
	CreateHistoryRecord(req *entity.Histories) error
	GetAllHistoryLaterRecord(req *entity.Histories) ([]*joins_model.HistoryJoins, error)
	DeleteHistoryRecord(req *entity.Histories) error
}
