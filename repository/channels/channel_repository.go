package channel_repository

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
)

type ChannelRepository interface {
	DetailChannel(channelId string, offset, limit int) ([]*joins_model.VideoChannelJoin, error)
	IsUserSubscribeThisChannel(userId, channelId string) bool
	GetOnlyChannelData(channelId string) (*entity.Channel, error)
	CreateChannel(req *entity.Channel) error
	IsUserHasChannel(userId string) error
	UpdateChannel(req *entity.Channels) error
}
