package joins_model

import "ocra_server/model/entity"

type CommentChannelJoin struct {
	Comment *entity.Comments `json:"comment" gorm:"embedded"`
	Channel *entity.Channels `json:"channel" gorm:"embedded"`
}
