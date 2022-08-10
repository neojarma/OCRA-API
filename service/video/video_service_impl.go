package video_service

import (
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
	videos_repository "ocra_server/repository/video"
	"sync"
)

type VideoServiceImpl struct {
	Repo videos_repository.VideosRepository
}

func NewVideoService(repo videos_repository.VideosRepository) VideoService {
	var doOnce sync.Once
	service := new(VideoServiceImpl)

	doOnce.Do(func() {
		service = &VideoServiceImpl{
			Repo: repo,
		}
	})

	return service
}

func (service *VideoServiceImpl) GetAllVideos(page, limit int) (*response.VideosResponse, error) {
	totalRows := make(chan int64)
	defer close(totalRows)

	offset := (page - 1) * limit
	resultVideos, err := service.Repo.GetAllVideos(offset, limit)
	if err != nil {
		return nil, err
	}

	go func() {
		totalRows <- service.Repo.CountTotalRows()
	}()

	rows := <-totalRows
	return &response.VideosResponse{
		Page:        page,
		Limit:       limit,
		TotalVideos: rows,
		Videos:      resultVideos,
	}, nil
}

func (service *VideoServiceImpl) GetDetailVideos(videoId string) (*joins_model.DetailVideoJoin, error) {
	return service.Repo.GetDetailVideos(videoId)
}

func (service *VideoServiceImpl) CreateVideo(req *entity.Videos) (*entity.Videos, error) {
	panic("not implemented") // TODO: Implement
}

func (service *VideoServiceImpl) UpdateVideo(req *entity.Videos) (*entity.Videos, error) {
	panic("not implemented") // TODO: Implement
}
