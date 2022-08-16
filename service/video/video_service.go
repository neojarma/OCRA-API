package video_service

import (
	"mime/multipart"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
)

type VideoService interface {
	GetAllVideos(page, limit string) (*response.VideosResponse, error)
	GetDetailVideos(videoId string) (*joins_model.DetailVideoJoin, error)
	CreateVideo(req *entity.Video, thumbnail, video *multipart.FileHeader) (*entity.Video, error)
	UpdateVideo(req *entity.Videos) (*entity.Videos, error)
}
