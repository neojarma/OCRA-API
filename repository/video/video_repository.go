package videos_repository

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
)

type VideosRepository interface {
	GetAllVideos(offset, limit int) ([]*joins_model.HomeVideoJoin, error)
	GetDetailVideos(videoId string) (*joins_model.DetailVideoJoin, error)
	CountTotalRows() int64
	CreateVideo(req *entity.Videos) (*entity.Videos, error)
	UpdateVideo(req *entity.Videos) (*entity.Videos, error)
}
