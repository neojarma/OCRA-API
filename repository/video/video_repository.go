package videos_repository

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
)

type VideosRepository interface {
	GetAllVideos(offset, limit int) ([]*joins_model.HomeVideoJoin, error)
	GetAllSubscribedVideos(offset, limit int, userId string) ([]*joins_model.HomeVideoJoin, error)
	GetDetailVideos(videoId string) (*joins_model.DetailVideoJoin, error)
	CountTotalRows() int64
	CreateVideo(req *entity.Video) error
	UpdateVideo(req *entity.Video) error
	IncrementViewsCount(videoId string) error
	Find(videosId ...string) ([]*joins_model.HomeVideoJoin, error)
}
