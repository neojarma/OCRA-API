package video_service

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
)

type VideoService interface {
	GetAllVideos(page, limit int) (*response.VideosResponse, error)
	GetDetailVideos(videoId string) (*joins_model.DetailVideoJoin, error)
	CreateVideo(req *entity.Videos) (*entity.Videos, error)
	UpdateVideo(req *entity.Videos) (*entity.Videos, error)
}
