package channel_service

import (
	"context"
	"mime/multipart"
	"ocra_server/helper"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/request"
	channel_repository "ocra_server/repository/channels"

	firebase_service "ocra_server/service/firebase"
	"sync"
)

type ChannelServiceImpl struct {
	ChannelRepository channel_repository.ChannelRepository
	FirebaseService   firebase_service.FirebaseService
}

func NewChannelService(repo channel_repository.ChannelRepository, firebaseService firebase_service.FirebaseService) ChannelService {
	var doOnce sync.Once
	service := new(ChannelServiceImpl)

	doOnce.Do(func() {
		service = &ChannelServiceImpl{
			ChannelRepository: repo,
			FirebaseService:   firebaseService,
		}
	})

	return service
}

func (service *ChannelServiceImpl) CreateChannel(req *entity.Channels, image *multipart.FileHeader) (*entity.Channels, error) {
	req.ChannelId = helper.GetRandomString(24)

	// is user already has a channel
	err := service.ChannelRepository.IsUserHasChannel(req.UserId)
	if err != nil {
		return nil, err
	}

	if image != nil {
		// upload image
		profileImagePath := helper.GetChannelProfileImagePath(req.ChannelId)
		profileImageLink, err := service.FirebaseService.CreateResource(context.Background(), profileImagePath, image)
		if err != nil {
			return nil, err
		}

		req.ProfileImage = &profileImageLink
	}

	domainReq := &entity.Channel{
		ChannelId:    req.ChannelId,
		Name:         req.Name,
		ProfileImage: req.ProfileImage,
		UserId:       req.UserId,
	}

	err = service.ChannelRepository.CreateChannel(domainReq)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (service *ChannelServiceImpl) DetailChannel(req *request.GetDetailChannelRequest) (*joins_model.ChannelVideoJoin, error) {
	offset, limit, _ := helper.ParseOffsetValue(req.Page, req.Limit)
	domainRes, err := service.ChannelRepository.DetailChannel(req.ChannelId, req.Exclude, offset, limit)
	if err != nil {
		return nil, err
	}

	channelDetail, err := service.ChannelRepository.GetOnlyChannelData(req.ChannelId)
	if err != nil {
		return nil, err
	}

	parsedResult := service.parsingDomainResult(domainRes)
	if req.UserId == "" {
		parsedResult.UserId = nil
		parsedResult.IsSubcribingChannel = false
	} else {
		parsedResult.UserId = &req.UserId
		parsedResult.IsSubcribingChannel = service.ChannelRepository.IsUserSubscribeThisChannel(req.UserId, req.ChannelId)
	}

	parsedResult.Channel = channelDetail

	return parsedResult, nil
}

func (service *ChannelServiceImpl) parsingDomainResult(domainRes []*joins_model.VideoChannelJoin) *joins_model.ChannelVideoJoin {
	join := new(joins_model.ChannelVideoJoin)

	if len(domainRes) == 0 {
		join.Videos = make([]*entity.Videos, 0)
		return join
	}

	for _, v := range domainRes {
		join.Videos = append(join.Videos, v.Video)
	}

	return join
}
