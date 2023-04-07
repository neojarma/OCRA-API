package channel_controller

import (
	"net/http"
	"ocra_server/model/entity"
	"ocra_server/model/request"
	"ocra_server/model/response"
	channel_service "ocra_server/service/channel"
	"sync"

	"github.com/labstack/echo/v4"
)

type ChannelControllerImpl struct {
	ChannelService channel_service.ChannelService
}

func NewChannelController(service channel_service.ChannelService) ChannelController {
	var doOnce sync.Once
	controller := new(ChannelControllerImpl)

	doOnce.Do(func() {
		controller = &ChannelControllerImpl{
			ChannelService: service,
		}
	})

	return controller
}

func (controller *ChannelControllerImpl) CreateChannel(ctx echo.Context) error {
	req := new(entity.Channels)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	image, err := ctx.FormFile("profile_image")
	if err != nil {
		image = nil
	}

	res, err := controller.ChannelService.CreateChannel(req, image)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessCreateChannel,
		Data:    res,
	})
}

func (controller *ChannelControllerImpl) DetailChannel(ctx echo.Context) error {
	req := new(request.GetDetailChannelRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	res, err := controller.ChannelService.DetailChannel(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessGetDetailChannel,
		Data:    res,
	})
}
