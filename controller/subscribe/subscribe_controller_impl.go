package subscribe_controller

import (
	"net/http"
	"ocra_server/model/entity"
	"ocra_server/model/response"
	subscribe_service "ocra_server/service/subscribe"
	"sync"

	"github.com/labstack/echo/v4"
)

type SubscribeControllerImpl struct {
	SubsService subscribe_service.SubscribeService
}

func NewSubsController(service subscribe_service.SubscribeService) SubscribeController {
	var doOnce sync.Once
	controller := new(SubscribeControllerImpl)

	doOnce.Do(func() {
		controller = &SubscribeControllerImpl{
			SubsService: service,
		}
	})

	return controller
}

func (controller *SubscribeControllerImpl) SubscribeChannel(ctx echo.Context) error {
	req := new(entity.Subscribes)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.SubsService.CreateSubsRecord(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccesSubscribe,
	})
}

func (controller *SubscribeControllerImpl) UnsubscribeChannel(ctx echo.Context) error {
	req := new(entity.Subscribes)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.SubsService.DeleteSubsRecord(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccesUnsubscribe,
	})
}
