package joins_model

import "ocra_server/model/entity"

type HomeVideoJoin struct {
	Video   entity.Videos   `json:"video" gorm:"embedded"`
	Channel entity.Channels `json:"channel" gorm:"embedded"`
}

type DetailVideoJoin struct {
	Video   entity.Video    `json:"video" gorm:"embedded"`
	Channel entity.Channels `json:"channel" gorm:"embedded"`
}
