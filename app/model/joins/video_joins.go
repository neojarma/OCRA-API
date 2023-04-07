package joins_model

import "ocra_server/model/entity"

type HomeVideoJoin struct {
	Video   *entity.Videos   `json:"video" gorm:"embedded"`
	Channel *entity.Channels `json:"channel" gorm:"embedded"`
}

type DetailVideoJoin struct {
	UserId         *string         `json:"userId"`
	IsSubscribe    bool            `json:"isSubscribe"`
	IsLikeVideo    bool            `json:"isLikeVideo"`
	IsDislikeVideo bool            `json:"isDislikeVideo"`
	Video          entity.Video    `json:"video" gorm:"embedded"`
	Channel        entity.Channels `json:"channel" gorm:"embedded"`
}

type VideoChannelJoin struct {
	Video   *entity.Videos  `json:"video" gorm:"embedded"`
	Channel *entity.Channel `json:"channel" gorm:"embedded"`
}

type ChannelVideoJoin struct {
	UserId              *string          `json:"userId"`
	IsSubcribingChannel bool             `json:"isSubscribing"`
	Channel             *entity.Channel  `json:"channel" gorm:"embedded"`
	Videos              []*entity.Videos `json:"videos" gorm:"embedded"`
}
