package watchlater_service

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
)

type HistoryService interface {
	GetAllHistoryRecords(userId string) ([]*joins_model.HistoryJoins, error)
	CreateHistoryRecord(req *entity.Histories) error
	DeleteHistoryRecord(req *entity.Histories) error
}
