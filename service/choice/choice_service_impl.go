package choice_service

import (
	"errors"
	"ocra_server/model/entity"
	"ocra_server/model/response"
	dislike_repository "ocra_server/repository/dislike"
	like_repository "ocra_server/repository/like"
	"sync"
)

type ChoiceServiceImpl struct {
	LikeRepo    like_repository.LikeRepository
	DislikeRepo dislike_repository.DislikeRepository
}

func NewChoiceService(likeRepo like_repository.LikeRepository, dislikeRepo dislike_repository.DislikeRepository) ChoiceService {
	var doOnce sync.Once
	repo := new(ChoiceServiceImpl)

	doOnce.Do(func() {
		repo = &ChoiceServiceImpl{
			LikeRepo:    likeRepo,
			DislikeRepo: dislikeRepo,
		}
	})

	return repo
}

func (service *ChoiceServiceImpl) CreateLikeRecord(req *entity.Likes) error {

	go service.DislikeRepo.DeleteDislike(&entity.Dislikes{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	})

	if service.LikeRepo.IsUserAlreadyLikeThisVideo(req) {
		return errors.New(response.MessageAlreadyLikeThisVideo)
	}

	return service.LikeRepo.CreateLike(req)
}

func (service *ChoiceServiceImpl) IsUserLikeTheVideo(req *entity.Likes) bool {
	return service.LikeRepo.IsUserAlreadyLikeThisVideo(&entity.Likes{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	})
}

func (service *ChoiceServiceImpl) CreateDislikeRecord(req *entity.Dislikes) error {

	go service.LikeRepo.DeleteLike(&entity.Likes{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	})

	if service.DislikeRepo.IsUserAlreadyDislikeThisVideo(req) {
		return errors.New(response.MessageAlreadyDislikeThisVideo)
	}

	return service.DislikeRepo.CreateDislike(&entity.Dislikes{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	})
}

func (service *ChoiceServiceImpl) IsUserDislikeTheVideo(req *entity.Dislikes) bool {
	return service.DislikeRepo.IsUserAlreadyDislikeThisVideo(&entity.Dislikes{
		UserId:  req.UserId,
		VideoId: req.VideoId,
	})
}
