package watchlater_controller

import (
	"net/http"
	"ocra_server/helper"
	"ocra_server/model/entity"
	"ocra_server/model/response"

	history_service "ocra_server/service/history"
	"sync"

	"github.com/labstack/echo/v4"
)

type HistoryControllerImpl struct {
	Service history_service.HistoryService
}

func NewHistoryController(service history_service.HistoryService) HistoryController {
	var doOnce sync.Once
	controller := new(HistoryControllerImpl)

	doOnce.Do(func() {
		controller = &HistoryControllerImpl{
			Service: service,
		}
	})

	return controller
}

func (controller *HistoryControllerImpl) GetAllHistoryRecords(ctx echo.Context) error {
	userIdFromCookie := ctx.Request().Header.Get("user-id")

	result, err := controller.Service.GetAllHistoryRecords(userIdFromCookie)
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

func (controller *HistoryControllerImpl) CreateHistoryRecord(ctx echo.Context) error {
	req := new(entity.Histories)
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

	if err := controller.Service.CreateHistoryRecord(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, &response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessCreateHistoryRecord,
	})
}

func (controller *HistoryControllerImpl) DeleteHistoryRecord(ctx echo.Context) error {
	req := new(entity.Histories)
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

	if err := controller.Service.DeleteHistoryRecord(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessDeleteHistoryRecord,
	})
}
