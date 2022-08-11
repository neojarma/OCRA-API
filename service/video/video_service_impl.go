package video_service

import (
	"context"
	"mime/multipart"
	"ocra_server/helper"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/response"
	videos_repository "ocra_server/repository/video"
	firebase_service "ocra_server/service/firebase"
	"strings"
	"sync"
)

type VideoServiceImpl struct {
	Repo            videos_repository.VideosRepository
	FirebaseService firebase_service.FirebaseService
}

func NewVideoService(repo videos_repository.VideosRepository, firebaseService firebase_service.FirebaseService) VideoService {
	var doOnce sync.Once
	service := new(VideoServiceImpl)

	doOnce.Do(func() {
		service = &VideoServiceImpl{
			Repo:            repo,
			FirebaseService: firebaseService,
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

func (service *VideoServiceImpl) CreateVideo(req *entity.Video, thumbnail, video *multipart.FileHeader) (*entity.Video, error) {
	req.VideoId = helper.GetRandomString(11)

	// parsing thumbnail
	thumbnailPath := helper.GetThumbnailFilePath(req.ChannelId, req.VideoId)
	thumbnailUrl, err := service.parsingFormFile(thumbnail, thumbnailPath)
	if err != nil {
		return nil, err
	}
	req.Thumbnail = thumbnailUrl

	// parsing video
	videoPath := helper.GetVideoFilePath(req.ChannelId, req.VideoId)
	videoUrl, err := service.parsingFormFile(video, videoPath)
	if err != nil {
		return nil, err
	}
	req.Video = videoUrl

	// insert db
	if err := service.Repo.CreateVideo(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (service *VideoServiceImpl) parsingFormFile(file *multipart.FileHeader, path string) (string, error) {
	openFile, err := file.Open()
	if err != nil {
		return "", err
	}

	defer openFile.Close()

	newFileName := strings.ReplaceAll(file.Filename, " ", "")
	path = path + newFileName
	service.FirebaseService.CreateResource(context.Background(), path, openFile)

	return service.FirebaseService.GenerateFirebaseStorageLink(path), nil
}

func (service *VideoServiceImpl) UpdateVideo(req *entity.Videos) (*entity.Videos, error) {
	return nil, nil
}
