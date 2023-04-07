package channel_service

import (
	"mime/multipart"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/request"
)

type ChannelService interface {
	CreateChannel(req *entity.Channels, image *multipart.FileHeader) (*entity.Channels, error)
	DetailChannel(req *request.GetDetailChannelRequest) (*joins_model.ChannelVideoJoin, error)
}
