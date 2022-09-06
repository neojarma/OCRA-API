package video_service

import (
	"context"
	"mime/multipart"
	"ocra_server/helper"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/request"
	"ocra_server/model/response"
	videos_repository "ocra_server/repository/video"
	choice_service "ocra_server/service/choice"
	firebase_service "ocra_server/service/firebase"
	subscribe_service "ocra_server/service/subscribe"
	"strconv"
	"sync"
)

type VideoServiceImpl struct {
	Repo             videos_repository.VideosRepository
	FirebaseService  firebase_service.FirebaseService
	SubscribeService subscribe_service.SubscribeService
	ChoiceService    choice_service.ChoiceService
}

func NewVideoService(
	repo videos_repository.VideosRepository,
	firebaseService firebase_service.FirebaseService,
	subsService subscribe_service.SubscribeService,
	choiceService choice_service.ChoiceService) VideoService {
	var doOnce sync.Once
	service := new(VideoServiceImpl)

	doOnce.Do(func() {
		service = &VideoServiceImpl{
			Repo:             repo,
			FirebaseService:  firebaseService,
			SubscribeService: subsService,
			ChoiceService:    choiceService,
		}
	})

	return service
}

func (service *VideoServiceImpl) GetAllVideos(req *request.AllVideosRequest) (*response.VideosResponse, error) {
	offset, limit, newPage := helper.ParseOffsetValue(req.Page, req.Size)

	res, err := strconv.ParseBool(req.IsSubscribed)
	if err == nil && res {
		subscribedList, err := service.Repo.GetAllSubscribedVideos(offset, limit, req.UserId)
		if err != nil {
			return nil, err
		}

		return &response.VideosResponse{
			Page:   newPage,
			Limit:  limit,
			Videos: subscribedList,
		}, nil
	}

	allVids, err := service.Repo.GetAllVideos(offset, limit)
	if err != nil {
		return nil, err
	}

	return &response.VideosResponse{
		Page:   newPage,
		Limit:  limit,
		Videos: allVids,
	}, nil
}

func (service *VideoServiceImpl) GetDetailVideos(req *request.DetailVideoRequest) (*joins_model.DetailVideoJoin, error) {

	result, err := service.Repo.GetDetailVideos(req.VideoId)
	if err != nil {
		return nil, err
	}

	if req.UserId == "" {
		result.UserId = nil
		result.IsSubscribe = false
		result.IsDislikeVideo = false
		result.IsLikeVideo = false
	} else {
		result.UserId = &req.UserId

		result.IsDislikeVideo = service.ChoiceService.IsUserDislikeTheVideo(&entity.Dislikes{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})

		result.IsLikeVideo = service.ChoiceService.IsUserLikeTheVideo(&entity.Likes{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})

		result.IsSubscribe = service.SubscribeService.IsUserSubscribeThisChannel(&entity.Subscribes{
			ChannelId: req.ChannelId,
			UserId:    req.UserId,
		})
	}

	result.Channel.ChannelId = req.ChannelId

	return result, err
}

func (service *VideoServiceImpl) CreateVideo(req *entity.Video, thumbnail, video *multipart.FileHeader) (*entity.Video, error) {
	req.VideoId = helper.GetRandomString(11)

	// parsing thumbnail
	thumbnailPath := helper.GetThumbnailFilePath(req.ChannelId, req.VideoId)
	thumbnailUrl, err := service.FirebaseService.CreateResource(context.Background(), thumbnailPath, thumbnail)
	if err != nil {
		return nil, err
	}
	req.Thumbnail = thumbnailUrl

	// parsing video
	videoPath := helper.GetVideoFilePath(req.ChannelId, req.VideoId)
	videoUrl, err := service.FirebaseService.CreateResource(context.Background(), videoPath, video)
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

func (service *VideoServiceImpl) UpdateVideo(req *entity.Videos) (*entity.Videos, error) {
	return nil, nil
}
