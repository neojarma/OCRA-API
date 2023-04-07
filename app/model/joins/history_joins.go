package joins_model

import "ocra_server/model/entity"

type HistoryJoins struct {
	HistoryId string         `json:"historyId"`
	Video     *entity.Videos `json:"video" gorm:"embedded"`
	Name      string         `json:"channelName"`
}
