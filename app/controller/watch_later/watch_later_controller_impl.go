package watchlater_controller

import (
	"net/http"
	"ocra_server/helper"
	"ocra_server/model/entity"
	"ocra_server/model/response"
	watchlater_service "ocra_server/service/watch_later"
	"sync"

	"github.com/labstack/echo/v4"
)

type WatchLaterControllerImpl struct {
	Service watchlater_service.WatchLaterService
}

func NewWatchLaterController(service watchlater_service.WatchLaterService) WatchLaterController {
	var doOnce sync.Once
	controller := new(WatchLaterControllerImpl)

	doOnce.Do(func() {
		controller = &WatchLaterControllerImpl{
			Service: service,
		}
	})

	return controller
}

func (controller *WatchLaterControllerImpl) GetAllWatchLaterRecords(ctx echo.Context) error {
	userIdFromCookie := ctx.Request().Header.Get("user-id")

	result, err := controller.Service.GetAllWatchLaterRecords(userIdFromCookie)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageDifferentUserId,
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessGetAllWatchLaterRecord,
		Data:    result,
	})
}

func (controller *WatchLaterControllerImpl) CreateWatchLaterRecord(ctx echo.Context) error {
	req := new(entity.Watch_Laters)
	userIdFromCookie := ctx.Request().Header.Get("user-id")

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if userIdFromCookie != req.UserId {
		return ctx.JSON(http.StatusUnauthorized, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageDifferentUserId,
		})
	}

	if err := helper.ValidateUserInput(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.Service.CreateWatchLaterRecord(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, &response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessCreateWatchLaterRecord,
	})
}

func (controller *WatchLaterControllerImpl) DeleteWatchLaterRecord(ctx echo.Context) error {
	req := new(entity.Watch_Laters)
	userIdFromCookie := ctx.Request().Header.Get("user-id")

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if userIdFromCookie != req.UserId {
		return ctx.JSON(http.StatusUnauthorized, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageDifferentUserId,
		})
	}

	if err := helper.ValidateUserInput(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.Service.DeleteWatchLaterRecord(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessDeleteWatchLaterRecord,
	})
}
